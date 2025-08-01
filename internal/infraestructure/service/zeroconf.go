package service

import (
	"synk/internal/domain"

	"github.com/grandcat/zeroconf"
)

type ZeroconfService struct {
	service *zeroconf.Server
	config  domain.Config
}

var service *ZeroconfService = nil

func NewZeroconfService(config domain.Config) (*ZeroconfService, error) {
	if service != nil {
		return service, nil
	}

	server, err := zeroconf.Register("Synk-"+config.Service.Name, config.Service.Type, config.Service.Domain, config.Service.Port, nil, nil)
	if err != nil {
		return nil, err
	}

	service = &ZeroconfService{service: server, config: config}

	return service, nil
}

func (z *ZeroconfService) Unregister() error {
	if z.service != nil {
		z.service.Shutdown()
		return nil
	}
	return nil
}
