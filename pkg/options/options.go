package options

import (
	"fmt"
	"os"
)

var (
	PPS_DEFAULT_PROJECT_ID      = "PPS_DEFAULT_PROJECT_ID"
	PPS_DEFAULT_ORGANIZATION_ID = "PPS_DEFAULT_ORGANIZATION_ID"
	PPS_DEFAULT_REGION          = "PPS_DEFAULT_REGION"
	PPS_DEFAULT_ZONE            = "PPS_DEFAULT_ZONE"
	PPS_MACHINE_TYPE            = "PPS_MACHINE_TYPE"
	PPS_IMAGE                   = "PPS_IMAGE"
	PPS_DISK_SIZE               = "PPS_DISK_SIZE"
	MACHINE_ID                  = "MACHINE_ID"
	MACHINE_FOLDER              = "MACHINE_FOLDER"
)

type Options struct {
	Image          string
	MachineType    string
	DiskSizeGB     string
	Zone           string
	OrganizationID string
	ProjectID      string
	ServerID       string

	MachineID     string
	MachineFolder string
}

func ConfigFromEnv() (Options, error) {
	return Options{
		Image:       os.Getenv(PPS_IMAGE),
		MachineType: os.Getenv(PPS_MACHINE_TYPE),
		Zone:        os.Getenv(PPS_DEFAULT_ZONE),
	}, nil
}

func FromEnv(init bool) (*Options, error) {
	retOptions := &Options{}

	var err error

	retOptions.Image, err = fromEnvOrError(PPS_IMAGE)
	if err != nil {
		return nil, err
	}
	retOptions.DiskSizeGB, err = fromEnvOrError(PPS_DISK_SIZE)
	if err != nil {
		return nil, err
	}
	retOptions.MachineType, err = fromEnvOrError(PPS_MACHINE_TYPE)
	if err != nil {
		return nil, err
	}

	retOptions.Zone, err = fromEnvOrError(PPS_DEFAULT_ZONE)
	if err != nil {
		return nil, err
	}
	retOptions.OrganizationID, err = fromEnvOrError(PPS_DEFAULT_ORGANIZATION_ID)
	if err != nil {
		return nil, err
	}
	retOptions.ProjectID, err = fromEnvOrError(PPS_DEFAULT_PROJECT_ID)
	if err != nil {
		return nil, err
	}

	// Return eraly if we're just doing init
	if init {
		return retOptions, nil
	}

	retOptions.MachineID, err = fromEnvOrError(MACHINE_ID)
	if err != nil {
		return nil, err
	}
	// prefix with devpod-
	retOptions.MachineID = "devpod-" + retOptions.MachineID

	retOptions.MachineFolder, err = fromEnvOrError(MACHINE_FOLDER)
	if err != nil {
		return nil, err
	}
	return retOptions, nil
}

func fromEnvOrError(name string) (string, error) {
	val := os.Getenv(name)
	if val == "" {
		return "", fmt.Errorf(
			"couldn't find option %s in environment, please make sure %s is defined",
			name,
			name,
		)
	}

	return val, nil
}
