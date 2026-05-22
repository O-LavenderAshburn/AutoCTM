package cli

import (
	"encoding/json"
	"net"
	"time"

	"sorcerer.nz/autoctm/internal/protocol"
)

func (c *CLI) send(cmd string, args any) error {
	command := protocol.Command{Cmd: cmd}
	if args != nil {
		encodedArgs, err := json.Marshal(args)
		if err != nil {
			return err
		}
		command.Args = encodedArgs
	}
	return c.enc.Encode(command)
}

func (c *CLI) recv() (protocol.Response, error) {
	var resp protocol.Response
	if err := c.dec.Decode(&resp); err != nil {
		return protocol.Response{}, err
	}
	return resp, nil
}

func (c *CLI) sendAndWait(cmd string, args any) (protocol.Response, error) {
	if err := c.send(cmd, args); err != nil {
		return protocol.Response{}, err
	}
	return c.recv()
}

type InstanceContext struct {
	ID        string
	Status    string
	StartedAt time.Time
	Active    bool
}

type CLI struct {
	conn    net.Conn
	enc     *json.Encoder
	dec     *json.Decoder
	context *InstanceContext
}

func New(conn net.Conn) *CLI {
	return &CLI{
		conn: conn,
		enc:  json.NewEncoder(conn),
		dec:  json.NewDecoder(conn),
	}
}
