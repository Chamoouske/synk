package cli

import (
	addCmd "synk/internal/infraestructure/command/add"
	initCmd "synk/internal/infraestructure/command/init"
	watchCmd "synk/internal/infraestructure/command/watch"
	factory "synk/internal/infraestructure/factory"
)

func InitializeCommandsFactory() *factory.CommandsFactory {
	initCmd.Init()
	addCmd.Init()
	watchCmd.Init()

	return factory.NewCommandsFactory()
}
