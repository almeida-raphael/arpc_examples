package main

import (
	"fmt"
	"os"

	"github.com/almeida-raphael/arpc_examples/models/arpc/typeuint16"

	"github.com/almeida-raphael/arpc/interfaces"
	"github.com/almeida-raphael/arpc_examples/utils"
)

// ServerDefinition struct to implement greetings aRPC service procedures
type ServerDefinition struct{}

func typeUInt16(request interfaces.Serializable) (interfaces.Serializable, error) {
	return request, nil
}

var metricsTypeUInt16 = utils.CollectServerMetrics(
	utils.Atoi(os.Getenv("SAMPLE_THREADS")), utils.Atoi(os.Getenv("TRIALS")), typeUInt16,
	fmt.Sprintf(
		"results/aRPC/type_uint16_%d_threads",
		utils.Atoi(os.Getenv("SAMPLE_THREADS")),
	)+"/server/%d.json",
)

// TypeUInt16 aRPC server function implementation
func (gs *ServerDefinition) TypeUInt16(amount *typeuint16.Request) (*typeuint16.EmptyResult, error) {
	var reqData interfaces.Serializable = amount
	respData := &typeuint16.EmptyResult{}

	response, err := metricsTypeUInt16(reqData)
	if err == nil {
		respData = response.(*typeuint16.EmptyResult)
	}

	return respData, err
}

func main() {
	aRPCController := utils.SetupServer()
	typeuint16.RegisterTypeuint16Server(aRPCController, &ServerDefinition{})
	utils.StartServer(aRPCController)
}
