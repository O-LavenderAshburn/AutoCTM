package protocol

type Response struct {
    OK   bool            `json:"ok"`
    Body json.RawMessage `json:"body"`
}

type Command struct {
    Cmd  	string          `json:"cmd"`
    Args 	json.RawMessage `json:"args,omitempty"` // omitted if nil
}

const (
    SocketPath = "/tmp/autoctm/autoctm-broker.sock"
    SocketDir  = "/tmp/autoctm"
)