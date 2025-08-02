package config

import (
	"encoding/json"
	"fmt"
	"os"
	"synk/internal/domain"

	"gopkg.in/gcfg.v1"
)

type SynkConfig struct {
	config domain.SynkConfig
}

var synkConfig *SynkConfig = &SynkConfig{}

func init() {
	err := gcfg.ReadFileInto(&synkConfig.config, "config/service.cfg")
	if err != nil {
		log.Error(fmt.Sprintf("falha ao ler o arquivo de configuração: %v", err))
		os.Exit(1)
	}
}

func GetSynkConfig() *SynkConfig {
	return synkConfig
}

func SaveSynkConfig(synk *SynkConfig) error {
	file, err := os.Create(".synk/synk.json")
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)

	if err := encoder.Encode(synk); err != nil {
		return fmt.Errorf("failed to encode device to JSON: %w", err)
	}

	return nil
}
