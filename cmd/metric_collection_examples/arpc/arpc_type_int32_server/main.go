package main

import (
	"fmt"
	"os"

	"github.com/almeida-raphael/arpc_examples/models/arpc/typeint32"

	"github.com/almeida-raphael/arpc/interfaces"
	"github.com/almeida-raphael/arpc_examples/utils"
)

// ServerDefinition struct to implement greetings aRPC service procedures
type ServerDefinition struct{}

func typeInt32(request interfaces.Serializable) (interfaces.Serializable, error) {
	return request, nil
}

var metricsTypeInt32 = utils.CollectServerMetrics(
	utils.Atoi(os.Getenv("SAMPLE_THREADS")), utils.Atoi(os.Getenv("TRIALS")), typeInt32,
	fmt.Sprintf(
		"results/aRPC/type_int32_%d_threads",
		utils.Atoi(os.Getenv("SAMPLE_THREADS")),
	)+"/server/%d.json",
)

// TypeInt32 aRPC server function implementation
func (gs *ServerDefinition) TypeInt32(amount *typeint32.Request) (*typeint32.EmptyResult, error) {
	var reqData interfaces.Serializable = amount
	respData := &typeint32.EmptyResult{}

	response, err := metricsTypeInt32(reqData)
	if err == nil {
		respData = response.(*typeint32.EmptyResult)
	}

	return respData, err
}

func main() {
	aRPCController := utils.SetupServer()
	typeint32.RegisterTypeint32Server(aRPCController, &ServerDefinition{})
	utils.StartServer(aRPCController)
}
