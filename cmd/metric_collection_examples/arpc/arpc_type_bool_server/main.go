package main

import (
	"fmt"
	"os"

	"github.com/almeida-raphael/arpc_examples/models/arpc/typebool"

	"github.com/almeida-raphael/arpc/interfaces"
	"github.com/almeida-raphael/arpc_examples/utils"
)

// ServerDefinition struct to implement greetings aRPC service procedures
type ServerDefinition struct{}

func typeBool(request interfaces.Serializable) (interfaces.Serializable, error) {
	return request, nil
}

var metricsTypeBool = utils.CollectServerMetrics(
	utils.Atoi(os.Getenv("SAMPLE_THREADS")), utils.Atoi(os.Getenv("TRIALS")), typeBool,
	fmt.Sprintf(
		"results/aRPC/type_bool_%d_threads",
		utils.Atoi(os.Getenv("SAMPLE_THREADS")),
	)+"/server/%d.json",
)

// TypeBool aRPC server function implementation
func (gs *ServerDefinition) TypeBool(amount *typebool.Request) (*typebool.EmptyResult, error) {
	var reqData interfaces.Serializable = amount
	respData := &typebool.EmptyResult{}

	response, err := metricsTypeBool(reqData)
	if err == nil {
		respData = response.(*typebool.EmptyResult)
	}

	return respData, err
}

func main() {
	aRPCController := utils.SetupServer()
	typebool.RegisterTypeboolServer(aRPCController, &ServerDefinition{})
	utils.StartServer(aRPCController)
}
