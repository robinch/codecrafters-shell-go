package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var commands []string

func main() {
	commands = []string{"echo", "exit", "type"}
	br := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$ ")

		input, err := br.ReadString('\n')

		if err != nil {
			panic(fmt.Sprintf("Could not read cmd, error: %v", err))
		}

		input = strings.TrimSpace(input)
		cmd := strings.SplitN(input, " ", 2)

		evalCommand(cmd)
	}
}

func evalCommand(cmd []string) {
	switch cmd[0] {
	case "echo":
		evalEcho(cmd)
	case "exit":
		evalExit()
	case "type":
		evalType(cmd)
	default:
		fmt.Printf("%s: command not found\n", cmd[0])
	}
}

func evalEcho(cmd []string) {
	fmt.Println(cmd[1])
}

func evalExit() {
	os.Exit(0)
}

func evalType(cmd []string) {
	if isCommand(cmd) {
		fmt.Printf("%s is a shell builtin\n", cmd[1])
	} else {
		fmt.Printf("%s: not found\n", cmd[1])
	}
}

func isCommand(cmd []string) bool {
	for _, c := range commands {
		if cmd[1] == c {
			return true
		}
	}

	return false
}
