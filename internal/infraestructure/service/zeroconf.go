package service

import (
	"fmt"
	"os"
	"os/signal"
	"synk/internal/domain"

	"synk/pkg/logger"
	"syscall"

	"github.com/grandcat/zeroconf"
)

type ZeroconfService struct {
	service *zeroconf.Server
	config  domain.Config
}

var service *ZeroconfService = nil

var log = logger.GetLogger("zeroconf")

func NewZeroconfService(config domain.Config) (*ZeroconfService, error) {
	if service != nil {
		return service, nil
	}

	service = &ZeroconfService{config: config}

	return service, nil
}

func (z *ZeroconfService) Start() error {
	server, err := zeroconf.Register(z.config.Service.Name, z.config.Service.Type, z.config.Service.Domain, z.config.Service.Port, nil, nil)
	if err != nil {
		return err
	}

	service.service = server
	log.Info(fmt.Sprintf("Zeroconf service registered: %s | porta: %d", z.config.Service.Name, z.config.Service.Port))

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigChan
		service.Stop()
		os.Exit(0)
	}()

	return nil
}

func (z *ZeroconfService) Stop() error {
	if z.service != nil {
		z.service.Shutdown()
		log.Info(fmt.Sprintf("Zeroconf service unregistered: %s", z.config.Service.Name))
		service = nil
		return nil
	}
	return nil
}
