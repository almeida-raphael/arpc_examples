package main

import (
	"context"
	"fmt"
	"os"

	"github.com/almeida-raphael/arpc_examples/grpc/echo"

	"github.com/almeida-raphael/arpc_examples/utils"
	"google.golang.org/protobuf/proto"
)

// ServerDefinition struct to implement greetings aRPC service procedures
type ServerDefinition struct {
	echo.UnimplementedEchoServer
}

// yell gRPC server function implementation
func yell(request proto.Message) (proto.Message, error) {
	return request, nil
}

var metricsYell = utils.CollectGRPCServerMetrics(
	utils.Atoi(os.Getenv("SAMPLE_THREADS")), 1000, yell,
	fmt.Sprintf("results/gRPC/largedata_%s", os.Getenv("VALUE"))+"/server/%d.json",
)

// Yell gRPC server function implementation
func (gs *ServerDefinition) Yell(ctx context.Context, request *echo.Numbers) (*echo.Numbers, error) {
	respEntries := &echo.Numbers{}

	response, err := metricsYell(request)
	if err == nil {
		respEntries = response.(*echo.Numbers)
	}

	return respEntries, err
}

func main() {
	listener, gRPCServer := utils.SetupGRPCServer()
	echo.RegisterEchoServer(gRPCServer, &ServerDefinition{})
	utils.StartGRPCServer(listener, gRPCServer)
}
