package cli

import (
	"fmt"
	"time"
	"net"
	"encoding/json"
)

type Command struct {
    Cmd  	string          `json:"cmd"`
    Args 	json.RawMessage `json:"args,omitempty"` // omitted if nil
}

// Send a command.
func (c *CLI) send(cmd string, args any) error {

	//Encode command and args.
	command := Command{Cmd:cmd}

	if args != nil {
		encodedArgs, err := json.Marshal(args)
		if err != nil {
            return err
        }
		command.Args = encodedArgs
	}
	data,err := json.Marshal(command)

    if err != nil {
        return err
    }

	//Send command over socket.
    length := uint32(len(data))
    if err := binary.Write(c.conn, binary.BigEndian, length); err != nil {
        return err
    }

    _, err = c.conn.Write(data)
    return err
}

func (c *CLI) sendAndWait(cmd string, args any) (Response, error) {
    if err := c.send(cmd, args); err != nil {
        return Response{}, err
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