package init

import (
	"fmt"
)

const CommandName = "init"

type InitCommand struct{}

func NewInitCommand() *InitCommand {
	return &InitCommand{}
}

func (c *InitCommand) Execute(args []string) error {
	fmt.Println("Initializing the application...")
	return nil
}
