package paperspace

import (
	"fmt"
	"time"
)

// Script is the structure of a script
type Script struct {
	ID          string    `json:"id"`
	OwnerType   string    `json:"ownerType"`
	OwnerID     string    `json:"ownerId"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	DtCreated   time.Time `json:"dtCreated"`
	IsEnabled   bool      `json:"isEnabled"`
	RunOnce     bool      `json:"runOnce"`
}

// CreateScriptParams are the parameters for creating a script
type CreateScriptParams struct {
	RequestParams
	ScriptName        string `json:"scriptName"`
	ScriptFile        string `json:"scriptFile,omitempty"`
	ScriptText        string `json:"scriptText,omitempty"`
	ScriptDescription string `json:"scriptDescription,omitempty"`
	IsEnabled         bool   `json:"isEnabled,omitempty"`
	RunOnce           bool   `json:"runOnce,omitempty"`
	MachineID         string `json:"machineId,omitempty"`
}

// DestroyScriptParams are the parameters for destroying a script
type DestroyScriptParams struct {
	RequestParams
	ScriptID string `json:"machineId,omitempty"`
}

// CreateScript creates a script
func (c Client) CreateScript(params CreateScriptParams) (Script, error) {
	script := Script{}

	url := "/scripts/createScript"
	_, err := c.Request("POST", url, params, &script, params.RequestParams)

	return script, err
}

// DestroyScript destroys a script
func (c Client) DestroyScript(params DestroyScriptParams) (Script, error) {
	script := Script{}

	url := fmt.Sprintf("/scripts/%s/destroy", params.ScriptID)
	_, err := c.Request("POST", url, nil, &script, params.RequestParams)

	return script, err
}
