package service

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"synk/config"
	"synk/internal/domain"
	"time"

	"synk/pkg/logger"
	"synk/pkg/utils"
	"syscall"

	"github.com/grandcat/zeroconf"
)

var log = logger.GetLogger("zeroconf")

var service *ZeroconfService

type ZeroconfService struct {
	server *zeroconf.Server
	config domain.Config
	device *domain.Device
	Port   int
}

func NewZeroconfService(config domain.Config, device *domain.Device) (*ZeroconfService, error) {
	if service == nil {
		service = &ZeroconfService{
			config: config,
			device: device,
			Port:   config.Service.Port,
		}
	}

	return service, nil
}

func (z *ZeroconfService) Start() error {
	metadata := []string{"id=" + z.device.ID}

	server, err := zeroconf.Register(
		z.config.Service.Name,
		z.config.Service.Type,
		z.config.Service.Domain,
		z.config.Service.Port,
		metadata,
		nil,
	)
	if err != nil {
		return err
	}

	z.server = server
	log.Info(fmt.Sprintf("Serviço registrado: %s | Porta: %d | ID: %s",
		z.config.Service.Name, z.config.Service.Port, z.device.ID))

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigChan
		z.Stop()
		os.Exit(0)
	}()

	return nil
}

func (z *ZeroconfService) Stop() error {
	if z.server != nil {
		z.server.Shutdown()
		log.Info("Serviço zeroconf finalizado: " + z.config.Service.Name)
	}

	return nil
}

func (z *ZeroconfService) GetPort() int {
	return z.Port
}

func (z *ZeroconfService) AddDeviceToConnect(ID string) error {
	resolver, err := zeroconf.NewResolver(nil)
	if err != nil {
		log.Error("Erro no resolvedor: " + err.Error())
		return err
	}

	entries := make(chan *zeroconf.ServiceEntry)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	go func() {
		z.findDeviceAndConnect(entries, ID)
	}()

	err = resolver.Browse(ctx, z.config.Service.Type, z.config.Service.Domain, entries)
	if err != nil {
		log.Error("Erro na busca: " + err.Error())
	}

	<-ctx.Done()
	return nil
}

func (z *ZeroconfService) findDeviceAndConnect(entries chan *zeroconf.ServiceEntry, ID string) error {
	for entry := range entries {
		deviceID := getIDFromMetadata(entry.Text)
		if deviceID != ID {
			continue
		}
		log.Info(fmt.Sprintf("ID encontrado: %s", deviceID))
		go connectToDevice(entry)

		device := config.GetDevice()
		if device == nil {
			log.Error("Dispositivo não encontrado, execute 'synk init' primeiro")
			continue
		}
		device.Connections = utils.RemoveDuplicates(append(device.Connections, deviceID))
		config.SaveDevice(device)
	}

	return nil
}

func getIDFromMetadata(txt []string) string {
	for _, item := range txt {
		if len(item) > 3 && item[:3] == "id=" {
			return item[3:]
		}
	}
	return ""
}

func connectToDevice(entry *zeroconf.ServiceEntry) {
	if len(entry.AddrIPv4) == 0 {
		log.Error("Sem endereço IPv4 para: " + entry.Instance)
		return
	}

	for ip := range entry.AddrIPv4 {
		if ip <= 0 {
			continue
		}
		address := fmt.Sprintf("Conectando a %s:%d", entry.AddrIPv4[ip], entry.Port)
		log.Info(address)
		conn, err := net.DialTimeout("tcp", address, 5*time.Second)
		if err != nil {
			continue
		}
		defer conn.Close()
		break
	}

	log.Info("Conectado com sucesso a: " + entry.Instance)
}
