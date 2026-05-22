package main

import (
	"fmt"
	"net"
	"os"

	"sorcerer.nz/autoctm/internal/cli"
	"sorcerer.nz/autoctm/internal/protocol"
)

// Attempt conection to broker
func attemptConn() (net.Conn, error) {
	return net.Dial("unix", protocol.SocketPath)
}

func main() {

	conn, err := attemptConn()
	if err != nil {
		fmt.Printf("No broker")
		os.Exit(1)
	}

	c := cli.New(conn)
	runner := cli.NewRunner(c)

	fmt.Println("starting runner")
	runner.Run()
}
