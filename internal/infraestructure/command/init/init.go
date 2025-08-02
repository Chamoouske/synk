package init

import (
	"fmt"
	"os"
	"synk/config"
	"synk/internal/domain"
	"synk/internal/infraestructure/factory"
	"synk/internal/infraestructure/pem"
	"synk/internal/infraestructure/service"
	"synk/pkg/logger"
)

const CommandName = "init"

var log = logger.GetLogger("init")

type InitCommand struct {
	service domain.Service
}

func Init() {
	commandsFactory := factory.NewCommandsFactory()
	commandsFactory.RegisterCommand(CommandName, NewInitCommand())
}

func NewInitCommand() *InitCommand {
	return &InitCommand{}
}

func (c *InitCommand) Execute(args []string) error {
	device, err := c.createKeys()
	if err != nil {
		log.Error("Erro ao criar chaves: " + err.Error())
		os.Exit(1)
	}

	err = c.registerService(*device)
	if err != nil {
		log.Error("Erro ao registrar serviço: " + err.Error())
		os.Exit(1)
	}

	c.service.Start()
	service.StartTCPServer(c.service.GetPort())

	defer c.service.Stop()

	select {}
}

func (c *InitCommand) createKeys() (*domain.Device, error) {
	device := config.GetDevice()
	if device != nil {
		log.Info("Dispositivo carregado: " + device.ID)
		return device, nil
	}
	privateKeyPEM, publicKeyPEM, err := pem.GenerateKeys()
	if err != nil {
		return nil, fmt.Errorf("erro ao gerar chaves: %w", err)
	}

	err = pem.SaveKeys(privateKeyPEM, publicKeyPEM)
	if err != nil {
		return nil, fmt.Errorf("erro ao salvar chaves: %w", err)
	}

	device, err = config.SaveDevice(string(privateKeyPEM), string(publicKeyPEM))

	log.Info("Dispositivo criado: " + device.ID)

	if err != nil {
		return nil, fmt.Errorf("erro ao salvar dispositivo: %w", err)
	}
	return device, nil
}

func (c *InitCommand) registerService(device domain.Device) error {
	config, err := config.GetConfigServer()
	if err != nil {
		return fmt.Errorf("erro ao obter configuração do servidor: %w", err)
	}

	config.Service.Name = "Synk-" + device.ID
	server, err := service.NewZeroconfService(config, &device)
	if err != nil {
		return fmt.Errorf("erro ao registrar serviço Zeroconf: %w", err)
	}

	c.service = server
	return nil
}
