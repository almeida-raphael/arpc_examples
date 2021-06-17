package utils

import (
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// SetupGRPCServer setup basic gRPC server config
func SetupGRPCServer() (net.Listener, *grpc.Server) {
	certFilePath := os.Getenv("CERT_FILE")
	keyFilePath := os.Getenv("KEY_FILE")
	serverAddress := os.Getenv("SERVER_ADDRESS")

	listen, err := net.Listen("tcp", serverAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	certificates, err := LoadCertificates(certFilePath, keyFilePath)
	if err != nil {
		log.Fatal(err)
	}

	credential := credentials.NewServerTLSFromCert(certificates)

	return listen, grpc.NewServer(grpc.Creds(credential),
		grpc.MaxRecvMsgSize(10*1024*1024),
		grpc.MaxSendMsgSize(10*1024*1024))
}

// StartGRPCServer start the server and panic on startup errors
func StartGRPCServer(listener net.Listener, server *grpc.Server) {
	log.Println("Serving...")
	if err := server.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
