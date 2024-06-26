package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var commands []string
var paths []string

func main() {
	commands = []string{"echo", "exit", "type"}
	paths = strings.Split(os.Getenv("PATH"), ":")

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
	if isBuiltin(cmd[1]) {
		fmt.Printf("%s is a shell builtin\n", cmd[1])
	} else if filepath, exists := isCommandFromPath(cmd[1]); exists {
		fmt.Printf("%s is %s\n", cmd[1], filepath)
	} else {
		fmt.Printf("%s: not found\n", cmd[1])
	}
}

func isBuiltin(c string) bool {
	for _, cm := range commands {
		if c == cm {
			return true
		}
	}

	return false
}

func isCommandFromPath(c string) (string, bool) {
	for _, path := range paths {
		files, _ := os.ReadDir(path)

		for _, file := range files {
			if !file.IsDir() && file.Name() == c {
				return fmt.Sprintf("%s/%s", path, file.Name()), true
			}
		}
	}

	return "", false
}
