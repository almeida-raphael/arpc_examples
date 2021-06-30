package main

import (
	"fmt"
	"os"

	"github.com/almeida-raphael/arpc_examples/models/arpc/typetext"

	"github.com/almeida-raphael/arpc/interfaces"
	"github.com/almeida-raphael/arpc_examples/utils"
)

// ServerDefinition struct to implement greetings aRPC service procedures
type ServerDefinition struct{}

func typeText(request interfaces.Serializable) (interfaces.Serializable, error) {
	return request, nil
}

var metricsTypeText = utils.CollectServerMetrics(
	utils.Atoi(os.Getenv("SAMPLE_THREADS")), utils.Atoi(os.Getenv("TRIALS")), typeText,
	fmt.Sprintf(
		"results/aRPC/type_text_%d_threads",
		utils.Atoi(os.Getenv("SAMPLE_THREADS")),
	)+"/server/%d.json",
)

// TypeText aRPC server function implementation
func (gs *ServerDefinition) TypeText(amount *typetext.Request) (*typetext.EmptyResult, error) {
	var reqData interfaces.Serializable = amount
	respData := &typetext.EmptyResult{}

	response, err := metricsTypeText(reqData)
	if err == nil {
		respData = response.(*typetext.EmptyResult)
	}

	return respData, err
}

func main() {
	aRPCController := utils.SetupServer()
	typetext.RegisterTypetextServer(aRPCController, &ServerDefinition{})
	utils.StartServer(aRPCController)
}
