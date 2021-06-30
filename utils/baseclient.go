package utils

import (
	"crypto/tls"
	"os"
	"strconv"
	"strings"

	"github.com/almeida-raphael/arpc/channel"
	"github.com/almeida-raphael/arpc/controller"
	"github.com/lucas-clemente/quic-go"
)

// SetupClient Setup basic client configs
func SetupClient() controller.RPC {
	rootCAPath := os.Getenv("CA_FILE")
	serverAddress := os.Getenv("SERVER_ADDRESS")

	splitAddress := strings.Split(serverAddress, ":")
	address := splitAddress[0]
	port, err := strconv.Atoi(splitAddress[1])
	if err != nil {
		panic(err)
	}

	rootCA, err := LoadCA(rootCAPath)
	if err != nil {
		panic(err)
	}

	tlsConfig := tls.Config{
		RootCAs:    rootCA,
		NextProtos: []string{"quic-arcp"},
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
