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
	z.findAndConnectToNode("471d424ac743c8fd")
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

func (z *ZeroconfService) findAndConnectToNode(targetID string) {
	fmt.Printf("Attempting to add device with ID: %s\n", targetID)
	timeout := 10 * time.Second
	resolver, err := zeroconf.NewResolver(nil)
	if err != nil {
		log.Error(fmt.Sprintf("Failed to initialize resolver: %s", err))
	}

	entries := make(chan *zeroconf.ServiceEntry)
	go func() {
		for entry := range entries {
			fmt.Printf("Dispositivo encontrado: %v\n", entry)
			fmt.Printf("  Nome: %s\n", entry.Instance)
			fmt.Printf("  Endereço: %v\n", entry.AddrIPv4)
			fmt.Printf("  Porta: %d\n", entry.Port)
			fmt.Printf("  Metadados: %v\n", entry.Text)
			if entry.Instance == fmt.Sprintf("Synk-%s", targetID) {
				fmt.Printf("Conectando ao dispositivo com ID: %s\n", targetID)
				connectToDevice(entry)
				return
			}
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	err = resolver.Browse(ctx, z.config.Service.Type, z.config.Service.Domain, entries)
	if err != nil {
		log.Error(fmt.Sprintf("Failed to browse for services: %s", err))
	}

	<-ctx.Done()
	fmt.Println("Busca de serviços concluída.")
}

func connectToDevice(entry *zeroconf.ServiceEntry) {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", entry.AddrIPv4[0], entry.Port), 5*time.Second)
	if err != nil {
		fmt.Println("Failed to connect to device:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Successfully connected to the device! Ready to exchange keys and synchronize.")
	// Aqui você implementaria a lógica de troca de chaves e sincronização
	// usando a chave pública do outro dispositivo.
}
