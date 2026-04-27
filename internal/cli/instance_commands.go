package cli

import (
    "fmt"
	"encoding/json"
)

type LogArgs struct {
    InstanceID string `json:"instanceId"`
    URL        string `json:"url"`
}

type InstanceArgs struct {
    InstanceID string `json:"instanceId"`
}

// AddLog adds a CT log to the active instance.
// The broker persists the log and updates the instance config.
func (c *CLI) AddLog(url string) error {
    if c.context == nil {
        return fmt.Errorf("no active instance selected")
    }

    return c.send("add-log", LogArgs{
        InstanceID: c.context.ID,
        URL:        url,
    })
}
// RemoveLog removes a CT log from the active instance.
func (c *CLI) RemoveLog(url string) error {
    if c.context == nil {
        return fmt.Errorf("no active instance selected")
    }

    return c.send("remove-log", LogArgs{
        InstanceID: c.context.ID,
        URL:        url,
    })
}

// Pause stops the active instance’s polling loop.
func (c *CLI) Pause() error {
    if c.context == nil {
        return fmt.Errorf("no active instance selected")
    }
    return c.send("pause", InstanceArgs{
        InstanceID: c.context.ID,
    })
}

// Resume restarts a paused instance.
func (c *CLI) Resume() error {
    if c.context == nil {
        return fmt.Errorf("no active instance selected")
    }

    return c.send("resume", InstanceArgs{
        InstanceID: c.context.ID,
    })
}

// Status prints the current state of the active instance.
func (c *CLI) Status() error {
    if c.context == nil {
        return fmt.Errorf("no active instance selected")
    }

    resp, err := c.sendAndWait("status", InstanceArgs{
        InstanceID: c.context.ID,
    })
    if err != nil {
        return err
    }

    var inst InstanceContext
    if err := json.Unmarshal(resp.Body, &inst); err != nil {
        return err
    }

    fmt.Printf(
        "Instance: %s | Status: %s\n",
        inst.ID,
        inst.Status,
    )
    return nil
}

// Shutdown stops the active instance.
func (c *CLI) Shutdown() error {
    if c.context == nil {
        return fmt.Errorf("no active instance selected")
    }

    return c.send("shutdown", InstanceArgs{
        InstanceID: c.context.ID,
    })
}