package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("$ ")

	br := bufio.NewReader(os.Stdin)
	input, err := br.ReadString('\n')

	input = strings.TrimSpace(input)

	if err != nil {
		panic(fmt.Sprintf("Could not read input, error: %v", err))
	}

	fmt.Printf("%s: command not found\n", input)
}
