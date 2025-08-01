package main

import (
	"fmt"
	"os"

	dependency_injection "synk/internal/dependenci_injection/cli"
)

func main() {
	os.MkdirAll(".synk", os.ModePerm)
	if len(os.Args) < 2 {
		fmt.Println("Usage: cli <command>")
		os.Exit(1)
	}
	command := os.Args[1]
	commandsFactory := dependency_injection.InitializeCommandsFactory()

	err := commandsFactory.ExecuteCommand(command, os.Args[2:])
	if err != nil {
		fmt.Printf("Error executing command '%s': %v\n", command, err)
		os.Exit(1)
	}

	os.Exit(0)
}
