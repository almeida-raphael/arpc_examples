package main

import (
	"context"
	"fmt"
	"os"

	echo2 "github.com/almeida-raphael/arpc_examples/models/grpc/echo"

	"github.com/almeida-raphael/arpc_examples/utils"
	"google.golang.org/protobuf/proto"
)

// ServerDefinition struct to implement greetings aRPC service procedures
type ServerDefinition struct {
	echo2.UnimplementedEchoServer
}

// yell gRPC server function implementation
func yell(request proto.Message) (proto.Message, error) {
	return request, nil
}

var metricsYell = utils.CollectGRPCServerMetrics(
	utils.Atoi(os.Getenv("SAMPLE_THREADS")), utils.Atoi(os.Getenv("TRIALS")), yell,
	fmt.Sprintf(
		"results/gRPC/largedata_%s_%d_threads", os.Getenv("VALUE"),
		utils.Atoi(os.Getenv("SAMPLE_THREADS")),
	)+"/server/%d.json",
)

// Yell gRPC server function implementation
func (gs *ServerDefinition) Yell(ctx context.Context, request *echo2.Numbers) (*echo2.Numbers, error) {
	respEntries := &echo2.Numbers{}

	response, err := metricsYell(request)
	if err == nil {
		respEntries = response.(*echo2.Numbers)
	}

	return respEntries, err
}

func main() {
	listener, gRPCServer := utils.SetupGRPCServer()
	echo2.RegisterEchoServer(gRPCServer, &ServerDefinition{})
	utils.StartGRPCServer(listener, gRPCServer)
}
