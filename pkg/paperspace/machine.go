package paperspace

import (
	"fmt"
	"time"
)

// MachineState represents the state of a Paperspace machine
const (
	Off          = "off"
	Starting     = "starting"
	Stopping     = "stopping"
	Restarting   = "restarting"
	ServiceReady = "serviceReady"
	Ready        = "ready"
	Upgrading    = "upgrading"
	Provisioning = "provisioning"
)

// Machine represents a Paperspace machine
type Machine struct {
	ID                     string `json:"id"`
	Name                   string `json:"name"`
	OS                     string `json:"os"`
	RAM                    string `json:"ram"`
	CPUs                   int    `json:"cpus"`
	GPU                    string `json:"gpu"`
	StorageTotal           string `json:"storageTotal"`
	StorageUsed            string `json:"storageUsed"`
	MachineType            string `json:"machineType"`
	UsageRate              string `json:"usageRate"`
	ShutdownTimeoutInHours int    `json:"shutdownTimeoutInHours"`
	ShutdownTimeoutForces  bool   `json:"shutdownTimeoutForces"`
	PerformAutoSnapshot    bool   `json:"performAutoSnapshot"`
	AutoSnapshotFrequency  string `json:"autoSnapshotFrequency"`
	AutoSnapshotSaveCount  int    `json:"autoSnapshotSaveCount"`
	DynamicPublicIP        bool   `json:"dynamicPublicIp"`
	AgentType              string `json:"agentType"`
	DtCreated              string `json:"dtCreated"`
	State                  string `json:"state"`
	UpdatesPending         bool   `json:"updatesPending"`
	NetworkID              string `json:"networkId"`
	PrivateIPAddress       string `json:"privateIpAddress"`
	PublicIPAddress        string `json:"publicIpAddress"`
	Region                 string `json:"region"`
	ScriptID               string `json:"scriptId"`
	DtLastRun              string `json:"dtLastRun"`
	RestorePointSnapshotID string `json:"restorePointSnapshotId"`
	RestorePointFrequency  string `json:"restorePointFrequency"`
	InternalID             int    `json:"internalId"`
}

// CreateMachineParams represents the parameters for CreateMachine method
type CreateMachineParams struct {
	RequestParams
	MachineName           string `json:"machineName"` // required
	MachineType           string `json:"machineType"` // required
	TemplateID            string `json:"templateId"`  // required
	Region                string `json:"region"`      // required
	Size                  int    `json:"size"`        // required
	BillingType           string `json:"billingType"` // required
	NetworkId             string `json:"networkId,omitempty"`
	AssignPublicIP        bool   `json:"assignPublicIp,omitempty"`
	DynamicPublicIP       bool   `json:"dynamicPublicIp,omitempty"`
	StartOnCreate         bool   `json:"startOnCreate,omitempty"`
	StartupScriptId       string `json:"startupScriptId,omitempty"`
	UserId                string `json:"userId,omitempty"`
	Email                 string `json:"email,omitempty"`
	Password              string `json:"password,omitempty"`
	FirstName             string `json:"firstName,omitempty"`
	LastName              string `json:"lastName,omitempty"`
	NotificationEmail     string `json:"notificationEmail,omitempty"`
	TakeInitialSnapshot   bool   `json:"takeInitialSnapshot,omitempty"`
	RestorePointEnabled   bool   `json:"restorePointEnabled,omitempty"`
	RestorePointFrequency string `json:"restorePointFrequency,omitempty"`
	EnableNvlink          bool   `json:"enableNvlink,omitempty"`
}

// StartMachineParams represents the parameters for StartMachine method
type StartMachineParams struct {
	RequestParams
	MachineID string `json:"machineId"` // required
}

// StopMachineParams represents the parameters for StopMachine method
type StopMachineParams struct {
	RequestParams
	MachineID string `json:"machineId"` // required
}

// RestartMachineParams represents the parameters for StopMachine method
type RestartMachineParams struct {
	RequestParams
	MachineID string `json:"machineId"` // required
}

