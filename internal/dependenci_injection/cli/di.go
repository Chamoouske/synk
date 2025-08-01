package cli

import (
	initcmd "synk/internal/infraestructure/command/init"
	"synk/internal/infraestructure/command/watch"
	factory "synk/internal/infraestructure/factory"
)

func InitializeCommandsFactory() *factory.CommandsFactory {
	commandsFactory := factory.NewCommandsFactory()

	commandsFactory.RegisterCommand(initcmd.CommandName, initcmd.NewInitCommand())
	commandsFactory.RegisterCommand(watch.CommandName, watch.NewWatchCommand())

	return commandsFactory
}
