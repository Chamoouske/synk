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

	server, err := zeroconf.Register("Synk-"+config.Service.Name, config.Service.Type, config.Service.Domain, config.Service.Port, nil, nil)
	if err != nil {
		return nil, err
	}

	service = &ZeroconfService{service: server, config: config}
	log.Info(fmt.Sprintf("Zeroconf service registered: %s", config.Service.Name))

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigChan
		service.Unregister()
		os.Exit(0)
	}()

	return service, nil
}

func (z *ZeroconfService) Unregister() error {
	if z.service != nil {
		z.service.Shutdown()
		log.Info(fmt.Sprintf("Zeroconf service unregistered: %s", z.config.Service.Name))
		service = nil
		return nil
	}
	return nil
}
