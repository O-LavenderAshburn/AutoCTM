package cli

import (
	"fmt"
	"time"
)

// Instruct the system to spawn a new monitor instance.
// (Later: this should call broker instead of store)
func (c *CLI) Start() error {
	id, err := c.broker.StartInstance()
	if err != nil {
		return err
	}

	fmt.Println("Started instance:", id)
	return nil
}

// List all registered instances, their IDs, and current status.
func (c *CLI) ListInstances() error {
	instances, err := c.store.List()
	if err != nil {
		return err
	}

	for _, inst := range instances {
		fmt.Printf(
			"ID: %s | Status: %s | PID: %d\n",
			inst.ID,
			inst.Status,
			inst.PID,
		)
	}

	return nil
}


// Set the active instance context for the CLI.
// Future commands will target this instance.
func (c *CLI) SetContext(instanceID string) (*InstanceContext, error) {

	inst, err := c.store.GetByID(instanceID)
	if err != nil {
		return nil, err
	}

	c.context = &InstanceContext{
		ID:        inst.ID,
		Status:    inst.Status,
		StartedAt: inst.StartedAt,
		Active:    true,
	}

	fmt.Println("Active instance:", inst.ID)

	return c.context, nil
}