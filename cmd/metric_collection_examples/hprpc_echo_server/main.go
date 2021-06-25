package main

import (
	"fmt"
	"os"

	"github.com/almeida-raphael/arpc/interfaces"
	"github.com/almeida-raphael/arpc_examples/hprpc/echo"
	"github.com/almeida-raphael/arpc_examples/utils"
)

// ServerDefinition struct to implement greetings aRPC service procedures
type ServerDefinition struct{}

// yell aRPC server function implementation
func yell(request interfaces.Serializable) (interfaces.Serializable, error) {
	return request, nil
}

var metricsYell = utils.CollectServerMetrics(
	utils.Atoi(os.Getenv("SAMPLE_THREADS")), 1000, yell,
	fmt.Sprintf("results/aRPC/largedata_%s", os.Getenv("VALUE"))+"/server/%d.json",
)

// Yell aRPC server function implementation
func (gs *ServerDefinition) Yell(amount *echo.Numbers) (*echo.Numbers, error) {
	var reqData interfaces.Serializable = amount
	respData := &echo.Numbers{}

	response, err := metricsYell(reqData)
	if err == nil {
		respData = response.(*echo.Numbers)
	}

	return respData, err
}

func main() {
	aRPCController := utils.SetupServer()
	echo.RegisterEchoServer(aRPCController, &ServerDefinition{})
	utils.StartServer(aRPCController)
}
