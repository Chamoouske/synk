package service

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"synk/internal/domain"
	"time"

	"synk/pkg/logger"
	"syscall"

	"github.com/grandcat/zeroconf"
)

var log = logger.GetLogger("zeroconf")

type ZeroconfService struct {
	server *zeroconf.Server
	config domain.Config
	device *domain.Device
}

func NewZeroconfService(config domain.Config, device *domain.Device) (*ZeroconfService, error) {
	return &ZeroconfService{
		config: config,
		device: device,
	}, nil
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

	go z.startTCPServer()

	go z.continuousDiscovery()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigChan
		z.Stop()
		os.Exit(0)
	}()

	return nil
}

func (z *ZeroconfService) startTCPServer() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", z.config.Service.Port))
	if err != nil {
		log.Error("Falha ao iniciar TCP server: " + err.Error())
		return
	}
	defer listener.Close()

	log.Info("Escutando em TCP porta: " + fmt.Sprintf("%d", z.config.Service.Port))

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Error("Erro de conexão: " + err.Error())
			continue
		}
		go z.handleConnection(conn)
	}
}

func (z *ZeroconfService) handleConnection(conn net.Conn) {
	defer conn.Close()
	remoteAddr := conn.RemoteAddr().String()
	log.Info("Conexão estabelecida com: " + remoteAddr)
}

func (z *ZeroconfService) continuousDiscovery() {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		z.discoverDevices()
		<-ticker.C
	}
}

func (z *ZeroconfService) discoverDevices() {
	resolver, err := zeroconf.NewResolver(nil)
	if err != nil {
		log.Error("Erro no resolvedor: " + err.Error())
		return
	}

	entries := make(chan *zeroconf.ServiceEntry)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	go func() {
		z.findDevices(entries)
	}()

	err = resolver.Browse(ctx, z.config.Service.Type, z.config.Service.Domain, entries)
	if err != nil {
		log.Error("Erro na busca: " + err.Error())
	}

	<-ctx.Done()
}

func (z *ZeroconfService) findDevices(entries chan *zeroconf.ServiceEntry) error {
	for entry := range entries {
		deviceID := getIDFromMetadata(entry.Text)
		if deviceID == "" {
			continue
		}
		if deviceID == z.device.ID {
			continue
		}
		log.Info(fmt.Sprintf("ID encontrado: %s", deviceID))
		go connectToDevice(entry)
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

	address := fmt.Sprintf("%s:%d", entry.AddrIPv4[0], entry.Port)
	conn, err := net.DialTimeout("tcp", address, 5*time.Second)
	if err != nil {
		log.Error("Conexão falhou: " + err.Error())
		return
	}
	defer conn.Close()

	log.Info("Conectado com sucesso a: " + entry.Instance)
}

func (z *ZeroconfService) Stop() error {
	if z.server != nil {
		z.server.Shutdown()
		log.Info("Serviço zeroconf finalizado: " + z.config.Service.Name)
	}

	return nil
}
