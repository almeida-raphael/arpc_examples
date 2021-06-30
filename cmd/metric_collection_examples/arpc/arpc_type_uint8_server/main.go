package main

import (
	"fmt"
	"os"

	"github.com/almeida-raphael/arpc_examples/models/arpc/typeuint8"

	"github.com/almeida-raphael/arpc/interfaces"
	"github.com/almeida-raphael/arpc_examples/utils"
)

// ServerDefinition struct to implement greetings aRPC service procedures
type ServerDefinition struct{}

func typeUInt8(request interfaces.Serializable) (interfaces.Serializable, error) {
	return request, nil
}

var metricsTypeUInt8 = utils.CollectServerMetrics(
	utils.Atoi(os.Getenv("SAMPLE_THREADS")), utils.Atoi(os.Getenv("TRIALS")), typeUInt8,
	fmt.Sprintf(
		"results/aRPC/type_uint8_%d_threads",
		utils.Atoi(os.Getenv("SAMPLE_THREADS")),
	)+"/server/%d.json",
)

// TypeUInt8 aRPC server function implementation
func (gs *ServerDefinition) TypeUInt8(amount *typeuint8.Request) (*typeuint8.EmptyResult, error) {
	var reqData interfaces.Serializable = amount
	respData := &typeuint8.EmptyResult{}

	response, err := metricsTypeUInt8(reqData)
	if err == nil {
		respData = response.(*typeuint8.EmptyResult)
	}

	return respData, err
}

func main() {
	aRPCController := utils.SetupServer()
	typeuint8.RegisterTypeuint8Server(aRPCController, &ServerDefinition{})
	utils.StartServer(aRPCController)
}
