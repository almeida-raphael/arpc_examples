package main

import (
	"fmt"
	"os"

	"github.com/almeida-raphael/arpc_examples/models/arpc/typebinary"

	"github.com/almeida-raphael/arpc/interfaces"
	"github.com/almeida-raphael/arpc_examples/utils"
)

// ServerDefinition struct to implement greetings aRPC service procedures
type ServerDefinition struct{}

func typeBinary(request interfaces.Serializable) (interfaces.Serializable, error) {
	return request, nil
}

var metricsTypeBinary = utils.CollectServerMetrics(
	utils.Atoi(os.Getenv("SAMPLE_THREADS")), utils.Atoi(os.Getenv("TRIALS")), typeBinary,
	fmt.Sprintf(
		"results/aRPC/type_binary_%d_threads",
		utils.Atoi(os.Getenv("SAMPLE_THREADS")),
	)+"/server/%d.json",
)

// TypeBinary aRPC server function implementation
func (gs *ServerDefinition) TypeBinary(amount *typebinary.Request) (*typebinary.EmptyResult, error) {
	var reqData interfaces.Serializable = amount
	respData := &typebinary.EmptyResult{}

	response, err := metricsTypeBinary(reqData)
	if err == nil {
		respData = response.(*typebinary.EmptyResult)
	}

	return respData, err
}

func main() {
	aRPCController := utils.SetupServer()
	typebinary.RegisterTypebinaryServer(aRPCController, &ServerDefinition{})
	utils.StartServer(aRPCController)
}
