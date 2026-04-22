package main

import (
	"flag"
	"fmt"
	"sorcerer.nz/autoctm/internal/cli"
)

func main() {

	initFlag := flag.Bool("init", false, "Initialize first instance")
	flag.Parse()

	// TODO: initialize store + broker implementations
	store := initStore()
	broker := initBroker()

	c := cli.New(store, broker)
	runner := cli.NewRunner(c)

	// Handle init flag (bootstrap only)
	if *initFlag {
		fmt.Println("Initializing first instance...")

		// TODO:
		// - check DB for existing instances
		// - if none exist, create default instance via broker
		// - set context automatically

		id, err := broker.StartInstance()
		if err != nil {
			fmt.Println("failed to initialize instance:", err)
			return
		}

		_, err = c.SetContext(id)
		if err != nil {
			fmt.Println("failed to set context:", err)
			return
		}
	}

	// Start CLI loop 
	runner.Run()

	
}