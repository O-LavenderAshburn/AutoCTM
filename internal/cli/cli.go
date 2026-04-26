package cli

import (
	"fmt"
	"time"
	"net"
	"sorcerer.nz/autoctm/internal/broker"
)

type InstanceContext struct {
	ID        string
	Status    string
	StartedAt time.Time
	Active    bool
}

type CLI struct {
    conn    net.Conn
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


func New(conn net.Conn) *CLI {
    return &CLI{
        conn: conn,
    }
}