package utils

import (
	"context"
	"crypto/tls"
	"github.com/almeida-raphael/arpc/channel"
	"github.com/almeida-raphael/arpc/controller"
	"github.com/lucas-clemente/quic-go"
	"os"
)

// SetupServer setup basic server config
func SetupServer() controller.RPC {
	certFilePath := os.Getenv("CERT_FILE")
	keyFilePath := os.Getenv("KEY_FILE")
	serverAddress := os.Getenv("SERVER_ADDRESS")

	certificates, err := LoadCertificates(certFilePath, keyFilePath)
	if err != nil {
		panic(err)
	}

	tlsConfig := tls.Config{
		Certificates:                []tls.Certificate{*certificates},
		NextProtos:                  []string{"quic-arcp"},
	}

	aRPCController := controller.NewRPCController(channel.NewQUICChannel(
		serverAddress,
		7653,
		&tlsConfig,
		&quic.Config{
			MaxIncomingStreams: 10000000,
		},
	))

	return aRPCController
}

// StartServer start the server and panic on startup errors
func StartServer(contr controller.RPC){
	err := contr.StartServer(context.Background())
	if err != nil {
		panic(err)
	}
}