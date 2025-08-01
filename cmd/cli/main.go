package main

import (
	"fmt"
	"os"

	dependency_injection "synk/internal/dependency_injection/cli"
)

func main() {
	os.MkdirAll(".synk", os.ModePerm)

	if len(os.Args) < 2 {
		fmt.Println("Usage: cli <command>")
		os.Exit(1)
	}
	command := os.Args[1]
	commandsFactory := dependency_injection.InitializeCommandsFactory()

	commandsFactory.ExecuteCommand(command, os.Args[2:])

	os.Exit(0)
}
