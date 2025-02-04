package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/FuturFusion/migration-manager/internal/batch"
	"github.com/FuturFusion/migration-manager/internal/instance"
	"github.com/FuturFusion/migration-manager/internal/server/auth"
	"github.com/FuturFusion/migration-manager/internal/server/response"
	"github.com/FuturFusion/migration-manager/shared/api"
)

var queueRootCmd = APIEndpoint{
	Path: "queue",

	Get: APIEndpointAction{Handler: queueRootGet, AccessHandler: allowPermission(auth.ObjectTypeServer, auth.EntitlementCanView)},
}

var queueCmd = APIEndpoint{
	Path: "queue/{uuid}",

	Get: APIEndpointAction{Handler: queueGet, AccessHandler: allowPermission(auth.ObjectTypeServer, auth.EntitlementCanView)},
}

var queueWorkerCmd = APIEndpoint{
	Path: "queue/{uuid}/worker",

	// Endpoints used by the migration worker which authenticates via a randomly-generated UUID unique to each instance.
	Get: APIEndpointAction{Handler: queueWorkerGet, AccessHandler: allowAuthenticated},
	Put: APIEndpointAction{Handler: queueWorkerPut, AccessHandler: allowAuthenticated},
}

// Authenticate a migration worker. Allow a GET for an existing instance so the worker can get its instructions,
// and for PUT require the secret token to be valid when the worker reports back.
func (d *Daemon) workerAccessTokenValid(r *http.Request) bool {
	// Only allow GET and PUT methods.
	if r.Method != http.MethodGet && r.Method != http.MethodPut {
		return false
	}

	// Limit to just queue status updates
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 4 {
		return false
	}

	if pathParts[2] != "queue" {
		return false
	}

	// Ensure we got a valid instance UUID.
	instanceUUID, err := uuid.Parse(pathParts[3])
	if err != nil {
		return false
	}

	// Get the instance.
	var i *instance.InternalInstance
	err = d.db.Transaction(r.Context(), func(ctx context.Context, tx *sql.Tx) error {
		dbInstance, err := d.db.GetInstance(tx, instanceUUID)
		if err != nil {
			return err
		}

		internalInstance, ok := dbInstance.(*instance.InternalInstance)
		if !ok {
			return fmt.Errorf("Wasn't given an InternalInstance?")
		}

		i = internalInstance
		return nil
	})
	if err != nil {
		return false
	}

	if r.Method == http.MethodPut {
		// Get the secret token.
		err = r.ParseForm()
		if err != nil {
			return false
		}

		secretUUID, err := uuid.Parse(r.Form.Get("secret"))
		if err != nil {
			return false
		}

		return secretUUID == i.GetSecretToken()
	}

	// Allow a GET for a valid instance.
	return r.Method == http.MethodGet
}

// swagger:operation GET /1.0/queue queue queueRoot_get
//
//	Get the current migration queue
//
//	Returns a list of all migrations underway (URLs).
//
//	---
//	produces:
//	  - application/json
//	responses:
//	  "200":
//	    description: Migration queue instances
//	    schema:
//	      type: object
//	      description: Sync response
//	      properties:
//	        type:
//	          type: string
//	          description: Response type
//	          example: sync
//	        status:
//	          type: string
//	          description: Status description
//	          example: Success
//	        status_code:
//	          type: integer
//	          description: Status code
//	          example: 200
//	        metadata:
//	          type: array
//	          description: List of migration items in the queue
//                items:
//                  type: string
//                example: |-
//                  [
//                    "/1.0/queue/26fa4eb7-8d4f-4bf8-9a6a-dd95d166dfad",
//                    "/1.0/queue/9aad7f16-0d2e-440e-872f-4e9df2d53367"
//                  ]
//	  "403":
//	    $ref: "#/responses/Forbidden"
//	  "500":
//	    $ref: "#/responses/InternalServerError"

