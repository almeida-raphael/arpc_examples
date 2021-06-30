package main

import (
	"fmt"
	"os"

	doubletype2 "github.com/almeida-raphael/arpc_examples/models/hprpc/doubletype"

	"github.com/almeida-raphael/arpc/interfaces"
	"github.com/almeida-raphael/arpc_examples/utils"
)

// ServerDefinition struct to implement greetings aRPC service procedures
type ServerDefinition struct{}

// getNumbers aRPC server function implementation
func average(request interfaces.Serializable) (interfaces.Serializable, error) {
	reqData := request.(*doubletype2.NumberList)
	return &doubletype2.Result{Value: utils.Average(reqData.Entries)}, nil
}

var metricsAverage = utils.CollectServerMetrics(
	utils.Atoi(os.Getenv("SAMPLE_THREADS")), utils.Atoi(os.Getenv("TRIALS")), average,
	fmt.Sprintf(
		"results/aRPC/doubletype_%s_%d_threads", os.Getenv("VALUE"),
		utils.Atoi(os.Getenv("SAMPLE_THREADS")),
	)+"/server/%d.json",
)

// Average aRPC server function implementation
func (gs *ServerDefinition) Average(request *doubletype2.NumberList) (*doubletype2.Result, error) {
	var reqData interfaces.Serializable = request
	respData := &doubletype2.Result{}

	response, err := metricsAverage(reqData)
	if err == nil {
		respData = response.(*doubletype2.Result)
	}

	return respData, err
}

func main() {
	aRPCController := utils.SetupServer()
	doubletype2.RegisterDoubletypeServer(aRPCController, &ServerDefinition{})
	utils.StartServer(aRPCController)
}
