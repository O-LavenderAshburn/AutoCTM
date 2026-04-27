package main

import (
	"os"
	"net"
	"sorcerer.nz/autoctm/internal/protocol"
	"sorcerer.nz/autoctm/internal/broker"
)


//handleSignals waits for a termination signal and cleans up before exit.
//Closing the listener causes Accept() in the main loop to return an error,
//which exits the broker gracefully.
func handleSignals(listener net.Listener) {
    sig := make(chan os.Signal, 1)
    signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
    <-sig
    os.Remove(BrokerSocket)
    listener.Close()
}

//Dispatch commands received.
func dispatch(req Request, b broker.Broker) Response {
	switch req.Command {
	case "start":
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
		return Response{Error: fmt.Sprintf("unknown command: %s", req.Command)}
	}
}


func handleConn(conn net.Conn, b broker.Broker) {
	defer conn.Close()

	dec := json.NewDecoder(conn)
	enc := json.NewEncoder(conn)

	for {
		var req Request
		if err := dec.Decode(&req); err != nil {
			return // client disconnected
		}

		resp := dispatch(req, b)
		enc.Encode(resp)
	}
}

func respond(data any, err error) Response {
	if err != nil {
		return Response{Error: err.Error()}
	}
	return Response{Data: data}
}


func main(){

	//Setup freash socket and listen.
	os.MkdirAll(SocketDir, 0755)
	os.Remove(SocketPath)
	
	listner,err := net.Listen("unix", SocketPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to listen: %v\n", err)
		os.Exit(1)
	}

	b := broker.New()

	//Handle termination signals
	go handleSignal(listner)

	//Accept connection.
	for {
		conn, err := listener.Accept()
		if err != nil {
			return
		}
		go handleConn(conn, b)
	}
}