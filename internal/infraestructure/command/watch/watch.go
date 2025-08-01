package watch

import (
	"fmt"
	"synk/internal/domain"
)

const CommandName = "watch"

type WatchCommand struct {
	notifyer domain.Notifyer
}

func NewWatchCommand(notifyer domain.Notifyer) *WatchCommand {
	return &WatchCommand{notifyer: notifyer}
}

func (c *WatchCommand) Execute(args []string) error {
	if len(args) == 0 {
		fmt.Println("No arguments provided for watch command.")
		return nil
	}
	fmt.Println("Watching for changes in:", args[0])
	return nil
}