// swagger:operation GET /1.0/queue?recursion=1 queue queueRoot_get_recursion
//
//	Get the current migration queue
//
//	Returns a list of all migrations underway (structs).
//
//	---
//	produces:
//	  - application/json
//	responses:
//	  "200":
//	    description: Migration queue instances
//	    schema:
//	      type: object
//	      description: Sync response
//	      properties:
//	        type:
//	          type: string
//	          description: Response type
//	          example: sync
//	        status:
//	          type: string
//	          description: Status description
//	          example: Success
//	        status_code:
//	          type: integer
//	          description: Status code
//	          example: 200
//	        metadata:
//	          type: array
//	          description: List of sources
//	          items:
//	            $ref: "#/definitions/QueueEntry"
//	  "403":
//	    $ref: "#/responses/Forbidden"
//	  "500":
//	    $ref: "#/responses/InternalServerError"
func queueRootGet(d *Daemon, r *http.Request) response.Response {
	// Parse the recursion field.
	recursion, err := strconv.Atoi(r.FormValue("recursion"))
	if err != nil {
		recursion = 0
	}

	queueItems := []api.QueueEntry{}

	// Get all batches.
	var batches []batch.Batch
	err = d.db.Transaction(r.Context(), func(ctx context.Context, tx *sql.Tx) error {
		dbBatches, err := d.db.GetAllBatches(tx)
		if err != nil {
			return err
		}

		batches = dbBatches
		return nil
	})
	if err != nil {
		return response.BadRequest(fmt.Errorf("Failed to get batches: %w", err))
	}

	// For each batch that has entered the "queued" state or later, add its instances.
	for _, b := range batches {
		if b.GetStatus() == api.BATCHSTATUS_UNKNOWN || b.GetStatus() == api.BATCHSTATUS_DEFINED || b.GetStatus() == api.BATCHSTATUS_READY {
			continue
		}

		id, err := b.GetDatabaseID()
		if err != nil {
			return response.BadRequest(fmt.Errorf("Failed to get database ID for batch '%s': %w", b.GetName(), err))
		}

		var instances []instance.Instance
		err = d.db.Transaction(r.Context(), func(ctx context.Context, tx *sql.Tx) error {
			dbInstances, err := d.db.GetAllInstancesForBatchID(tx, id)
			if err != nil {
				return err
			}

			instances = dbInstances
			return nil
		})
		if err != nil {
			return response.BadRequest(fmt.Errorf("Failed to get instances for batch '%s': %w", b.GetName(), err))
		}

		for _, i := range instances {
			queueItems = append(queueItems, api.QueueEntry{
				InstanceUUID:          i.GetUUID(),
				InstanceName:          i.GetName(),
				MigrationStatus:       i.GetMigrationStatus(),
				MigrationStatusString: i.GetMigrationStatusString(),
				BatchID:               id,
				BatchName:             b.GetName(),
			})
		}
	}

	if recursion == 1 {
		return response.SyncResponse(true, queueItems)
	}

	result := make([]string, 0, len(queueItems))
	for _, q := range queueItems {
		result = append(result, fmt.Sprintf("/%s/queue/%s", api.APIVersion, q.InstanceUUID))
	}

	return response.SyncResponse(true, result)
}