// DestroyMachineParams represents the parameters for StopMachine method
type DestroyMachineParams struct {
	RequestParams
	MachineID       string `json:"machineId"` // required
	ReleasePublicIP bool   `json:"releasePublicIp,omitempty"`
}

// UpdateMachineParams represents the parameters for UpdateMachine method
type UpdateMachineParams struct {
	RequestParams
	MachineID              string `json:"machineId"`                        // required
	MachineName            string `json:"machineName,omitempty"`            // optional
	ShutdownTimeoutInHours int    `json:"shutdownTimeoutInHours,omitempty"` // optional
	ShutdownTimeoutForces  int    `json:"shutdownTimeoutForces,omitempty"`  // optional
	PerformAutoSnapshot    bool   `json:"performAutoSnapshot,omitempty"`    // optional
	AutoSnapshotFrequency  string `json:"autoSnapshotFrequency,omitempty"`  // optional
	AutoSnapshotSaveCount  int    `json:"autoSnapshotSaveCount,omitempty"`  // optional
	DynamicPublicIP        bool   `json:"dynamicPublicIp,omitempty"`        // optional
	AssignPublicIP         bool   `json:"assignPublicIp,omitempty"`         // optional
}

// GetMachineParams represents the parameters for GetMachine method
type GetMachineParams struct {
	RequestParams
}

type GetMachineResponse struct {
	ID                     string     `json:"id"`
	Name                   string     `json:"name"`
	OS                     string     `json:"os"`
	RAM                    int64      `json:"ram,string"`
	CPUs                   int        `json:"cpus"`
	GPU                    string     `json:"gpu"`
	StorageTotal           int64      `json:"storageTotal,string"`
	StorageUsed            int64      `json:"storageUsed,string"`
	MachineType            string     `json:"machineType"`
	UsageRate              string     `json:"usageRate"`
	ShutdownTimeoutInHours int        `json:"shutdownTimeoutInHours"`
	ShutdownTimeoutForces  bool       `json:"shutdownTimeoutForces"`
	PerformAutoSnapshot    bool       `json:"performAutoSnapshot"`
	AutoSnapshotFrequency  string     `json:"autoSnapshotFrequency"`
	AutoSnapshotSaveCount  int        `json:"autoSnapshotSaveCount"`
	DynamicPublicIP        bool       `json:"dynamicPublicIp"`
	AgentType              string     `json:"agentType"`
	DtCreated              time.Time  `json:"dtCreated"`
	State                  string     `json:"state"`
	UpdatesPending         bool       `json:"updatesPending"`
	NetworkID              string     `json:"networkId"`
	PrivateIPAddress       string     `json:"privateIpAddress"`
	PublicIPAddress        string     `json:"publicIpAddress"`
	Region                 string     `json:"region"`
	ScriptID               *string    `json:"scriptId"`               // pointer to handle null value
	DtLastRun              *time.Time `json:"dtLastRun"`              // pointer to handle null value
	RestorePointSnapshotID *string    `json:"restorePointSnapshotId"` // pointer to handle null value
	RestorePointFrequency  *string    `json:"restorePointFrequency"`  // pointer to handle null value
	Events                 []Event    `json:"events"`                 // slice of Event structs
}

type Event struct {
	Name       string    `json:"name"`
	State      string    `json:"state"`
	ErrorMsg   string    `json:"errorMsg"`
	Handle     string    `json:"handle"`
	DtModified time.Time `json:"dtModified"`
	DtFinished time.Time `json:"dtFinished"`
	DtCreated  time.Time `json:"dtCreated"`
}

