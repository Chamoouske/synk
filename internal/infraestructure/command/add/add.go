package add

import (
	"fmt"
	"synk/config"
	"synk/internal/domain"
	"synk/internal/infraestructure/factory"
	"synk/internal/infraestructure/service"
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

	cfg, err := config.GetConfigServer()
	if err != nil {
		return fmt.Errorf("erro ao obter configuração do servidor: %w", err)
	}

	deviceID := args[0]
	log.Info("Adicionando dispositivo: " + deviceID)
	device := config.GetDevice()
	if device == nil {
		return fmt.Errorf("Dispositivo não encontrado, execute 'synk init' primeiro")
	}

	serv, err := service.NewZeroconfService(cfg, device)
	if err != nil {
		return fmt.Errorf("erro ao criar serviço Zeroconf: %w", err)
	}

	c.service = serv
	err = c.service.AddDeviceToConnect(deviceID)

	return nil
}
