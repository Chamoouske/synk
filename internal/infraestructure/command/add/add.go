package add

import (
	"synk/internal/domain"
	"synk/internal/infraestructure/factory"
	"synk/pkg/logger"
)

const CommandName = "add"

var log = logger.GetLogger("add")

type AddCommand struct {
	service domain.Service
}

func Init() {
	commandsFactory := factory.NewCommandsFactory()
	commandsFactory.RegisterCommand(CommandName, NewAddCommand())
}

func NewAddCommand() *AddCommand {
	return &AddCommand{}
}

func (c *AddCommand) Execute(args []string) error {
	if len(args) < 1 {
		log.Error("Nenhum dispositivo especificado para adicionar.")
		return nil
	}

	deviceID := args[0]
	log.Info("Adicionando dispositivo: " + deviceID)

	return nil
}
