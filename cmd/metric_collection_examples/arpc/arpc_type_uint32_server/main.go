package main

import (
	"fmt"
	"os"

	"github.com/almeida-raphael/arpc_examples/models/arpc/typeuint32"

	"github.com/almeida-raphael/arpc/interfaces"
	"github.com/almeida-raphael/arpc_examples/utils"
)

// ServerDefinition struct to implement greetings aRPC service procedures
type ServerDefinition struct{}

func typeUInt32(request interfaces.Serializable) (interfaces.Serializable, error) {
	return request, nil
}

var metricsTypeUInt32 = utils.CollectServerMetrics(
	utils.Atoi(os.Getenv("SAMPLE_THREADS")), utils.Atoi(os.Getenv("TRIALS")), typeUInt32,
	fmt.Sprintf(
		"results/aRPC/type_uint32_%d_threads",
		utils.Atoi(os.Getenv("SAMPLE_THREADS")),
	)+"/server/%d.json",
)

// TypeUInt32 aRPC server function implementation
func (gs *ServerDefinition) TypeUInt32(amount *typeuint32.Request) (*typeuint32.EmptyResult, error) {
	var reqData interfaces.Serializable = amount
	respData := &typeuint32.EmptyResult{}

	response, err := metricsTypeUInt32(reqData)
	if err == nil {
		respData = response.(*typeuint32.EmptyResult)
	}

	return respData, err
}

func main() {
	aRPCController := utils.SetupServer()
	typeuint32.RegisterTypeuint32Server(aRPCController, &ServerDefinition{})
	utils.StartServer(aRPCController)
}
