package cli

import (
	command "synk/internal/infraestructure/command"
	factory "synk/internal/infraestructure/factory"
)

func InitializeCommandsFactory() *factory.CommandsFactory {
	command.InitCommands()

	return factory.NewCommandsFactory()
}
