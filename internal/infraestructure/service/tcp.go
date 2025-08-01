package service

import (
	"fmt"
	"net"
)

func StartTCPServer(port int) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Error("Falha ao iniciar TCP server: " + err.Error())
		return
	}
	defer listener.Close()

	log.Info("Escutando em TCP porta: " + fmt.Sprintf("%d", port))

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Error("Erro de conexão: " + err.Error())
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	remoteAddr := conn.RemoteAddr().String()
	log.Info("Conexão estabelecida com: " + remoteAddr)
}
