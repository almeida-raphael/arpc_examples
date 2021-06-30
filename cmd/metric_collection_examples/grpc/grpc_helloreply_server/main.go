package main

import (
	"context"
	"fmt"
	"os"

	helloreply2 "github.com/almeida-raphael/arpc_examples/models/grpc/helloreply"

	"github.com/almeida-raphael/arpc_examples/utils"
	"google.golang.org/protobuf/proto"
)

// ServerDefinition struct to implement greetings aRPC service procedures
type ServerDefinition struct {
	helloreply2.UnimplementedHelloReplyServer
}

// sayHello gRPC server function implementation
func sayHello(request proto.Message) (proto.Message, error) {
	reqData := request.(*helloreply2.Text)
	return &helloreply2.Text{Data: fmt.Sprintf("Hello %s", reqData.Data)}, nil
}

var metricsSayHello = utils.CollectGRPCServerMetrics(
	utils.Atoi(os.Getenv("SAMPLE_THREADS")), utils.Atoi(os.Getenv("TRIALS")), sayHello,
	"results/gRPC/helloreply/server/%d.json",
)

// SayHello gRPC server function implementation
func (gs *ServerDefinition) SayHello(ctx context.Context, request *helloreply2.Text) (*helloreply2.Text, error) {
	respData := &helloreply2.Text{}

	response, err := metricsSayHello(request)
	if err == nil {
		respData = response.(*helloreply2.Text)
	}

	return respData, err
}

func main() {
	listener, gRPCServer := utils.SetupGRPCServer()
	helloreply2.RegisterHelloReplyServer(gRPCServer, &ServerDefinition{})
	utils.StartGRPCServer(listener, gRPCServer)
}
