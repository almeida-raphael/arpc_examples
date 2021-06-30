package main

import (
	"fmt"
	"os"

	echo2 "github.com/almeida-raphael/arpc_examples/models/hprpc/echo"

	"github.com/almeida-raphael/arpc/interfaces"
	"github.com/almeida-raphael/arpc_examples/utils"
)

// ServerDefinition struct to implement greetings aRPC service procedures
type ServerDefinition struct{}

// yell aRPC server function implementation
func yell(request interfaces.Serializable) (interfaces.Serializable, error) {
	return request, nil
}

var metricsYell = utils.CollectServerMetrics(
	utils.Atoi(os.Getenv("SAMPLE_THREADS")), utils.Atoi(os.Getenv("TRIALS")), yell,
	fmt.Sprintf(
		"results/aRPC/largedata_%s_%d_threads", os.Getenv("VALUE"),
		utils.Atoi(os.Getenv("SAMPLE_THREADS")),
	)+"/server/%d.json",
)

// Yell aRPC server function implementation
func (gs *ServerDefinition) Yell(amount *echo2.Numbers) (*echo2.Numbers, error) {
	var reqData interfaces.Serializable = amount
	respData := &echo2.Numbers{}

	response, err := metricsYell(reqData)
	if err == nil {
		respData = response.(*echo2.Numbers)
	}

	return respData, err
}

func main() {
	aRPCController := utils.SetupServer()
	echo2.RegisterEchoServer(aRPCController, &ServerDefinition{})
	utils.StartServer(aRPCController)
}
