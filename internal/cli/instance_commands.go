package cli

import (
	"fmt"

)

// AddLog adds a CT log to the active instance.
// The broker persists the log and updates the instance config.
func (c *CLI) AddLog(url string) error {
	if c.context == nil {
		return fmt.Errorf("no active instance selected")
	}

	return c.broker.AddLog(c.context.ID, url)
}

// RemoveLog removes a CT log from the active instance.
func (c *CLI) RemoveLog(url string) error {
	if c.context == nil {
		return fmt.Errorf("no active instance selected")
	}

	return c.broker.RemoveLog(c.context.ID, url)
}

// Pause stops the active instance’s polling loop.
func (c *CLI) Pause() error {
	if c.context == nil {
		return fmt.Errorf("no active instance selected")
	}

	return c.broker.Pause(c.context.ID)
}

// Resume restarts a paused instance.
func (c *CLI) Resume() error {
	if c.context == nil {
		return fmt.Errorf("no active instance selected")
	}

	return c.broker.Resume(c.context.ID)
}

// Status prints the current state of the active instance.
func (c *CLI) Status() error {
	if c.context == nil {
		return fmt.Errorf("no active instance selected")
	}

	inst, err := c.broker.GetInstance(c.context.ID)
	if err != nil {
		return err
	}

	fmt.Printf(
		"Instance: %s | Status: %s | PID: %d\n",
		inst.ID,
		inst.Status,
		inst.PID,
	)

	return nil
}

// Shutdown stops the active instance.
func (c *CLI) Shutdown() error {
	if c.context == nil {
		return fmt.Errorf("no active instance selected")
	}

	err := c.broker.StopInstance(c.context.ID)
	if err != nil {
		return err
	}

	fmt.Println("Instance shutdown:", c.context.ID)

	// optional: clear context after shutdown
	c.context = nil

	return nil
}