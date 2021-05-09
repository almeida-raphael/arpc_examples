package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/almeida-raphael/arpc/channel"
	"github.com/almeida-raphael/arpc/controller"
	"github.com/almeida-raphael/arpc_examples/examples/greetings"
	"github.com/almeida-raphael/arpc_examples/utils"
	"os"
)

// GreetingsServerDefinition struct to implement greetings aRPC service procedures
type GreetingsServerDefinition struct {}

// SayHi aRPC Greetings.SayHi function implementation
func (gs *GreetingsServerDefinition)SayHi(request *greetings.SayHiRequest)(*greetings.SayHiResponse, error){
	return &greetings.SayHiResponse{
		Response: fmt.Sprintf("hi %s with title %d", request.Person.Name, request.Person.Title),
	}, nil
}

func main(){
	certFilePath := os.Getenv("CERT_FILE")
	keyFilePath := os.Getenv("KEY_FILE")
	serverAddress := os.Getenv("SERVER_ADDRESS")

	certificates, err := utils.LoadCertificates(certFilePath, keyFilePath)
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
		nil,
	))

	greetings.RegisterGreetingsServer(aRPCController, &GreetingsServerDefinition{})

	err = aRPCController.StartServer(context.Background())
	if err != nil {
		panic(err)
	}
}

