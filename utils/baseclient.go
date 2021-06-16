package utils

import (
	"crypto/tls"
	"github.com/almeida-raphael/arpc/channel"
	"github.com/almeida-raphael/arpc/controller"
	"github.com/lucas-clemente/quic-go"
	"os"
)

// SetupClient Setup basic client configs
func SetupClient() controller.RPC {
	rootCAPath := os.Getenv("CA_FILE")
	serverAddress := os.Getenv("SERVER_ADDRESS")

	rootCA, err := LoadCA(rootCAPath)
	if err != nil {
		panic(err)
	}

	tlsConfig := tls.Config{
		RootCAs:                     rootCA,
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
