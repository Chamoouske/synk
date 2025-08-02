package watch

import (
	"fmt"
	"synk/internal/infraestructure/command"
	"synk/internal/infraestructure/factory"
)

const CommandName = "watch"

type WatchCommand struct {
}

func (w *WatchCommand) Init(factory *factory.CommandsFactory) {
	factory.RegisterCommand(CommandName, w)
}

func init() {
	command.RegisterCommand(&WatchCommand{})
}

func (c *WatchCommand) Execute(args []string) error {
	if len(args) == 0 {
		fmt.Println("No arguments provided for watch command.")
		return nil
	}
	fmt.Println("Watching for changes in:", args[0])
	return nil
}
