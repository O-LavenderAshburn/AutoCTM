package main

import (
	"flag"
)

func main() {
	initFlag := flag.Bool("init", false, "Initialize first instance")

	flag.Parse()

	if *initFlag {
		// TODO: Run instace initialization verification.
		// TODO: If no instace, initilize intance.
	}

}
