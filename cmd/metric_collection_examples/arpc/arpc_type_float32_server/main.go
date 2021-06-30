package main

import (
	"fmt"
	"os"

	"github.com/almeida-raphael/arpc_examples/models/arpc/typefloat32"

	"github.com/almeida-raphael/arpc/interfaces"
	"github.com/almeida-raphael/arpc_examples/utils"
)

// ServerDefinition struct to implement greetings aRPC service procedures
type ServerDefinition struct{}

func typeFloat32(request interfaces.Serializable) (interfaces.Serializable, error) {
	return request, nil
}

var metricsTypeFloat32 = utils.CollectServerMetrics(
	utils.Atoi(os.Getenv("SAMPLE_THREADS")), utils.Atoi(os.Getenv("TRIALS")), typeFloat32,
	fmt.Sprintf(
		"results/aRPC/type_float32_%d_threads",
		utils.Atoi(os.Getenv("SAMPLE_THREADS")),
	)+"/server/%d.json",
)

// TypeFloat32 aRPC server function implementation
func (gs *ServerDefinition) TypeFloat32(amount *typefloat32.Request) (*typefloat32.EmptyResult, error) {
	var reqData interfaces.Serializable = amount
	respData := &typefloat32.EmptyResult{}

	response, err := metricsTypeFloat32(reqData)
	if err == nil {
		respData = response.(*typefloat32.EmptyResult)
	}

	return respData, err
}

func main() {
	aRPCController := utils.SetupServer()
	typefloat32.RegisterTypefloat32Server(aRPCController, &ServerDefinition{})
	utils.StartServer(aRPCController)
}
