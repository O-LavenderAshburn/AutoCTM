package cli

import (
	"encoding/json"
	"fmt"
)

type SetContextArgs struct {
	InstanceID string `json:"instanceId"`
}

// Send command to broker
func (c *CLI) Start() error {
	resp, err := c.sendAndWait("start-instance", nil)
	if err != nil {
		return err
	}
	if resp.Error != "" {
		return fmt.Errorf(resp.Error)
	}
	fmt.Println("Instance started:", string(resp.Data))
	return nil
}

// List all registered instances, their IDs, and current status.
func (c *CLI) ListInstances() error {
	response, err := c.sendAndWait("list-instances", nil)
	if err != nil {
		return err
	}

	// Unmarshal the body into a slice
	var instances []InstanceContext
	if err := json.Unmarshal(response.Data, &instances); err != nil {
		return err
	}

	for _, inst := range instances {
		fmt.Printf("ID: %s | Status: %s\n", inst.ID, inst.Status)
	}
	return nil
}

// Set the active instance context for the CLI.
// Future commands will target this instance.
func (c *CLI) SetContext(instanceID string) error {
	data, err := c.sendAndWait("set-context", SetContextArgs{
		InstanceID: instanceID,
	})
	if err != nil {
		return err
	}

	// parse the response into InstanceContext
	var inst InstanceContext
	if err := json.Unmarshal(data.Data, &inst); err != nil {
		return err
	}

	c.context = &InstanceContext{
		ID:        inst.ID,
		Status:    inst.Status,
		StartedAt: inst.StartedAt,
		Active:    true,
	}

	fmt.Println("Active instance:", inst.ID)
	return nil // missing
}

func (c *CLI) ExitContext() {
	c.context = nil
}
