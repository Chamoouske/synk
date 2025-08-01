package factory

import (
	"fmt"
	"synk/internal/domain"
)

var factoryInstance *CommandsFactory = nil

type CommandsFactory struct {
	commands map[string]domain.Command
}

func NewCommandsFactory() *CommandsFactory {
	if factoryInstance != nil {
		return factoryInstance
	}
	factoryInstance = &CommandsFactory{
		commands: make(map[string]domain.Command),
	}
	return factoryInstance
}

func (f *CommandsFactory) RegisterCommand(name string, command domain.Command) {
	f.commands[name] = command
}

func (f *CommandsFactory) ExecuteCommand(name string, args []string) error {
	command, exists := f.commands[name]
	if !exists {
		return fmt.Errorf("command %s not found", name)
	}
	return command.Execute(args)
}
