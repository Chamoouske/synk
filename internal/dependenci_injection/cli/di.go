package cli

import (
	initcmd "synk/internal/infraestructure/command/init"
	"synk/internal/infraestructure/command/watch"
	factory "synk/internal/infraestructure/factory"
	notifyer "synk/internal/infraestructure/notifyer/cli"
)

func InitializeCommandsFactory() *factory.CommandsFactory {
	commandsFactory := factory.NewCommandsFactory()

	commandsFactory.RegisterCommand(initcmd.CommandName, initcmd.NewInitCommand(&notifyer.CliNotifyer{}))
	commandsFactory.RegisterCommand(watch.CommandName, watch.NewWatchCommand(&notifyer.CliNotifyer{}))

	return commandsFactory
}
