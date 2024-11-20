package main

import (
	"context"
	"database/sql"
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/lxc/incus/v6/shared/logger"

	"github.com/FuturFusion/migration-manager/internal/batch"
	"github.com/FuturFusion/migration-manager/internal/instance"
	"github.com/FuturFusion/migration-manager/internal/source"
	"github.com/FuturFusion/migration-manager/internal/target"
	"github.com/FuturFusion/migration-manager/shared/api"
)

func (d *Daemon) runPeriodicTask(f func() bool, interval time.Duration) {
	go func() {
		for {
			done := f()
			if done {
				return
			}

			t := time.NewTimer(interval)

			select {
			case <-d.shutdownCtx.Done():
				t.Stop()
				return
			case <-t.C:
				t.Stop()
			}
		}
	}()
}

func (d *Daemon) syncInstancesFromSources() bool {
	loggerCtx := logger.Ctx{"method": "syncInstancesFromSources"}

	// Ensure at least one target exists.
	targets := []target.Target{}
	err := d.db.Transaction(d.shutdownCtx, func(ctx context.Context, tx *sql.Tx) error {
		var err error
		targets, err = d.db.GetAllTargets(tx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		logger.Warn(err.Error(), loggerCtx)
		return false
	}
	if len(targets) == 0 {
		logger.Debug("No targets defined, skipping instance sync from sources", loggerCtx)
		return false
	}

	// For now, just default to the first target defined.
	targetId, err := targets[0].GetDatabaseID()
	if err != nil {
		logger.Warn(err.Error(), loggerCtx)
		return false
	}

	// Get the list of configured sources.
	sources := []source.Source{}
	err = d.db.Transaction(d.shutdownCtx, func(ctx context.Context, tx *sql.Tx) error {
		var err error
		sources, err = d.db.GetAllSources(tx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		logger.Warn(err.Error(), loggerCtx)
		return false
	}

	// Check each source for any new, changed, or deleted instances.
	for _, s := range sources {
		err := s.Connect(d.shutdownCtx)
		if err != nil {
			logger.Warn(err.Error(), loggerCtx)
			continue
		}

		instances, err := s.GetAllVMs(d.shutdownCtx)
		if err != nil {
			logger.Warn(err.Error(), loggerCtx)
			continue
		}

		currentInstancesFromSource := make(map[uuid.UUID]bool)

		// Iterate each instance from this source.
		for _, i := range instances {
			// Check if this instance is already in the database.
			existingInstance := &instance.InternalInstance{}
			err := d.db.Transaction(d.shutdownCtx, func(ctx context.Context, tx *sql.Tx) error {
				inst, err := d.db.GetInstance(tx, i.GetUUID())
				if err != nil {
					return err
				}

				existingInstance = inst.(*instance.InternalInstance)
				return nil
			})

			if err == nil {
				// An instance already exists in the database; update with any changes from the source.
				instanceUpdated := false

				// First, check any fields that cannot be changed through the migration manager
				if existingInstance.Name != i.Name {
					existingInstance.Name = i.Name
					instanceUpdated = true
				}

				if existingInstance.Architecture != i.Architecture {
					existingInstance.Architecture = i.Architecture
					instanceUpdated = true
				}

				if existingInstance.OS != i.OS {
					existingInstance.OS = i.OS
					instanceUpdated = true
				}

				if existingInstance.OSVersion != i.OSVersion {
					existingInstance.OSVersion = i.OSVersion
					instanceUpdated = true
				}

				if !slices.Equal(existingInstance.Disks, i.Disks) {
					existingInstance.Disks = i.Disks
					instanceUpdated = true
				}

				if !slices.Equal(existingInstance.NICs, i.NICs) {
					existingInstance.NICs = i.NICs
					instanceUpdated = true
				}

				if existingInstance.UseLegacyBios != i.UseLegacyBios {
					existingInstance.UseLegacyBios = i.UseLegacyBios
					instanceUpdated = true
				}

				if existingInstance.SecureBootEnabled != i.SecureBootEnabled {
					existingInstance.SecureBootEnabled = i.SecureBootEnabled
					instanceUpdated = true
				}

				if existingInstance.TPMPresent != i.TPMPresent {
					existingInstance.TPMPresent = i.TPMPresent
					instanceUpdated = true
				}

				// Next, check fields that can be updated, but only sync if this instance hasn't been manually updated.
				if existingInstance.LastManualUpdate.IsZero() {
					if existingInstance.NumberCPUs != i.NumberCPUs {
						existingInstance.NumberCPUs = i.NumberCPUs
						instanceUpdated = true
					}

					if existingInstance.MemoryInMiB != i.MemoryInMiB {
						existingInstance.MemoryInMiB = i.MemoryInMiB
						instanceUpdated = true
					}
				} else {
					logger.Debug("Instance " +  i.GetName() + " (" + i.GetUUID().String() + ") has been manually updated, skipping some automatic sync updates", loggerCtx)
				}

				if instanceUpdated {
					logger.Info("Syncing changes to instance " + i.GetName() + " (" + i.GetUUID().String() + ") from source " + s.GetName(), loggerCtx)
					existingInstance.LastUpdateFromSource = i.LastUpdateFromSource
					err := d.db.Transaction(d.shutdownCtx, func(ctx context.Context, tx *sql.Tx) error {
						err := d.db.UpdateInstance(tx, existingInstance)
						if err != nil {
							return err
						}

						return nil
					})

					if err != nil {
						logger.Warn(err.Error(), loggerCtx)
						continue
					}
				}
			} else {
				// Add a new instance to the database.
				logger.Info("Adding instance " + i.GetName() + " (" + i.GetUUID().String() + ") from source " + s.GetName() + " to database", loggerCtx)
				i.TargetID = targetId

				err := d.db.Transaction(d.shutdownCtx, func(ctx context.Context, tx *sql.Tx) error {
					err := d.db.AddInstance(tx, &i)
					if err != nil {
						return err
					}

					return nil
				})

				if err != nil {
					logger.Warn(err.Error(), loggerCtx)
					continue
				}
			}

			// Record that this instance exists.
			currentInstancesFromSource[i.GetUUID()] = true
		}

		// Remove instances that no longer exist in this source.
		allDBInstances := []instance.Instance{}
		err = d.db.Transaction(d.shutdownCtx, func(ctx context.Context, tx *sql.Tx) error {
			instances, err := d.db.GetAllInstances(tx)
			if err != nil {
				return err
			}

			allDBInstances = instances
			return nil
		})
		if err != nil {
			logger.Warn(err.Error(), loggerCtx)
			continue
		}

		for _, i := range allDBInstances {
			_, instanceExists := currentInstancesFromSource[i.GetUUID()]
			if !instanceExists {
				logger.Info("Instance " + i.GetName() + " (" + i.GetUUID().String() + ") removed from source " + s.GetName(), loggerCtx)
				err := d.db.Transaction(d.shutdownCtx, func(ctx context.Context, tx *sql.Tx) error {
					err := d.db.DeleteInstance(tx, i.GetUUID())
					if err != nil {
						return err
					}

					return nil
				})
				if err != nil {
					logger.Warn(err.Error(), loggerCtx)
					continue
				}
			}
		}
	}

	return false
}

func (d *Daemon) processReadyBatches() bool {
	loggerCtx := logger.Ctx{"method": "processReadyBatches"}

	// Get any batches in the "ready" state.
	batches := []batch.Batch{}
	err := d.db.Transaction(d.shutdownCtx, func(ctx context.Context, tx *sql.Tx) error {
		var err error
		batches, err = d.db.GetAllBatchesByState(tx, api.BATCHSTATUS_READY)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		logger.Warn(err.Error(), loggerCtx)
		return false
	}

	// Do some basic sanity check of each batch before adding it to the queue.
	for _, b := range batches {
		logger.Info("Batch '" + b.GetName() + "' status is 'Ready', processing....", loggerCtx)
		batchID, err := b.GetDatabaseID()
		if err != nil {
			logger.Warn(err.Error(), loggerCtx)
			continue
		}

		// If a migration window is defined, ensure sure it makes sense.
		if !b.GetMigrationWindowStart().IsZero() && !b.GetMigrationWindowEnd().IsZero() && b.GetMigrationWindowEnd().Before(b.GetMigrationWindowStart()) {
			logger.Error("Batch '" + b.GetName() + "' window end time is before its start time", loggerCtx)

			err = d.db.Transaction(d.shutdownCtx, func(ctx context.Context, tx *sql.Tx) error {
				err := d.db.UpdateBatchStatus(tx, batchID, api.BATCHSTATUS_ERROR, "Migration window end before start")
				if err != nil {
					return err
				}

				return nil
			})
			if err != nil {
				logger.Warn(err.Error(), loggerCtx)
				continue
			}

			continue
		}
		if !b.GetMigrationWindowEnd().IsZero() && b.GetMigrationWindowEnd().Before(time.Now().UTC()) {
			logger.Error("Batch '" + b.GetName() + "' window end time has already passed", loggerCtx)

			err = d.db.Transaction(d.shutdownCtx, func(ctx context.Context, tx *sql.Tx) error {
				err := d.db.UpdateBatchStatus(tx, batchID, api.BATCHSTATUS_ERROR, "Migration window end has already passed")
				if err != nil {
					return err
				}

				return nil
			})
			if err != nil {
				logger.Warn(err.Error(), loggerCtx)
				continue
			}

			continue
		}

		// Get all instances for this batch.
		instances := []instance.Instance{}
		err = d.db.Transaction(d.shutdownCtx, func(ctx context.Context, tx *sql.Tx) error {
			var err error
			instances, err = d.db.GetAllInstancesForBatchID(tx, batchID)
			if err != nil {
				return err
			}

			return nil
		})
		if err != nil {
			logger.Warn(err.Error(), loggerCtx)
			continue
		}

		// If no instances apply to this batch, return an error.
		if len(instances) == 0 {
			logger.Error("Batch '" + b.GetName() + "' has no instances", loggerCtx)

			err = d.db.Transaction(d.shutdownCtx, func(ctx context.Context, tx *sql.Tx) error {
				err := d.db.UpdateBatchStatus(tx, batchID, api.BATCHSTATUS_ERROR, "No instances assigned")
				if err != nil {
					return err
				}

				return nil
			})
			if err != nil {
				logger.Warn(err.Error(), loggerCtx)
				continue
			}

			continue
		}

		// No issues detected, move to "queued" status.
		logger.Info("Updating batch '" + b.GetName() + "' status to 'Queued'", loggerCtx)

		err = d.db.Transaction(d.shutdownCtx, func(ctx context.Context, tx *sql.Tx) error {
			var state api.BatchStatusType = api.BATCHSTATUS_QUEUED
			err := d.db.UpdateBatchStatus(tx, batchID, state, state.String())
			if err != nil {
				return err
			}

			return nil
		})
		if err != nil {
			logger.Warn(err.Error(), loggerCtx)
			continue
		}
	}

	return false
}

func (d *Daemon) processQueuedBatches() bool {
	loggerCtx := logger.Ctx{"method": "processQueuedBatches"}

	// Get any batches in the "queued" state.
	batches := []batch.Batch{}
	err := d.db.Transaction(d.shutdownCtx, func(ctx context.Context, tx *sql.Tx) error {
		var err error
		batches, err = d.db.GetAllBatchesByState(tx, api.BATCHSTATUS_QUEUED)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		logger.Warn(err.Error(), loggerCtx)
		return false
	}

	// See if we can start running this batch.
	for _, b := range batches {
		logger.Info("Batch '" + b.GetName() + "' status is 'Queued', processing....", loggerCtx)
		batchID, err := b.GetDatabaseID()
		if err != nil {
			logger.Warn(err.Error(), loggerCtx)
			continue
		}

		if !b.GetMigrationWindowStart().IsZero() && b.GetMigrationWindowStart().After(time.Now().UTC()) {
			logger.Debug("Start of migration window hasn't arrived yet", loggerCtx)
			continue
		}

		if !b.GetMigrationWindowEnd().IsZero() && b.GetMigrationWindowEnd().Before(time.Now().UTC()) {
			logger.Error("Batch '" + b.GetName() + "' window end time has already passed", loggerCtx)

			err = d.db.Transaction(d.shutdownCtx, func(ctx context.Context, tx *sql.Tx) error {
				err := d.db.UpdateBatchStatus(tx, batchID, api.BATCHSTATUS_ERROR, "Migration window end has already passed")
				if err != nil {
					return err
				}

				return nil
			})
			if err != nil {
				logger.Warn(err.Error(), loggerCtx)
				continue
			}

			continue
		}

		// Get all instances for this batch.
		instances := []instance.Instance{}
		err = d.db.Transaction(d.shutdownCtx, func(ctx context.Context, tx *sql.Tx) error {
			var err error
			instances, err = d.db.GetAllInstancesForBatchID(tx, batchID)
			if err != nil {
				return err
			}

			return nil
		})
		if err != nil {
			logger.Warn(err.Error(), loggerCtx)
			continue
		}

		// Instantiate each new empty VM in Incus.
		for _, i := range instances {
			// Update the instance status.
			err := d.db.Transaction(d.shutdownCtx, func(ctx context.Context, tx *sql.Tx) error {
				var state api.MigrationStatusType = api.MIGRATIONSTATUS_CREATING
				err := d.db.UpdateInstanceStatus(tx, i.GetUUID(), state, state.String(), true)
				if err != nil {
					return err
				}

				return nil
			})
			if err != nil {
				logger.Warn(err.Error(), loggerCtx)
				continue
			}

			// Get the target for this instance.
			var t target.Target
			err = d.db.Transaction(d.shutdownCtx, func(ctx context.Context, tx *sql.Tx) error {
				var err error
				t, err = d.db.GetTargetByID(tx, i.GetTargetID())
				if err != nil {
					return err
				}

				return nil
			})
			if err != nil {
				logger.Warn(err.Error(), loggerCtx)
				_ = d.db.Transaction(d.shutdownCtx, func(ctx context.Context, tx *sql.Tx) error {
					var state api.MigrationStatusType = api.MIGRATIONSTATUS_ERROR
					err := d.db.UpdateInstanceStatus(tx, i.GetUUID(), state, err.Error(), true)
					if err != nil {
						return err
					}

					return nil
				})
				continue
			}

			// Connect to the target.
			err = t.Connect(d.shutdownCtx)
			if err != nil {
				logger.Warn(err.Error(), loggerCtx)
				_ = d.db.Transaction(d.shutdownCtx, func(ctx context.Context, tx *sql.Tx) error {
					var state api.MigrationStatusType = api.MIGRATIONSTATUS_ERROR
					err := d.db.UpdateInstanceStatus(tx, i.GetUUID(), state, err.Error(), true)
					if err != nil {
						return err
					}

					return nil
				})
				continue
			}

			// Create the instance.
			internalInstance, _ := i.(*instance.InternalInstance)
			instanceDef := t.CreateVMDefinition(*internalInstance)
			creationErr := t.CreateNewVM(instanceDef)
			if creationErr == nil {
				// Creation was successful, update the instance state to 'Idle'.
				err := d.db.Transaction(d.shutdownCtx, func(ctx context.Context, tx *sql.Tx) error {
					var state api.MigrationStatusType = api.MIGRATIONSTATUS_IDLE
					err := d.db.UpdateInstanceStatus(tx, i.GetUUID(), state, state.String(), true)
					if err != nil {
						return err
					}

					return nil
				})
				if err != nil {
					logger.Warn(err.Error(), loggerCtx)
					continue
				}
			} else {
				logger.Warn(creationErr.Error(), loggerCtx)
				err := d.db.Transaction(d.shutdownCtx, func(ctx context.Context, tx *sql.Tx) error {
					var state api.MigrationStatusType = api.MIGRATIONSTATUS_ERROR
					err := d.db.UpdateInstanceStatus(tx, i.GetUUID(), state, creationErr.Error(), true)
					if err != nil {
						return err
					}

					return nil
				})
				if err != nil {
					logger.Warn(err.Error(), loggerCtx)
					continue
				}
			}

			// Start the instance.
			startErr := t.StartVM(i.GetName())
			if startErr != nil {
				logger.Warn(startErr.Error(), loggerCtx)
				err := d.db.Transaction(d.shutdownCtx, func(ctx context.Context, tx *sql.Tx) error {
					var state api.MigrationStatusType = api.MIGRATIONSTATUS_ERROR
					err := d.db.UpdateInstanceStatus(tx, i.GetUUID(), state, startErr.Error(), true)
					if err != nil {
						return err
					}

					return nil
				})
				if err != nil {
					logger.Warn(err.Error(), loggerCtx)
					continue
				}
			}

			// Inject the worker binary.
			pushErr := t.PushFile(i.GetName(), "./migration-manager-worker", "/root/")
			if pushErr != nil {
				logger.Warn(pushErr.Error(), loggerCtx)
				err := d.db.Transaction(d.shutdownCtx, func(ctx context.Context, tx *sql.Tx) error {
					var state api.MigrationStatusType = api.MIGRATIONSTATUS_ERROR
					err := d.db.UpdateInstanceStatus(tx, i.GetUUID(), state, pushErr.Error(), true)
					if err != nil {
						return err
					}

					return nil
				})
				if err != nil {
					logger.Warn(err.Error(), loggerCtx)
					continue
				}
			}

			// Start the worker binary.
			workerStartErr := t.ExecWithoutWaiting(i.GetName(), []string{"/root/migration-manager-worker", "--endpoint", d.getEndpoint(), "--uuid", i.GetUUID().String()})
			if workerStartErr != nil {
				logger.Warn(workerStartErr.Error(), loggerCtx)
				err := d.db.Transaction(d.shutdownCtx, func(ctx context.Context, tx *sql.Tx) error {
					var state api.MigrationStatusType = api.MIGRATIONSTATUS_ERROR
					err := d.db.UpdateInstanceStatus(tx, i.GetUUID(), state, workerStartErr.Error(), true)
					if err != nil {
						return err
					}

					return nil
				})
				if err != nil {
					logger.Warn(err.Error(), loggerCtx)
					continue
				}
			}
		}

		// Move batch to "running" status.
		logger.Info("Updating batch '" + b.GetName() + "' status to 'Running'", loggerCtx)

		err = d.db.Transaction(d.shutdownCtx, func(ctx context.Context, tx *sql.Tx) error {
			var state api.BatchStatusType = api.BATCHSTATUS_RUNNING
			err := d.db.UpdateBatchStatus(tx, batchID, state, state.String())
			if err != nil {
				return err
			}

			return nil
		})
		if err != nil {
			logger.Warn(err.Error(), loggerCtx)
			continue
		}
	}
	return false
}

func (d *Daemon) finalizeCompleteInstances() bool {
	loggerCtx := logger.Ctx{"method": "finalizeCompleteInstances"}

	// Get any instances in the "complete" state.
	instances := []instance.Instance{}
	err := d.db.Transaction(d.shutdownCtx, func(ctx context.Context, tx *sql.Tx) error {
		var err error
		instances, err = d.db.GetAllInstancesByState(tx, api.MIGRATIONSTATUS_IMPORT_COMPLETE)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		logger.Warn(err.Error(), loggerCtx)
		return false
	}

	for _, i := range instances {
		logger.Info("Finalizing migration steps for instance " + i.GetName(), loggerCtx)
		// Get the target for this instance.
		var t target.Target
		err = d.db.Transaction(d.shutdownCtx, func(ctx context.Context, tx *sql.Tx) error {
			var err error
			t, err = d.db.GetTargetByID(tx, i.GetTargetID())
			if err != nil {
				return err
			}

			return nil
		})
		if err != nil {
			logger.Warn(err.Error(), loggerCtx)
			_ = d.db.Transaction(d.shutdownCtx, func(ctx context.Context, tx *sql.Tx) error {
				var state api.MigrationStatusType = api.MIGRATIONSTATUS_ERROR
				err := d.db.UpdateInstanceStatus(tx, i.GetUUID(), state, err.Error(), true)
				if err != nil {
					return err
				}

				return nil
			})
			continue
		}

		// Connect to the target.
		err = t.Connect(d.shutdownCtx)
		if err != nil {
			logger.Warn(err.Error(), loggerCtx)
			_ = d.db.Transaction(d.shutdownCtx, func(ctx context.Context, tx *sql.Tx) error {
				var state api.MigrationStatusType = api.MIGRATIONSTATUS_ERROR
				err := d.db.UpdateInstanceStatus(tx, i.GetUUID(), state, err.Error(), true)
				if err != nil {
					return err
				}

				return nil
			})
			continue
		}

		// Stop the instance.
		stopErr := t.StopVM(i.GetName())
		if stopErr != nil {
			logger.Warn(stopErr.Error(), loggerCtx)
			err := d.db.Transaction(d.shutdownCtx, func(ctx context.Context, tx *sql.Tx) error {
				var state api.MigrationStatusType = api.MIGRATIONSTATUS_ERROR
				err := d.db.UpdateInstanceStatus(tx, i.GetUUID(), state, stopErr.Error(), true)
				if err != nil {
					return err
				}

				return nil
			})
			if err != nil {
				logger.Warn(err.Error(), loggerCtx)
				continue
			}
		}

		// Get the instance definition.
		apiDef, etag, err := t.GetInstance(i.GetName())
		if err != nil {
			logger.Warn(err.Error(), loggerCtx)
			err := d.db.Transaction(d.shutdownCtx, func(ctx context.Context, tx *sql.Tx) error {
				var state api.MigrationStatusType = api.MIGRATIONSTATUS_ERROR
				err := d.db.UpdateInstanceStatus(tx, i.GetUUID(), state, err.Error(), true)
				if err != nil {
					return err
				}

				return nil
			})
			if err != nil {
				logger.Warn(err.Error(), loggerCtx)
				continue
			}
		}

		// Add NIC(s).
		//internalInstance, _ := i.(*instance.InternalInstance)
		for i, nic := range i.(*instance.InternalInstance).NICs {
			deviceName := fmt.Sprintf("eth%d", i)
			apiDef.Devices[deviceName] = make(map[string]string)
			apiDef.Devices[deviceName]["type"] = "nic"
			apiDef.Devices[deviceName]["nictype"] = "macvlan"
			apiDef.Devices[deviceName]["parent"] = "wan4350"
			apiDef.Devices[deviceName]["name"] = deviceName
			apiDef.Devices[deviceName]["hwaddr"] = nic.Hwaddr
			/*for _, profileDevice := range profile.Devices {
				if profileDevice["type"] == "nic" && profileDevice["network"] == "vmware" { // FIXME need to fix up network mappings
					ret.Devices[deviceName] = make(map[string]string)
					for k, v := range profileDevice {
						ret.Devices[deviceName][k] = v
					}
					ret.Devices[deviceName]["hwaddr"] = nic.Hwaddr
				}
			}*/
		}

		// Remove the migration ISO image.
		delete(apiDef.Devices, "migration-iso")

		// Don't set any profiles by default.
		apiDef.Profiles = []string{}

		// Handle Windows-specific completion steps.
		if strings.Contains(apiDef.Config["image.os"], "swodniw") {
			// Remove the drivers ISO image.
			delete(apiDef.Devices, "drivers")

			// Fixup the OS name.
			apiDef.Config["image.os"] = strings.Replace(apiDef.Config["image.os"], "swodniw", "windows", 1)
		}

		// Update the instance in Incus.
		op, updateErr := t.UpdateInstance(i.GetName(), apiDef.Writable(), etag)
		if updateErr != nil {
			logger.Warn(updateErr.Error(), loggerCtx)
			err := d.db.Transaction(d.shutdownCtx, func(ctx context.Context, tx *sql.Tx) error {
				var state api.MigrationStatusType = api.MIGRATIONSTATUS_ERROR
				err := d.db.UpdateInstanceStatus(tx, i.GetUUID(), state, updateErr.Error(), true)
				if err != nil {
					return err
				}

				return nil
			})
			if err != nil {
				logger.Warn(err.Error(), loggerCtx)
				continue
			}
		}
		updateErr = op.Wait()
		if updateErr != nil {
			logger.Warn(updateErr.Error(), loggerCtx)
			err := d.db.Transaction(d.shutdownCtx, func(ctx context.Context, tx *sql.Tx) error {
				var state api.MigrationStatusType = api.MIGRATIONSTATUS_ERROR
				err := d.db.UpdateInstanceStatus(tx, i.GetUUID(), state, updateErr.Error(), true)
				if err != nil {
					return err
				}

				return nil
			})
			if err != nil {
				logger.Warn(err.Error(), loggerCtx)
				continue
			}
		}

		// Update the instance status.
		err = d.db.Transaction(d.shutdownCtx, func(ctx context.Context, tx *sql.Tx) error {
			var state api.MigrationStatusType = api.MIGRATIONSTATUS_FINISHED
			err := d.db.UpdateInstanceStatus(tx, i.GetUUID(), state, state.String(), true)
			if err != nil {
				return err
			}

			return nil
		})
		if err != nil {
			logger.Warn(err.Error(), loggerCtx)
			continue
		}

		// Power on the completed instance.
		startErr := t.StartVM(i.GetName())
		if startErr != nil {
			logger.Warn(startErr.Error(), loggerCtx)
			continue
		}
	}
	return false
}
