package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type CLIRunner struct {
	cli *CLI
}

func NewRunner(cli *CLI) *CLIRunner {
	return &CLIRunner{
		cli: cli,
	}
}

func (r *CLIRunner) Run() {
	reader := bufio.NewReader(os.Stdin)

	for {
		r.printPrompt()

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("error reading input:", err)
			continue
		}

		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}

		if input == "exit" || input == "quit" {
			fmt.Println("bye")
			return
		}

		r.handle(input)
	}
}