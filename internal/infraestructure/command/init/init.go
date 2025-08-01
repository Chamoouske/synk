package init

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"os"
	"synk/internal/domain"
	"synk/internal/infraestructure/pem"
	"synk/internal/infraestructure/service"
	"synk/pkg/logger"

	"gopkg.in/gcfg.v1"
)

const CommandName = "init"

var log = logger.GetLogger("init")

type InitCommand struct {
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
	return nil
}

func (c *InitCommand) createKeys() (*domain.Device, error) {
	device := c.LoadDevice()
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
	device = &domain.Device{ID: generateRandomID(), PublicKey: string(publicKeyPEM), PrivateKey: string(privateKeyPEM)}

	err = c.SaveDevice(device)

	log.Info("Dispositivo criado: " + device.ID)

	if err != nil {
		return nil, fmt.Errorf("erro ao salvar dispositivo: %w", err)
	}
	return device, nil
}

func (c *InitCommand) registerService(device domain.Device) error {
	config, err := getConfigServer()
	if err != nil {
		return fmt.Errorf("erro ao obter configuração do servidor: %w", err)
	}

	config.Service.Name = device.ID
	server, err := service.NewZeroconfService(config)
	if err != nil {
		return fmt.Errorf("erro ao registrar serviço Zeroconf: %w", err)
	}
	log.Info("Serviço Zeroconf registrado com sucesso.")
	defer server.Unregister()
	// return nil
	select {}
}

func (c *InitCommand) UnregisterService() error {
	return nil
}

func getConfigServer() (domain.Config, error) {
	var config domain.Config
	err := gcfg.ReadFileInto(&config, "config/service.cfg")
	if err != nil {
		return config, fmt.Errorf("falha ao ler o arquivo de configuração: %v", err)
	}
	return config, nil
}

func (c *InitCommand) SaveDevice(device *domain.Device) error {
	file, err := os.Create(".synk/device.json")
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)

	if err := encoder.Encode(device); err != nil {
		return fmt.Errorf("failed to encode device to JSON: %w", err)
	}

	return nil
}

func (c *InitCommand) LoadDevice() *domain.Device {
	device := &domain.Device{}
	file, err := os.ReadFile(".synk/device.json")
	if err != nil {
		return nil
	}
	err = json.Unmarshal(file, device)
	if err != nil {
		return nil
	}
	if device.ID == "" || device.PublicKey == "" || device.PrivateKey == "" {
		return nil
	}
	return device
}

func generateRandomID() string {
	b := make([]byte, 8)
	_, err := rand.Read(b)
	if err != nil {
		log.Error("Failed to generate random ID", "error", err)
	}
	return fmt.Sprintf("%x", b)
}
