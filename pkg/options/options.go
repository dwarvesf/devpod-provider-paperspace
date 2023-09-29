package options

import (
	"fmt"
	"os"
)

var (
	PPS_DEFAULT_REGION   = "PPS_DEFAULT_REGION"
	PPS_MACHINE_TYPE     = "PPS_MACHINE_TYPE"
	PPS_IMAGE            = "PPS_IMAGE"
	PPS_DISK_SIZE        = "PPS_DISK_SIZE"
	PPS_MACHINE_NAME     = "PPS_MACHINE_NAME"
	PPS_MACHINE_TEMPLATE = "PPS_MACHINE_TEMPLATE"
	MACHINE_ID           = "MACHINE_ID"
	SSH_FOLDER           = "SSH_FOLDER"
)

type Options struct {
	Image       string
	MachineType string
	DiskSizeGB  string
	Region      string

	MachineID       string
	MachineTemplate string
	SSHFolder       string
}

func ConfigFromEnv() (Options, error) {
	return Options{
		Image:       os.Getenv(PPS_IMAGE),
		MachineType: os.Getenv(PPS_MACHINE_TYPE),
		Region:      os.Getenv(PPS_DEFAULT_REGION),
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

	retOptions.Region, err = fromEnvOrError(PPS_DEFAULT_REGION)
	if err != nil {
		return nil, err
	}

	retOptions.SSHFolder, err = fromEnvOrError(SSH_FOLDER)
	if err != nil {
		return nil, err
	}

	// Return early if we're just doing init
	if init {
		return retOptions, nil
	}

	retOptions.MachineID, err = fromEnvOrError(MACHINE_ID)
	if err != nil {
		return nil, err
	}
	// prefix with devpod-
	retOptions.MachineID = "devpod-" + retOptions.MachineID

	retOptions.MachineTemplate, err = fromEnvOrError(PPS_MACHINE_TEMPLATE)
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