// swagger:operation GET /1.0/queue/{uuid} queue queue_get
//
//	Get migration entry from queue
//
//	Returns details about the specified queue entry.
//
//	---
//	produces:
//	  - application/json
//	responses:
//	  "200":
//	    description: Queue entry
//	    schema:
//	      type: object
//	      description: Sync response
//	      properties:
//	        type:
//	          type: string
//	          description: Response type
//	          example: sync
//	        status:
//	          type: string
//	          description: Status description
//	          example: Success
//	        status_code:
//	          type: integer
//	          description: Status code
//	          example: 200
//	        metadata:
//	          $ref: "#/definitions/QueueEntry"
//	  "403":
//	    $ref: "#/responses/Forbidden"
//	  "500":
//	    $ref: "#/responses/InternalServerError"
func queueGet(d *Daemon, r *http.Request) response.Response {
	UUIDString := r.PathValue("uuid")

	UUID, err := uuid.Parse(UUIDString)
	if err != nil {
		return response.SmartError(err)
	}

	// Get the instance.
	var i *instance.InternalInstance
	err = d.db.Transaction(r.Context(), func(ctx context.Context, tx *sql.Tx) error {
		dbInstance, err := d.db.GetInstance(tx, UUID)
		if err != nil {
			return err
		}

		internalInstance, ok := dbInstance.(*instance.InternalInstance)
		if !ok {
			return fmt.Errorf("Wasn't given an InternalInstance?")
		}

		i = internalInstance
		return nil
	})
	if err != nil {
		return response.BadRequest(fmt.Errorf("Failed to get instance '%s': %w", UUID, err))
	}

	// Don't return info for instances that aren't in the migration queue.
	if i.GetBatchID() == nil || !i.IsMigrating() {
		return response.BadRequest(fmt.Errorf("Instance '%s' isn't in the migration queue", i.GetName()))
	}

	// Get the corresponding batch.
	var b batch.Batch
	err = d.db.Transaction(r.Context(), func(ctx context.Context, tx *sql.Tx) error {
		dbBatch, err := d.db.GetBatchByID(tx, *i.GetBatchID())
		if err != nil {
			return err
		}

		b = dbBatch
		return nil
	})
	if err != nil {
		return response.BadRequest(fmt.Errorf("Failed to get batch: %w", err))
	}

	batchID, _ := b.GetDatabaseID()
	ret := api.QueueEntry{
		InstanceUUID:          i.GetUUID(),
		InstanceName:          i.GetName(),
		MigrationStatus:       i.GetMigrationStatus(),
		MigrationStatusString: i.GetMigrationStatusString(),
		BatchID:               batchID,
		BatchName:             b.GetName(),
	}

	return response.SyncResponseETag(true, ret, ret)
}

