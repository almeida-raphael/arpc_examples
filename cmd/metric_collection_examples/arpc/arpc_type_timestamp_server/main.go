package main

import (
	"fmt"
	"os"

	"github.com/almeida-raphael/arpc_examples/models/arpc/typetimestamp"

	"github.com/almeida-raphael/arpc/interfaces"
	"github.com/almeida-raphael/arpc_examples/utils"
)

// ServerDefinition struct to implement greetings aRPC service procedures
type ServerDefinition struct{}

func typeTimestamp(request interfaces.Serializable) (interfaces.Serializable, error) {
	return request, nil
}

var metricsTypeTimestamp = utils.CollectServerMetrics(
	utils.Atoi(os.Getenv("SAMPLE_THREADS")), utils.Atoi(os.Getenv("TRIALS")), typeTimestamp,
	fmt.Sprintf(
		"results/aRPC/type_timestamp_%d_threads",
		utils.Atoi(os.Getenv("SAMPLE_THREADS")),
	)+"/server/%d.json",
)

// TypeTimestamp aRPC server function implementation
func (gs *ServerDefinition) TypeTimestamp(amount *typetimestamp.Request) (*typetimestamp.EmptyResult, error) {
	var reqData interfaces.Serializable = amount
	respData := &typetimestamp.EmptyResult{}

	response, err := metricsTypeTimestamp(reqData)
	if err == nil {
		respData = response.(*typetimestamp.EmptyResult)
	}

	return respData, err
}

func main() {
	aRPCController := utils.SetupServer()
	typetimestamp.RegisterTypetimestampServer(aRPCController, &ServerDefinition{})
	utils.StartServer(aRPCController)
}
