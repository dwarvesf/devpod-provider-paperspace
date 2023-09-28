package paperspace

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/dwarvesf/devpod-provider-paperspace/pkg/options"
	"github.com/loft-sh/devpod/pkg/client"
	"github.com/loft-sh/log"
	"github.com/pkg/errors"
)

type PaperspaceProvider struct {
	Config *options.Options
	Client *Client
}

func NewProvider(logs log.Logger, init bool) (*PaperspaceProvider, error) {
	ppsAPIKey := os.Getenv("PPS_API_KEY")
	if ppsAPIKey == "" {
		return nil, errors.Errorf("PPS_API_KEY is not set")
	}

	config, err := options.FromEnv(init)

	if err != nil {
		return nil, err
	}

	client := NewClient(ppsAPIKey)
	if err != nil {
		return nil, err
	}
	provider := &PaperspaceProvider{
		Config: config,
		Client: client,
	}
	return provider, nil
}

func GetDevpodInstance(paperspaceProvider *PaperspaceProvider) (*GetMachineResponse, error) {
	servers, err := paperspaceProvider.Client.GetMachines(GetMachinesParams{
		Name: paperspaceProvider.Config.MachineName,
	})
	if err != nil {
		return nil, err
	}

	if len(servers) == 0 {
		return nil, fmt.Errorf("no devpod instance found")
	}

	machine, err := paperspaceProvider.Client.GetMachine(GetMachineParams{
		MachineID: servers[0].ID,
	})
	if err != nil {
		return nil, err
	}
	return &machine, nil
}

func Create(paperspaceProvider *PaperspaceProvider) error {
	sizeGB, _ := strconv.Atoi(paperspaceProvider.Config.DiskSizeGB)
	_, err := paperspaceProvider.Client.CreateMachine(CreateMachineParams{
		MachineName:     paperspaceProvider.Config.MachineName,
		TemplateID:      paperspaceProvider.Config.MachineTemplate,
		MachineType:     paperspaceProvider.Config.MachineType,
		Region:          paperspaceProvider.Config.Region,
		Size:            sizeGB,
		BillingType:     "hourly",
		StartOnCreate:   true,
		DynamicPublicIP: true,
	})
	if err != nil {
		return err
	}

	stillCreating := true
	for stillCreating {
		devPodInstance, err := GetDevpodInstance(paperspaceProvider)
		if err != nil {
			return err
		}

		state := devPodInstance.State

		if state == Ready {
			stillCreating = false
		} else {
			time.Sleep(2 * time.Second)
		}
	}

	return nil
}

func Delete(paperspaceProvider *PaperspaceProvider) error {
	devPodInstance, err := GetDevpodInstance(paperspaceProvider)
	if err != nil {
		return err
	}

	_, err = paperspaceProvider.Client.DestroyMachine(DestroyMachineParams{
		MachineID: devPodInstance.ID,
	})
	if err != nil {
		return err
	}

	return nil
}

func Start(paperspaceProvider *PaperspaceProvider) error {
	devPodInstance, err := GetDevpodInstance(paperspaceProvider)
	if err != nil {
		return err
	}

	_, err = paperspaceProvider.Client.StartMachine(StartMachineParams{
		MachineID: devPodInstance.ID,
	})
	if err != nil {
		return err
	}

	stillCreating := true
	for stillCreating {
		devPodInstance, err := GetDevpodInstance(paperspaceProvider)
		if err != nil {
			return err
		}

		state := devPodInstance.State

		if state == Ready {
			stillCreating = false
		} else {
			time.Sleep(2 * time.Second)
		}
	}

	return nil
}

func Status(paperspaceProvider *PaperspaceProvider) (client.Status, error) {
	devPodInstance, err := GetDevpodInstance(paperspaceProvider)
	if err != nil {
		return client.StatusNotFound, nil
	}

	state := devPodInstance.State

	switch {
	case state == Ready:
		return client.StatusRunning, nil
	case state == Off:
		return client.StatusStopped, nil
	default:
		return client.StatusBusy, nil
	}
}

func Stop(paperspaceProvider *PaperspaceProvider) error {
	devPodInstance, err := GetDevpodInstance(paperspaceProvider)
	if err != nil {
		return err
	}

	_, err = paperspaceProvider.Client.StopMachine(StopMachineParams{
		MachineID: devPodInstance.ID,
	})
	if err != nil {
		return err
	}

	stillCreating := true
	for stillCreating {
		devPodInstance, err := GetDevpodInstance(paperspaceProvider)
		if err != nil {
			return err
		}

		state := devPodInstance.State

		if state == Ready {
			stillCreating = false
		} else {
			time.Sleep(2 * time.Second)
		}
	}

	return nil
}

func Init(paperspaceProvider *PaperspaceProvider) error {
	_, err := paperspaceProvider.Client.GetMachines(GetMachinesParams{})
	if err != nil {
		return err
	}
	return nil
}