// swagger:operation GET /1.0/queue/{uuid}/worker queue queue_worker_get
//
//	Get worker command for instance
//
//	Gets a worker command, if any, for this queued instance.
//
//	---
//	produces:
//	  - application/json
//	responses:
//	  "200":
//	    description: WorkerCommand
//	    schema:
//	      type: object
//	      description: Sync response
//	      properties:
//	        type:
//	          type: string
//	          description: Response type
//	          example: sync
//	        status:
//	          type: string
//	          description: Status description
//	          example: Success
//	        status_code:
//	          type: integer
//	          description: Status code
//	          example: 200
//	        metadata:
//	          $ref: "#/definitions/WorkerCommand"
//	  "403":
//	    $ref: "#/responses/Forbidden"
//	  "500":
//	    $ref: "#/responses/InternalServerError"
func queueWorkerGet(d *Daemon, r *http.Request) response.Response {
	UUIDString := r.PathValue("uuid")

	UUID, err := uuid.Parse(UUIDString)
	if err != nil {
		return response.SmartError(err)
	}

	// Get the instance.
	var i *instance.InternalInstance
	err = d.db.Transaction(r.Context(), func(ctx context.Context, tx *sql.Tx) error {
		dbInstance, err := d.db.GetInstance(tx, UUID)
		if err != nil {
			return err
		}

		internalInstance, ok := dbInstance.(*instance.InternalInstance)
		if !ok {
			return fmt.Errorf("Wasn't given an InternalInstance?")
		}

		i = internalInstance
		return nil
	})
	if err != nil {
		return response.BadRequest(fmt.Errorf("Failed to get instance '%s': %w", UUID, err))
	}

	// Don't return info for instances that aren't in the migration queue.
	if i.GetBatchID() == nil || !i.IsMigrating() {
		return response.BadRequest(fmt.Errorf("Instance '%s' isn't in the migration queue", i.GetName()))
	}

	// If the instance is already doing something, don't start something else.
	if i.MigrationStatus != api.MIGRATIONSTATUS_IDLE {
		return response.BadRequest(fmt.Errorf("Instance '%s' isn't idle: %s (%s)", i.GetInventoryPath(), i.MigrationStatus.String(), i.MigrationStatusString))
	}

	// Setup the default "idle" command
	cmd := api.WorkerCommand{
		Command:       api.WORKERCOMMAND_IDLE,
		InventoryPath: i.InventoryPath,
		SourceType:    api.SOURCETYPE_UNKNOWN,
		Source:        []byte(`{}`),
		OS:            i.OS,
		OSVersion:     i.OSVersion,
	}

	// Fetch the source for the instance.
	ms, err := d.source.GetByID(r.Context(), i.SourceID)
	if err != nil {
		return response.BadRequest(fmt.Errorf("Failed to get source '%s': %w", UUID, err))
	}

	s := api.Source{
		DatabaseID: ms.ID,
		Name:       ms.Name,
		Insecure:   ms.Insecure,
		SourceType: ms.SourceType,
		Properties: ms.Properties,
	}

	// Fetch the batch for the instance.
	var b batch.Batch
	err = d.db.Transaction(r.Context(), func(ctx context.Context, tx *sql.Tx) error {
		dbBatch, err := d.db.GetBatchByID(tx, *i.GetBatchID())
		if err != nil {
			return err
		}

		b = dbBatch
		return nil
	})
	if err != nil {
		return response.BadRequest(fmt.Errorf("Failed to get batch '%d': %w", *i.GetBatchID(), err))
	}

	// Determine what action, if any, the worker should start.
	if i.NeedsDiskImport && disksSupportDifferentialSync(i.Disks) {
		// If we can do a background disk sync, kick it off.
		cmd.Command = api.WORKERCOMMAND_IMPORT_DISKS
		cmd.SourceType = api.SOURCETYPE_VMWARE
		cmd.Source, _ = json.Marshal(s)

		i.MigrationStatus = api.MIGRATIONSTATUS_BACKGROUND_IMPORT
		i.MigrationStatusString = i.MigrationStatus.String()
	} else if !b.GetMigrationWindowStart().IsZero() && b.GetMigrationWindowStart().Before(time.Now().UTC()) {
		// If a migration window has been defined, and we have passed the start time, begin the final migration.
		cmd.Command = api.WORKERCOMMAND_FINALIZE_IMPORT
		cmd.SourceType = api.SOURCETYPE_VMWARE
		cmd.Source, _ = json.Marshal(s)

		i.MigrationStatus = api.MIGRATIONSTATUS_FINAL_IMPORT
		i.MigrationStatusString = api.MIGRATIONSTATUS_FINAL_IMPORT.String()
	} else if b.GetMigrationWindowStart().IsZero() {
		// If no migration window start time has been defined, go ahead and begin the final migration.
		cmd.Command = api.WORKERCOMMAND_FINALIZE_IMPORT
		cmd.SourceType = api.SOURCETYPE_VMWARE
		cmd.Source, _ = json.Marshal(s)

		i.MigrationStatus = api.MIGRATIONSTATUS_FINAL_IMPORT
		i.MigrationStatusString = api.MIGRATIONSTATUS_FINAL_IMPORT.String()
	}

	// Update instance in the database.
	err = d.db.Transaction(r.Context(), func(ctx context.Context, tx *sql.Tx) error {
		return d.db.UpdateInstanceStatus(tx, UUID, i.MigrationStatus, i.MigrationStatusString, i.NeedsDiskImport)
	})
	if err != nil {
		return response.SmartError(fmt.Errorf("Failed updating instance '%s': %w", i.GetUUID(), err))
	}

	return response.SyncResponseETag(true, cmd, cmd)
}

