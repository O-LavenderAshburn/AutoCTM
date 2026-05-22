package protocol

import "encoding/json"

type Response struct {
	Data  json.RawMessage `json:"data,omitempty"`
	Error string          `json:"error,omitempty"`
}

type Command struct {
	Cmd  string          `json:"cmd"`
	Args json.RawMessage `json:"args,omitempty"`
	ID   string          `json:"id,omitempty"`
	URL  string          `json:"url,omitempty"`
}

const (
	SocketPath = "/tmp/autoctm/autoctm-broker.sock"
	SocketDir  = "/tmp/autoctm"
)
