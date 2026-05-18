package cli

import (
	"encoding/binary"
	"encoding/json"
	"io"
	"net"
	"time"

	"sorcerer.nz/autoctm/internal/protocol"
)

func (c *CLI) recv() (protocol.Response, error) {
    // Read length prefix 
	var length uint32
	if err := binary.Read(c.conn, binary.BigEndian, &length); err != nil {
		return protocol.Response{}, err
	}
    // Read the body 
	buf := make([]byte, length)
	if _, err := io.ReadFull(c.conn, buf); err != nil {
		return protocol.Response{}, err
	}

	var resp protocol.Response
	if err := json.Unmarshal(buf, &resp); err != nil {
		return protocol.Response{}, err
	}

	return resp, nil
}

func (c *CLI) send(cmd string, args any) error {
	command := protocol.Command{Cmd: cmd}

	if args != nil {
		encodedArgs, err := json.Marshal(args)
		if err != nil {
			return err
		}
		command.Args = encodedArgs
	}

	data, err := json.Marshal(command)
	if err != nil {
		return err
	}

	length := uint32(len(data))
	if err := binary.Write(c.conn, binary.BigEndian, length); err != nil {
		return err
	}

	_, err = c.conn.Write(data)
	return err
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
	context *InstanceContext
}

func New(conn net.Conn) *CLI {
	return &CLI{
		conn: conn,
	}
}