package config

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"os"
	"synk/internal/domain"
	"synk/pkg/logger"

	"gopkg.in/gcfg.v1"
)

var log = logger.GetLogger("config")

func GetConfigServer() (domain.Config, error) {
	var config domain.Config
	err := gcfg.ReadFileInto(&config, "config/service.cfg")
	if err != nil {
		return config, fmt.Errorf("falha ao ler o arquivo de configuração: %v", err)
	}
	return config, nil
}

func GetDevice() *domain.Device {
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

func SaveDevice(privateKeyPEM string, publicKeyPEM string) (*domain.Device, error) {
	device := &domain.Device{ID: generateRandomID(), PublicKey: string(publicKeyPEM), PrivateKey: string(privateKeyPEM)}
	file, err := os.Create(".synk/device.json")
	if err != nil {
		return nil, fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)

	if err := encoder.Encode(device); err != nil {
		return nil, fmt.Errorf("failed to encode device to JSON: %w", err)
	}

	return device, nil
}

func generateRandomID() string {
	b := make([]byte, 8)
	_, err := rand.Read(b)
	if err != nil {
		log.Error("Failed to generate random ID", "error", err)
	}
	return fmt.Sprintf("%x", b)
}
