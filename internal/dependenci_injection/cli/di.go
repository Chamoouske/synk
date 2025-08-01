package cli

import (
	initcmd "synk/internal/infraestructure/command/init"
	"synk/internal/infraestructure/command/watch"
	factory "synk/internal/infraestructure/factory"
	"synk/internal/infraestructure/notifyer/cli"
)

func InitializeCommandsFactory() *factory.CommandsFactory {
	commandsFactory := factory.NewCommandsFactory()

	commandsFactory.RegisterCommand(initcmd.CommandName, initcmd.NewInitCommand(&cli.CliNotifyer{}))
	commandsFactory.RegisterCommand(watch.CommandName, watch.NewWatchCommand(&cli.CliNotifyer{}))

	return commandsFactory
}
