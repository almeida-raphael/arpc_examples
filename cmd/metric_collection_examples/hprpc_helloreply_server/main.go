package main

import (
	"fmt"

	"github.com/almeida-raphael/arpc/interfaces"
	"github.com/almeida-raphael/arpc_examples/hprpc/helloreply"
	"github.com/almeida-raphael/arpc_examples/utils"
)

// ServerDefinition struct to implement greetings aRPC service procedures
type ServerDefinition struct{}

func sayHello(request interfaces.Serializable) (interfaces.Serializable, error) {
	reqData := request.(*helloreply.Text)
	return &helloreply.Text{Data: fmt.Sprintf("Hello %s", reqData.Data)}, nil
}

var metricsSayHello = utils.CollectServerMetrics(
	20, 1000, sayHello,
	"results/helloreply/server/%d.json",
)

// SayHello aRPC server function implementation
func (gs *ServerDefinition) SayHello(request *helloreply.Text) (*helloreply.Text, error) {
	var reqData interfaces.Serializable = request
	respData := &helloreply.Text{}

	response, err := metricsSayHello(reqData)
	if err == nil {
		respData = response.(*helloreply.Text)
	}

	return respData, err
}

func main() {
	aRPCController := utils.SetupServer()
	helloreply.RegisterHelloreplyServer(aRPCController, &ServerDefinition{})
	utils.StartServer(aRPCController)
}
