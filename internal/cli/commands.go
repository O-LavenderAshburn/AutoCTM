package cli

import (
    "fmt"
    "encoding/json"
)


type SetContextArgs struct {
	InstanceID string `json:"instanceId"`
}



//Send command to broker
func (c *CLI) Start() error {
    return c.send("start-instance", nil)
}


// List all registered instances, their IDs, and current status.
//TODO: Fix in Specs document
func (c *CLI) ListInstances() error {
    response, err := c.sendAndWait("list-instances", nil)
    if err != nil {
        return err
    }

    // Unmarshal the body into a slice
    var instances []InstanceContext
    if err := json.Unmarshal(response.Body, &instances); err != nil {
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
    if err := json.Unmarshal(data.Body, &inst); err != nil {
        return err
    }

    c.context = &InstanceContext{
        ID:        inst.ID,
        Status:    inst.Status,
        StartedAt: inst.StartedAt,
        Active:    true,
    }

    fmt.Println("Active instance:", inst.ID)
    return nil  // missing
}

func (c *CLI) ExitContext(){
	c.context = nil
}

