package cli

import("fmt")

// AddLog adds a CT log to the instance’s configuration.
// The broker persists the log and notifies the instance to reload config.
func (c *CLI) AddLog(url string) error {
	if c.context == nil {
		return fmt.Errorf("no active instance selected")
	}

	return c.broker.AddLog(c.context.ID, url)
}


// RemoveLog removes a CT log from the instance’s configuration.
// The broker updates the database and notifies the instance to reload config.
func (c *CLI) RemoveLog(url string) error {
	if c.context == nil {
		return fmt.Errorf("no active instance selected")
	}

	return c.broker.RemoveLog(c.context.ID, url)
}

//Pause the instance’s polling loop.
func (c *CLI) Pause() error {
	if c.context == nil {
		return fmt.Errorf("no active instance selected")
	}

	return c.broker.Pause(c.context.ID)
}

//Resume a paused instance.
func (c *CLI) Resume() error {
	if c.context == nil {
		return fmt.Errorf("no active instance selected")
	}

	return c.broker.Resume(c.context.ID)
}

//Print the instance’s current status, including configured logs, last known 
//STH per log, and alert counts
func (c *CLI) Status() error {
	if c.context == nil {
		return fmt.Errorf("no active instance selected")
	}

	inst, err := c.broker.Status(c.context.ID)
	if err != nil {
		return err
	}

	fmt.Printf("Instance: %s | Status: %s\n", inst.ID, inst.Status)
	return nil
}

//Shutdown the active instance only.
func (c *CLI) Shutdown() error {
	if c.context == nil {
		return fmt.Errorf("no active instance selected")
	}

	err := c.broker.Shutdown(c.context.ID)
	if err != nil {
		return err
	}

	fmt.Println("Instance shutdown:", c.context.ID)
	return nil
}






