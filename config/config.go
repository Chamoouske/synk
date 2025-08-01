package config

import (
	"fmt"
	"synk/internal/domain"

	"gopkg.in/gcfg.v1"
)

func GetConfigServer() (domain.Config, error) {
	var config domain.Config
	err := gcfg.ReadFileInto(&config, "config/service.cfg")
	if err != nil {
		return config, fmt.Errorf("falha ao ler o arquivo de configuração: %v", err)
	}
	return config, nil
}
