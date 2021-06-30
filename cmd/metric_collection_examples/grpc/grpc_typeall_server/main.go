package main

import (
	"context"
	"fmt"
	"os"

	"github.com/almeida-raphael/arpc_examples/models/grpc/typeall"

	"github.com/almeida-raphael/arpc_examples/utils"
	"google.golang.org/protobuf/proto"
)

// ServerDefinition struct to implement greetings aRPC service procedures
type ServerDefinition struct {
	typeall.UnimplementedTypeAllServer
}

// typeAll gRPC server function implementation
func typeAll(request proto.Message) (proto.Message, error) {
	return request, nil
}

var metricsTypeAll = utils.CollectGRPCServerMetrics(
	utils.Atoi(os.Getenv("SAMPLE_THREADS")), utils.Atoi(os.Getenv("TRIALS")), typeAll,
	fmt.Sprintf(
		"results/gRPC/type_all_%d_threads",
		utils.Atoi(os.Getenv("SAMPLE_THREADS")),
	)+"/server/%d.json",
)

// TypeAll gRPC server function implementation
func (gs *ServerDefinition) TypeAll(ctx context.Context, request *typeall.Request) (*typeall.EmptyResult, error) {
	respData := &typeall.EmptyResult{}

	response, err := metricsTypeAll(request)
	if err == nil {
		respData = response.(*typeall.EmptyResult)
	}

	return respData, err
}

func main() {
	listener, gRPCServer := utils.SetupGRPCServer()
	typeall.RegisterTypeAllServer(gRPCServer, &ServerDefinition{})
	utils.StartGRPCServer(listener, gRPCServer)
}
