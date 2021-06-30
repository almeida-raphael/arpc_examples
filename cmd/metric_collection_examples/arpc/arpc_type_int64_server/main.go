package main

import (
	"fmt"
	"os"

	"github.com/almeida-raphael/arpc_examples/models/arpc/typeint64"

	"github.com/almeida-raphael/arpc/interfaces"
	"github.com/almeida-raphael/arpc_examples/utils"
)

// ServerDefinition struct to implement greetings aRPC service procedures
type ServerDefinition struct{}

func typeInt64(request interfaces.Serializable) (interfaces.Serializable, error) {
	return request, nil
}

var metricsTypeInt64 = utils.CollectServerMetrics(
	utils.Atoi(os.Getenv("SAMPLE_THREADS")), utils.Atoi(os.Getenv("TRIALS")), typeInt64,
	fmt.Sprintf(
		"results/aRPC/type_int64_%d_threads",
		utils.Atoi(os.Getenv("SAMPLE_THREADS")),
	)+"/server/%d.json",
)

// TypeInt64 aRPC server function implementation
func (gs *ServerDefinition) TypeInt64(amount *typeint64.Request) (*typeint64.EmptyResult, error) {
	var reqData interfaces.Serializable = amount
	respData := &typeint64.EmptyResult{}

	response, err := metricsTypeInt64(reqData)
	if err == nil {
		respData = response.(*typeint64.EmptyResult)
	}

	return respData, err
}

func main() {
	aRPCController := utils.SetupServer()
	typeint64.RegisterTypeint64Server(aRPCController, &ServerDefinition{})
	utils.StartServer(aRPCController)
}
