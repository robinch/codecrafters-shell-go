package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	br := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$ ")

		cmd, err := br.ReadString('\n')
		cmd = strings.TrimSpace(cmd)

		if err != nil {
			panic(fmt.Sprintf("Could not read cmd, error: %v", err))
		}

		evalCommand(cmd)
	}
}

func evalCommand(cmd string) {
	if strings.HasPrefix(cmd, "exit ") {
		os.Exit(0)
	} else {
		fmt.Printf("%s: command not found\n", cmd)
	}
}
