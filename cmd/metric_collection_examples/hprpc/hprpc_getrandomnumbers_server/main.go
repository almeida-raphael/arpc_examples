package main

import (
	"fmt"
	"os"

	getrandomnumbers2 "github.com/almeida-raphael/arpc_examples/models/hprpc/getrandomnumbers"

	"github.com/almeida-raphael/arpc/interfaces"
	"github.com/almeida-raphael/arpc_examples/utils"
)

// ServerDefinition struct to implement greetings aRPC service procedures
type ServerDefinition struct{}

// getNumbers aRPC server function implementation
func getNumbers(request interfaces.Serializable) (interfaces.Serializable, error) {
	reqData := request.(*getrandomnumbers2.Amount)
	return &getrandomnumbers2.NumberList{Entries: utils.GenerateNumbers(int(reqData.Value))}, nil
}

var metricsGetNumbers = utils.CollectServerMetrics(
	utils.Atoi(os.Getenv("SAMPLE_THREADS")), utils.Atoi(os.Getenv("TRIALS")), getNumbers,
	fmt.Sprintf(
		"results/aRPC/getrandomnumbers_%s_%d_threads", os.Getenv("VALUE"),
		utils.Atoi(os.Getenv("SAMPLE_THREADS")),
	)+"/server/%d.json",
)

// GetNumbers aRPC server function implementation
func (gs *ServerDefinition) GetNumbers(amount *getrandomnumbers2.Amount) (*getrandomnumbers2.NumberList, error) {
	var reqData interfaces.Serializable = amount
	respData := &getrandomnumbers2.NumberList{}

	response, err := metricsGetNumbers(reqData)
	if err == nil {
		respData = response.(*getrandomnumbers2.NumberList)
	}

	return respData, err
}

func main() {
	aRPCController := utils.SetupServer()
	getrandomnumbers2.RegisterGetrandomnumbersServer(aRPCController, &ServerDefinition{})
	utils.StartServer(aRPCController)
}
