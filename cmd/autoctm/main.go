package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"sorcerer.nz/autoctm/internal/cli"
)

const BrokerSocket = "/tmp/autoctm/autoctm-broker.sock"

// Attempt conection to broker
func attemptConn() (net.Conn, error) {
    return net.Dial("unix", BrokerSocket)
}

// Initilize Broker
func initBroker(){
	fmt.Printf("Broker Init")
}

func main() {
	initFlag := flag.Bool("init", false, "Initialize first instance")
	flag.Parse()

	// If its the first time running the program
	if *initFlag {
		fmt.Println("Initializing first instance...")
		//TODO: Init Broker
		initBroker()
		// TODO:
		// - check DB for existing instances
		// - if none exist, create default instance via broker
		// - set context automatically
	}

	// Attempt to connect to broker
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