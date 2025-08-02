package command

import (
	"synk/internal/infraestructure/factory"
)

var commands = []CommandInitializer{}

type CommandInitializer interface {
	Init(factory *factory.CommandsFactory)
}

func InitCommands() {
	factoryCommands := factory.NewCommandsFactory()
	for _, cmd := range commands {
		cmd.Init(factoryCommands)
	}
}

func RegisterCommand(initializer CommandInitializer) {
	commands = append(commands, initializer)
}
