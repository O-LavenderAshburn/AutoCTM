package cli

import (
	"fmt"
	"time"

	"sorcerer.nz/autoctm/internal/broker"
	"sorcerer.nz/autoctm/internal/instance"
)

type InstanceContext struct {
	ID        string
	Status    string
	StartedAt time.Time
	Active    bool
}

type CLI struct {
	broker  broker.Broker
	context *InstanceContext
}

func (c *CLI) Start() error {
	id, err := c.broker.StartInstance()
	if err != nil {
		return err
	}

	fmt.Println("Started instance:", id)
	return nil
}


func New(b broker.Broker) *CLI {
	return &CLI{
		broker: b,
	}
}