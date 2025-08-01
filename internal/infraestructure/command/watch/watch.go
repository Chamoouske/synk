package watch

import (
	"fmt"
)

const CommandName = "watch"

type WatchCommand struct {
}

func NewWatchCommand() *WatchCommand {
	return &WatchCommand{}
}

func (c *WatchCommand) Execute(args []string) error {
	if len(args) == 0 {
		fmt.Println("No arguments provided for watch command.")
		return nil
	}
	fmt.Println("Watching for changes in:", args[0])
	return nil
}
