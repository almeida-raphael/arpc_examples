package main

import (
	"fmt"
	"os"

	"github.com/almeida-raphael/arpc_examples/models/arpc/typeall"

	"github.com/almeida-raphael/arpc/interfaces"
	"github.com/almeida-raphael/arpc_examples/utils"
)

// ServerDefinition struct to implement greetings aRPC service procedures
type ServerDefinition struct{}

func typeAll(request interfaces.Serializable) (interfaces.Serializable, error) {
	return &typeall.EmptyResult{}, nil
}

var metricsTypeAll = utils.CollectServerMetrics(
	utils.Atoi(os.Getenv("SAMPLE_THREADS")), utils.Atoi(os.Getenv("TRIALS")), typeAll,
	fmt.Sprintf(
		"results/aRPC/type_all_%s_%d_threads",
		os.Getenv("VALUE"), utils.Atoi(os.Getenv("SAMPLE_THREADS")),
	)+"/server/%d.json",
)

// TypeAll aRPC server function implementation
func (gs *ServerDefinition) TypeAll(amount *typeall.Request) (*typeall.EmptyResult, error) {
	var reqData interfaces.Serializable = amount
	respData := &typeall.EmptyResult{}

	response, err := metricsTypeAll(reqData)
	if err == nil {
		respData = response.(*typeall.EmptyResult)
	}

	return respData, err
}

func main() {
	aRPCController := utils.SetupServer()
	typeall.RegisterTypeallServer(aRPCController, &ServerDefinition{})
	utils.StartServer(aRPCController)
}
