package main

import (
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/almeida-raphael/arpc/channel"
	"github.com/almeida-raphael/arpc/controller"
	arpcerrors "github.com/almeida-raphael/arpc/errors"
	"github.com/almeida-raphael/arpc_examples/examples/greetings"
	"github.com/almeida-raphael/arpc_examples/utils"
	"os"
)


func main(){
	rootCAPath := os.Getenv("CA_FILE")
	serverAddress := os.Getenv("SERVER_ADDRESS")

	rootCA, err := utils.LoadCA(rootCAPath)
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
		nil,
	))
	greetingsService := greetings.NewGreetings(&aRPCController)

	err = aRPCController.StartClient()
	if err != nil {
		panic(err)
	}

	hiResponse, err := greetingsService.SayHi(&greetings.SayHiRequest{
		Person: &greetings.Person{
			Title: 10,
			Name:  "Raphael",
		},
	})
	if err != nil {
		if errors.Is(err, &arpcerrors.Remote{}) {
			fmt.Printf("Remote Error: %v", err)
		}else{
			panic(err)
		}
	}else{
		fmt.Printf("Response: %s", hiResponse.Response)
	}
}

