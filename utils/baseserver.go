package utils

import (
	"context"
	"crypto/tls"
	"os"
	"strconv"
	"strings"

	"github.com/almeida-raphael/arpc/channel"
	"github.com/almeida-raphael/arpc/controller"
	"github.com/lucas-clemente/quic-go"
)

// SetupServer setup basic server config
func SetupServer() controller.RPC {
	certFilePath := os.Getenv("CERT_FILE")
	keyFilePath := os.Getenv("KEY_FILE")
	serverAddress := os.Getenv("SERVER_ADDRESS")

	splitAddress := strings.Split(serverAddress, ":")
	address := splitAddress[0]
	port, err := strconv.Atoi(splitAddress[1])
	if err != nil {
		panic(err)
	}

	certificates, err := LoadCertificates(certFilePath, keyFilePath)
	if err != nil {
		panic(err)
	}

	tlsConfig := tls.Config{
		Certificates: []tls.Certificate{*certificates},
		NextProtos:   []string{"quic-arcp"},
	}

	aRPCController := controller.NewRPCController(channel.NewQUICChannel(
		address,
		port,
		&tlsConfig,
		&quic.Config{
			MaxIncomingStreams: 10000000,
		},
	))

	return aRPCController
}

// StartServer start the server and panic on startup errors
func StartServer(contr controller.RPC) {
	err := contr.StartServer(context.Background())
	if err != nil {
		panic(err)
	}
}
