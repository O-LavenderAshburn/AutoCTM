package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"sorcerer.nz/autoctm/internal/broker"
	"sorcerer.nz/autoctm/internal/protocol"
)

// handleSignals waits for a termination signal and cleans up before exit.
// Closing the listener causes Accept() in the main loop to return an error,
// which exits the broker gracefully.
func handleSignals(listener net.Listener) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	os.Remove(protocol.SocketPath)
	listener.Close()
}

// Dispatch commands received.
func dispatch(req protocol.Command, b broker.Broker) protocol.Response {
	switch req.Cmd {
	case "start-instance":
		id, err := b.StartInstance()
		return respond(id, err)
	case "stop":
		return respond(nil, b.StopInstance(req.ID))
	case "pause":
		return respond(nil, b.Pause(req.ID))
	case "resume":
		return respond(nil, b.Resume(req.ID))
	case "list":
		instances, err := b.ListInstances()
		return respond(instances, err)
	case "get":
		inst, err := b.GetInstance(req.ID)
		return respond(inst, err)
	case "add-log":
		return respond(nil, b.AddLog(req.ID, req.URL))
	case "remove-log":
		return respond(nil, b.RemoveLog(req.ID, req.URL))
	default:
		return protocol.Response{Error: fmt.Sprintf("unknown command: %s", req.Cmd)}
	}
}

func handleConn(conn net.Conn, b broker.Broker) {
	defer conn.Close()

	dec := json.NewDecoder(conn)
	enc := json.NewEncoder(conn)

	for {
		var req protocol.Command
		if err := dec.Decode(&req); err != nil {
			enc.Encode(protocol.Response{Error: fmt.Sprintf("decode error: %v", err)})
			return
		}
		resp := dispatch(req, b)
		if err := enc.Encode(resp); err != nil {
			return
		}
	}
}

func respond(data any, err error) protocol.Response {
	if err != nil {
		return protocol.Response{Error: err.Error()}
	}
	if data == nil {
		return protocol.Response{}
	}
	raw, err := json.Marshal(data)
	if err != nil {
		return protocol.Response{Error: err.Error()}
	}
	return protocol.Response{Data: raw}
}

func main() {

	//Setup fresh socket and listen.
	fmt.Println("Starting socket")
	os.MkdirAll(protocol.SocketDir, 0755)
	os.Remove(protocol.SocketPath)

	listener, err := net.Listen("unix", protocol.SocketPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to listen: %v\n", err)
		os.Exit(1)
	} else {
		fmt.Printf("Socket listening on: %s\n", protocol.SocketPath)
	}

	b := broker.New()

	//Handle termination signals
	go handleSignals(listener)

	//Accept connection.
	for {
		conn, err := listener.Accept()
		if err != nil {
			return
		}
		fmt.Println("ACCEPTED connection from:", conn.RemoteAddr())
		go handleConn(conn, b)
	}
}
