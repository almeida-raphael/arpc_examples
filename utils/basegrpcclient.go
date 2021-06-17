package utils

import (
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// SetupGRPCClient Setup basic grpc client configs
func SetupGRPCClient() (*grpc.ClientConn, error) {
	rootCAPath := os.Getenv("CA_FILE")
	serverAddress := os.Getenv("SERVER_ADDRESS")

	rootCA, err := LoadCA(rootCAPath)
	if err != nil {
		log.Fatal(err)
	}

	return grpc.Dial(
		serverAddress,
		grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(rootCA, "")),
		grpc.WithBlock(),
	)
}
