package main

import (
	"context"

	"github.com/almeida-raphael/arpc_examples/grpc/doubletype"
	"github.com/almeida-raphael/arpc_examples/utils"
	"google.golang.org/protobuf/proto"
)

// ServerDefinition struct to implement greetings aRPC service procedures
type ServerDefinition struct {
	doubletype.UnimplementedDoubleTypeServer
}

// getNumbers aRPC server function implementation
func average(request proto.Message) (proto.Message, error) {
	reqData := request.(*doubletype.NumberList)
	return &doubletype.Result{Value: utils.Average(reqData.Entries)}, nil
}

var metricsAverage = utils.CollectGRPCServerMetrics(
	20, 1000, average,
	"results/doubletype_grpc/server/%d.json",
)

// Average aRPC server function implementation
func (gs *ServerDefinition) Average(ctx context.Context, request *doubletype.NumberList) (*doubletype.Result, error) {
	respData := &doubletype.Result{}

	response, err := metricsAverage(request)
	if err == nil {
		respData = response.(*doubletype.Result)
	}

	return respData, err
}

func main() {
	listener, gRPCServer := utils.SetupGRPCServer()
	doubletype.RegisterDoubleTypeServer(gRPCServer, &ServerDefinition{})
	utils.StartGRPCServer(listener, gRPCServer)
}
