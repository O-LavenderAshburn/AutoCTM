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
func (r *CLIRunner) handle(input string) {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return
	}

	// Execute commands on the CLI or print errors.
	switch parts[0] {
	case "/start":
		err := r.cli.Start()
		if err != nil {
			fmt.Println("error:", err)
		}

	case "/list":
		err := r.cli.ListInstances()
		if err != nil {
			fmt.Println("error:", err)
		}

	}

}

func (r *CLIRunner) Run() {
	reader := bufio.NewReader(os.Stdin)

	for {

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("error reading input:", err)
			continue
		}

		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}

		if input == "quit" {
			fmt.Println("bye")
			return
		}

		r.handle(input)
	}
}