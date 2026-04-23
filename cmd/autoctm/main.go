package main

import (
	"flag"
	"fmt"
	"sorcerer.nz/autoctm/internal/broker"
	"sorcerer.nz/autoctm/internal/cli"
)

func main() {
	initFlag := flag.Bool("init", false, "Initialize first instance")
	flag.Parse()

	// TODO: initialize store + broker implementations
	b := broker.New()

	c := cli.New(b)
	runner := cli.NewRunner(c)
	
	if *initFlag {
		fmt.Println("Initializing first instance...")

		// TODO:
		// - check DB for existing instances
		// - if none exist, create default instance via broker
		// - set context automatically
	}
	fmt.Println("starting runner")
	runner.Run()
}