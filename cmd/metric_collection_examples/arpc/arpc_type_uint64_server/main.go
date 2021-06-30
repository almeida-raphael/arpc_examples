package main

import (
	"fmt"
	"os"

	"github.com/almeida-raphael/arpc_examples/models/arpc/typeuint64"

	"github.com/almeida-raphael/arpc/interfaces"
	"github.com/almeida-raphael/arpc_examples/utils"
)

// ServerDefinition struct to implement greetings aRPC service procedures
type ServerDefinition struct{}

func typeUInt64(request interfaces.Serializable) (interfaces.Serializable, error) {
	return request, nil
}

var metricsTypeUInt64 = utils.CollectServerMetrics(
	utils.Atoi(os.Getenv("SAMPLE_THREADS")), utils.Atoi(os.Getenv("TRIALS")), typeUInt64,
	fmt.Sprintf(
		"results/aRPC/type_uint64_%d_threads",
		utils.Atoi(os.Getenv("SAMPLE_THREADS")),
	)+"/server/%d.json",
)

// TypeUInt64 aRPC server function implementation
func (gs *ServerDefinition) TypeUInt64(amount *typeuint64.Request) (*typeuint64.EmptyResult, error) {
	var reqData interfaces.Serializable = amount
	respData := &typeuint64.EmptyResult{}

	response, err := metricsTypeUInt64(reqData)
	if err == nil {
		respData = response.(*typeuint64.EmptyResult)
	}

	return respData, err
}

func main() {
	aRPCController := utils.SetupServer()
	typeuint64.RegisterTypeuint64Server(aRPCController, &ServerDefinition{})
	utils.StartServer(aRPCController)
}
