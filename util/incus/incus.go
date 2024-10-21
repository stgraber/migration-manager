package incus

import (
	"context"
	"fmt"
	"os"
	"path"

	"github.com/lxc/incus/v6/client"
	"github.com/lxc/incus/v6/shared/api"
	"github.com/lxc/incus/v6/shared/cliconfig"

	"github.com/FuturFusion/migration-manager/util"
	"github.com/FuturFusion/migration-manager/util/revert"
)

type IncusClient struct {
	bootableISOPool   string
	bootableISOSource string
	client            incus.InstanceServer
	ctx               context.Context
}

func NewIncusClient(ctx context.Context, incusRemoteName string, bootableISOPool string, bootableISOSource string) (*IncusClient, error) {
	var incusServer incus.InstanceServer
	var err error

	if incusRemoteName == "" {
		incusServer, err = connectLocal()
	} else {
		incusServer, err = connectRemote(incusRemoteName)
	}

	if err != nil {
		return nil, err
	}

	return &IncusClient{
		bootableISOPool:   bootableISOPool,
		bootableISOSource: bootableISOSource,
		client:            incusServer,
		ctx:               ctx,
	}, nil
}

func connectLocal() (incus.InstanceServer, error) {
	args := incus.ConnectionArgs{}
	args.UserAgent = fmt.Sprintf("LXC-MIGRATE %s", "1.2.3") // FIXME was pulling in internal version string from Incus

	return incus.ConnectIncusUnix("", &args)
}

func connectRemote(remoteName string) (incus.InstanceServer, error) {
	incusConfig, err := cliconfig.LoadConfig(path.Join(os.Getenv("HOME"), ".config", "incus", "config.yml"))
	if err != nil {
		return nil, err
	}

	return incusConfig.GetInstanceServer(remoteName)
}

func (c *IncusClient) CreateInstance(instanceArgs api.InstancesPost, disks []util.DiskInfo, nics []util.NICInfo) error {
        revert := revert.New()
        defer revert.Fail()

	// Fetch default profile to get proper root device definition.
	profile, _, err := c.client.GetProfile("default")
	if err != nil && len(disks) != 0 {
		return err
	}

	instanceArgs.Devices = make(map[string]map[string]string)

	// Get the existing root device, if it exists.
	defaultRoot, exists := profile.Devices["root"]
	if !exists {
		defaultRoot = map[string]string{
			"path": "/",
			"pool": "default",
			"type": "disk",
		}
	}

	// Add empty disk(s) from VM definition that will be synced later.
	for i, disk := range disks {
		fmt.Printf("    Adding disk with capacity %0.2fGiB from source '%s'\n", float32(disk.Size)/1024/1024/1024, disk.Name)
		diskKey := "root"
		if i != 0 {
			diskKey = "disk" + string(i)
		}

		instanceArgs.Devices[diskKey] = make(map[string]string)
		for k, v := range defaultRoot {
			instanceArgs.Devices[diskKey][k] = v
		}
		instanceArgs.Devices[diskKey]["size"] = fmt.Sprintf("%dB", disk.Size)

		if i != 0 {
			instanceArgs.Devices[diskKey]["path"] = diskKey
		}
	}

	// Attach bootable ISO to run migration of this VM
	instanceArgs.Devices["migration-iso"] = map[string]string{
		"type": "disk",
		"pool": c.bootableISOPool,
		"source": c.bootableISOSource,
		"boot.priority": "10",
	}

	for _, nic := range nics {
		// FIXME actually add to instance
		fmt.Printf("    Adding NIC for network %q with MAC %s\n", nic.Network, nic.Hwaddr)
	}

        // Create the instance.
        op, err := c.client.CreateInstance(instanceArgs)
        if err != nil {
                return err
        }

        revert.Add(func() {
		_, _ = c.client.DeleteInstance(instanceArgs.Name)
        })

	err = op.Wait()
	if err != nil {
		return err
	}

	revert.Success()

	return nil
}

func (c *IncusClient) GetVMNames() ([]string, error) {
	ret := []string{}
	instances, err := c.client.GetInstances(api.InstanceTypeVM)
	if err != nil {
		return ret, err
	}

	for _, instance := range instances {
		ret = append(ret, instance.Name)
	}

	return ret, nil
}

func (c *IncusClient) DeleteVM(vm string) error {
	op, err := c.client.DeleteInstance(vm)
	if err != nil {
		return err
	}

	return op.Wait()
}
