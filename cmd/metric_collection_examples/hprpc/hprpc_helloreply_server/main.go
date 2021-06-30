package main

import (
	"fmt"
	"os"

	helloreply2 "github.com/almeida-raphael/arpc_examples/models/hprpc/helloreply"

	"github.com/almeida-raphael/arpc/interfaces"
	"github.com/almeida-raphael/arpc_examples/utils"
)

// ServerDefinition struct to implement greetings aRPC service procedures
type ServerDefinition struct{}

func sayHello(request interfaces.Serializable) (interfaces.Serializable, error) {
	reqData := request.(*helloreply2.Text)
	return &helloreply2.Text{Data: fmt.Sprintf("Hello %s", reqData.Data)}, nil
}

var metricsSayHello = utils.CollectServerMetrics(
	utils.Atoi(os.Getenv("SAMPLE_THREADS")), utils.Atoi(os.Getenv("TRIALS")), sayHello,
	"results/aRPC/helloreply/server/%d.json",
)

// SayHello aRPC server function implementation
func (gs *ServerDefinition) SayHello(request *helloreply2.Text) (*helloreply2.Text, error) {
	var reqData interfaces.Serializable = request
	respData := &helloreply2.Text{}

	response, err := metricsSayHello(reqData)
	if err == nil {
		respData = response.(*helloreply2.Text)
	}

	return respData, err
}

func main() {
	aRPCController := utils.SetupServer()
	helloreply2.RegisterHelloreplyServer(aRPCController, &ServerDefinition{})
	utils.StartServer(aRPCController)
}
