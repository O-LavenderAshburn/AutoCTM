package cli

import (
	"time"
	"net"
	"encoding/json"
    "encoding/binary"
    "io"
)


type Response struct {
    OK   bool            `json:"ok"`
    Body json.RawMessage `json:"body"`
}

type Command struct {
    Cmd  	string          `json:"cmd"`
    Args 	json.RawMessage `json:"args,omitempty"` // omitted if nil
}

func (c *CLI) recv() (Response, error) {
    // Read length prefix
    var length uint32
    if err := binary.Read(c.conn, binary.BigEndian, &length); err != nil {
        return Response{}, err
    }

    // Read the body
    buf := make([]byte, length)
    if _, err := io.ReadFull(c.conn, buf); err != nil {
        return Response{}, err
    }

    // Unmarshal into Response
    var resp Response
    if err := json.Unmarshal(buf, &resp); err != nil {
        return Response{}, err
    }

    return resp, nil
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