package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var commands []string
var paths []string

func main() {
	commands = []string{"echo", "exit", "type", "pwd"}
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
	if cmd[0] == "echo" {
		evalEcho(cmd)
	} else if cmd[0] == "exit" {
		evalExit()
	} else if cmd[0] == "type" {
		evalType(cmd)
	} else if cmd[0] == "pwd" {
		evalPwd()
	} else if cmd[0] == "cd" {
		evalCd(cmd)
	} else if filepath, exists := isCommandFromPath(cmd[0]); exists {
		runCommandFromPath(filepath, cmd[1])
	} else {
		fmt.Printf("%s: command not found\n", cmd[0])
	}
}

func evalEcho(cmd []string) {
	fmt.Println(cmd[1])
}

func evalExit() {
	os.Exit(0)
}

func evalPwd() {
	path, err := os.Getwd()

	if err != nil {
		panic(fmt.Sprintf("Could not get working directory, err: %s", err))
	}

	fmt.Println(path)
}

func evalCd(cmd []string) {
	err := os.Chdir(cmd[1])

	if err != nil {
		fmt.Printf("cd: %s: No such file or directory\n", cmd[1])
	}
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

func runCommandFromPath(filepath string, args string) {
	command := exec.Command(filepath, args)
	var out strings.Builder
	command.Stdout = &out
	err := command.Run()

	if err != nil {
		panic(fmt.Sprintf("could not run command from path, error: %s", err))
	}

	fmt.Print(out.String())
}
