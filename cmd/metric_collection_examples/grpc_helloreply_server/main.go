package main

import (
	"context"
	"fmt"

	"github.com/almeida-raphael/arpc_examples/grpc/helloreply"

	"github.com/almeida-raphael/arpc_examples/utils"
	"google.golang.org/protobuf/proto"
)

// ServerDefinition struct to implement greetings aRPC service procedures
type ServerDefinition struct {
	helloreply.UnimplementedHelloReplyServer
}

// sayHello gRPC server function implementation
func sayHello(request proto.Message) (proto.Message, error) {
	reqData := request.(*helloreply.Text)
	return &helloreply.Text{Data: fmt.Sprintf("Hello %s", reqData.Data)}, nil
}

var metricsSayHello = utils.CollectGRPCServerMetrics(
	20, 1000, sayHello,
	"results/helloreply_grpc/server/%d.json",
)

// SayHello gRPC server function implementation
func (gs *ServerDefinition) SayHello(ctx context.Context, request *helloreply.Text) (*helloreply.Text, error) {
	respData := &helloreply.Text{}

	response, err := metricsSayHello(request)
	if err == nil {
		respData = response.(*helloreply.Text)
	}

	return respData, err
}

func main() {
	listener, gRPCServer := utils.SetupGRPCServer()
	helloreply.RegisterHelloReplyServer(gRPCServer, &ServerDefinition{})
	utils.StartGRPCServer(listener, gRPCServer)
}
