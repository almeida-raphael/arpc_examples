package main

import (
	"fmt"
	"github.com/almeida-raphael/arpc_examples/examples/greetings"
	"github.com/almeida-raphael/arpc_examples/utils"
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
	aRPCController := utils.SetupServer()
	greetings.RegisterGreetingsServer(aRPCController, &GreetingsServerDefinition{})
	utils.StartServer(aRPCController)
}

