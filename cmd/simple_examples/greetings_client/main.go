package main

import (
	"fmt"
	"github.com/almeida-raphael/arpc_examples/examples/greetings"
	"github.com/almeida-raphael/arpc_examples/utils"
)

func main(){
	aRPCController := utils.SetupClient()

	greetingsService := greetings.NewGreetings(&aRPCController)
	err := aRPCController.StartClient()
	if err != nil {
		panic(err)
	}

	requestData := greetings.SayHiRequest{
		Person: &greetings.Person{
			Title: 10,
			Name:  "Raphael",
		},
	}

	hiResponse, err := greetingsService.SayHi(&requestData)

	if !utils.HandleRemoteError(err){
		fmt.Printf("Response: %s", hiResponse.Response)
	}
}

