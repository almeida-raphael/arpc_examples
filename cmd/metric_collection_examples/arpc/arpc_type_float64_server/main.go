package main

import (
	"fmt"
	"os"

	"github.com/almeida-raphael/arpc_examples/models/arpc/typefloat64"

	"github.com/almeida-raphael/arpc/interfaces"
	"github.com/almeida-raphael/arpc_examples/utils"
)

// ServerDefinition struct to implement greetings aRPC service procedures
type ServerDefinition struct{}

func typeFloat64(request interfaces.Serializable) (interfaces.Serializable, error) {
	return request, nil
}

var metricsTypeFloat64 = utils.CollectServerMetrics(
	utils.Atoi(os.Getenv("SAMPLE_THREADS")), utils.Atoi(os.Getenv("TRIALS")), typeFloat64,
	fmt.Sprintf(
		"results/aRPC/type_float64_%d_threads",
		utils.Atoi(os.Getenv("SAMPLE_THREADS")),
	)+"/server/%d.json",
)

// TypeFloat64 aRPC server function implementation
func (gs *ServerDefinition) TypeFloat64(amount *typefloat64.Request) (*typefloat64.EmptyResult, error) {
	var reqData interfaces.Serializable = amount
	respData := &typefloat64.EmptyResult{}

	response, err := metricsTypeFloat64(reqData)
	if err == nil {
		respData = response.(*typefloat64.EmptyResult)
	}

	return respData, err
}

func main() {
	aRPCController := utils.SetupServer()
	typefloat64.RegisterTypefloat64Server(aRPCController, &ServerDefinition{})
	utils.StartServer(aRPCController)
}