func disksSupportDifferentialSync(disks []api.InstanceDiskInfo) bool {
	for _, disk := range disks {
		if disk.Type == "HDD" && disk.DifferentialSyncSupported {
			return true
		}
	}

	return false
}

// swagger:operation PUT /1.0/queue/{uuid}/worker queue queue_worker_put
//
//	Sets worker response for instance
//
//	Sets the response from the worker for this queued instance.
//
//	---
//	consumes:
//	  - application/json
//	produces:
//	  - application/json
//	parameters:
//	  - in: body
//	    name: response
//	    description: WorkerResponse definition
//	    required: true
//	    schema:
//	      $ref: "#/definitions/WorkerResponse"
//	responses:
//	  "200":
//	    $ref: "#/responses/EmptySyncResponse"
//	  "400":
//	    $ref: "#/responses/BadRequest"
//	  "403":
//	    $ref: "#/responses/Forbidden"
//	  "412":
//	    $ref: "#/responses/PreconditionFailed"
//	  "500":
//	    $ref: "#/responses/InternalServerError"
func queueWorkerPut(d *Daemon, r *http.Request) response.Response {
	UUIDString := r.PathValue("uuid")

	UUID, err := uuid.Parse(UUIDString)
	if err != nil {
		return response.SmartError(err)
	}

	// Get the instance.
	var i *instance.InternalInstance
	err = d.db.Transaction(r.Context(), func(ctx context.Context, tx *sql.Tx) error {
		dbInstance, err := d.db.GetInstance(tx, UUID)
		if err != nil {
			return err
		}

		internalInstance, ok := dbInstance.(*instance.InternalInstance)
		if !ok {
			return fmt.Errorf("Wasn't given an InternalInstance?")
		}

		i = internalInstance
		return nil
	})
	if err != nil {
		return response.BadRequest(fmt.Errorf("Failed to get instance '%s': %w", UUID, err))
	}

	// Don't update instances that aren't in the migration queue.
	if i.GetBatchID() == nil || !i.IsMigrating() {
		return response.BadRequest(fmt.Errorf("Instance '%s' isn't in the migration queue", i.GetName()))
	}

	// Decode the command response.
	var resp api.WorkerResponse
	err = json.NewDecoder(r.Body).Decode(&resp)
	if err != nil {
		return response.BadRequest(err)
	}

	// Process the response.
	switch resp.Status {
	case api.WORKERRESPONSE_RUNNING:
		i.MigrationStatusString = resp.StatusString
	case api.WORKERRESPONSE_SUCCESS:
		switch i.MigrationStatus {
		case api.MIGRATIONSTATUS_BACKGROUND_IMPORT:
			i.NeedsDiskImport = false
			i.MigrationStatus = api.MIGRATIONSTATUS_IDLE
			i.MigrationStatusString = api.MIGRATIONSTATUS_IDLE.String()
		case api.MIGRATIONSTATUS_FINAL_IMPORT:
			i.MigrationStatus = api.MIGRATIONSTATUS_IMPORT_COMPLETE
			i.MigrationStatusString = api.MIGRATIONSTATUS_IMPORT_COMPLETE.String()
		}

	case api.WORKERRESPONSE_FAILED:
		i.MigrationStatus = api.MIGRATIONSTATUS_ERROR
		i.MigrationStatusString = resp.StatusString
	}

	// Update instance in the database.
	err = d.db.Transaction(r.Context(), func(ctx context.Context, tx *sql.Tx) error {
		return d.db.UpdateInstanceStatus(tx, UUID, i.MigrationStatus, i.MigrationStatusString, i.NeedsDiskImport)
	})
	if err != nil {
		return response.SmartError(fmt.Errorf("Failed updating instance '%s': %w", i.GetUUID(), err))
	}

	return response.SyncResponse(true, nil)
}
