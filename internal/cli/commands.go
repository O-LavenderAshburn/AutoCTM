package cli

import (
	"fmt"
)



// List all registered instances, their IDs, and current status.
//TODO: Fix in Specs document
func (c *CLI) ListInstances() error {
	instances, err := c.broker.ListInstances()
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
func (c *CLI) SetContext(instanceID string) {
	inst, err := c.broker.GetInstance(instanceID)
	if err != nil {
		return
	}

	c.context = &InstanceContext{
		ID:        inst.ID,
		Status:    inst.Status,
		StartedAt: inst.StartedAt,
		Active:    true,
	}

	fmt.Println("Active instance:", inst.ID)
}

func (c *CLI) ExitContext() {
	c.context = nil
	fmt.Println("Returned to global context")
}

