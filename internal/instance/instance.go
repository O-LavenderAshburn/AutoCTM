package instance

import (
    "time"
    "encoding/json"
)
        
type Instance struct {
    ID        string
    PID       int
    Status    string
    StartedAt time.Time
    Config    json.RawMessage
}

