package main

import (
	"fmt"
	"os"

	"github.com/almeida-raphael/arpc/interfaces"
	"github.com/almeida-raphael/arpc_examples/hprpc/getrandomnumbers"
	"github.com/almeida-raphael/arpc_examples/utils"
)

// ServerDefinition struct to implement greetings aRPC service procedures
type ServerDefinition struct{}

// getNumbers aRPC server function implementation
func getNumbers(request interfaces.Serializable) (interfaces.Serializable, error) {
	reqData := request.(*getrandomnumbers.Amount)
	return &getrandomnumbers.NumberList{Entries: utils.GenerateNumbers(int(reqData.Value))}, nil
}

var metricsGetNumbers = utils.CollectServerMetrics(
	20, 1000, getNumbers,
	fmt.Sprintf("results/aRPC/getrandomnumbers_%s", os.Getenv("VALUE"))+"/server/%d.json",
)

// GetNumbers aRPC server function implementation
func (gs *ServerDefinition) GetNumbers(amount *getrandomnumbers.Amount) (*getrandomnumbers.NumberList, error) {
	var reqData interfaces.Serializable = amount
	respData := &getrandomnumbers.NumberList{}

	response, err := metricsGetNumbers(reqData)
	if err == nil {
		respData = response.(*getrandomnumbers.NumberList)
	}

	return respData, err
}

func main() {
	aRPCController := utils.SetupServer()
	getrandomnumbers.RegisterGetrandomnumbersServer(aRPCController, &ServerDefinition{})
	utils.StartServer(aRPCController)
}