// ListMachinesParams represents the parameters object for ListMachines method
type GetMachinesParams struct {
	RequestParams
	Limit                  string `json:"limit,omitempty"`
	Skip                   string `json:"skip,omitempty"`
	MachineID              string `json:"machineId,omitempty"`
	Name                   string `json:"name,omitempty"`
	OS                     string `json:"os,omitempty"`
	RAM                    string `json:"ram,omitempty"`
	CPUs                   int    `json:"cpu,omitempty"`
	GPU                    string `json:"gpu,omitempty"`
	StorageTotal           string `json:"storageTotal,omitempty"`
	StorageUsed            string `json:"storageUsed,omitempty"`
	UsageRate              string `json:"usageRate,omitempty"`
	ShutdownTimeoutInHours int    `json:"shutdownTimeoutInHours,omitempty"`
	PerformAutoSnapshot    bool   `json:"performAutoSnapshot,omitempty"`
	AutoSnapshotFrequency  string `json:"autoSnapshotFrequency,omitempty"`
	AutoSnapshotSaveCount  int    `json:"autoSnapshotSaveCount,omitempty"`
	AgentType              string `json:"agentType,omitempty"`
	DtCreated              string `json:"dtCreated,omitempty"`
	State                  string `json:"state,omitempty"`
	UpdatesPending         string `json:"updatesPending,omitempty"`
	NetworkID              string `json:"networkId,omitempty"`
	PrivateIPAddress       string `json:"privateIpAddress,omitempty"`
	PublicIPAddress        string `json:"publicIpAddress,omitempty"`
	Region                 string `json:"region,omitempty"`
	UserID                 string `json:"userId,omitempty"`
	TeamID                 string `json:"teamId,omitempty"`
	ScriptID               string `json:"scriptId,omitempty"`
	DtLastRun              string `json:"dtLastRun,omitempty"`
}

// CreateMachine creates a machine
func (c Client) CreateMachine(params CreateMachineParams) (Machine, error) {
	machine := Machine{}

	url := fmt.Sprintf("/machines/createSingleMachinePublic")
	_, err := c.Request("POST", url, params, &machine, params.RequestParams)

	return machine, err
}

// StartMachine starts a machine
func (c Client) StartMachine(id string, params StartMachineParams) (Machine, error) {
	machine := Machine{}

	url := fmt.Sprintf("/machines/%s/start", id)
	_, err := c.Request("POST", url, nil, &machine, params.RequestParams)

	return machine, err
}

// StopMachine stops a machine
func (c Client) StopMachine(id string, params StopMachineParams) (Machine, error) {
	machine := Machine{}

	url := fmt.Sprintf("/machines/%s/stop", id)
	_, err := c.Request("POST", url, nil, &machine, params.RequestParams)

	return machine, err
}

// RestartMachine restarts a machine
func (c Client) RestartMachine(id string, params RestartMachineParams) (Machine, error) {
	machine := Machine{}

	url := fmt.Sprintf("/machines/%s/restart", id)
	_, err := c.Request("POST", url, nil, &machine, params.RequestParams)

	return machine, err
}

// DestroyMachine destroys a machine
func (c Client) DestroyMachine(id string, params DestroyMachineParams) (Machine, error) {
	machine := Machine{}

	url := fmt.Sprintf("/machines/%s/destroyMachine", id)
	_, err := c.Request("POST", url, nil, &machine, params.RequestParams)

	return machine, err
}

// UpdateMachine updates a machine
func (c Client) UpdateMachine(id string, params UpdateMachineParams) (Machine, error) {
	machine := Machine{}

	url := fmt.Sprintf("/machines/%s/restart", id)
	_, err := c.Request("POST", url, nil, &machine, params.RequestParams)

	return machine, err
}

// GetMachines gets a list of machines
func (c Client) GetMachines(params GetMachinesParams) ([]Machine, error) {
	machines := []Machine{}

	url := fmt.Sprintf("/machines/getMachines")
	_, err := c.Request("GET", url, params, &machines, params.RequestParams)

	return machines, err
}

// CreateMachine creates a machine
func (c Client) GetMachine(id string, params GetMachineParams) (GetMachineResponse, error) {
	machine := GetMachineResponse{}

	url := fmt.Sprintf("/machines/getMachinePublic?machineId=%s", id)
	_, err := c.Request("GET", url, params, &machine, params.RequestParams)

	return machine, err
}
