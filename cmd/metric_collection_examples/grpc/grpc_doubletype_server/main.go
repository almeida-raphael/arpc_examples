package main

import (
	"context"
	"fmt"
	"os"

	doubletype2 "github.com/almeida-raphael/arpc_examples/models/grpc/doubletype"

	"github.com/almeida-raphael/arpc_examples/utils"
	"google.golang.org/protobuf/proto"
)

// ServerDefinition struct to implement greetings aRPC service procedures
type ServerDefinition struct {
	doubletype2.UnimplementedDoubleTypeServer
}

// getNumbers aRPC server function implementation
func average(request proto.Message) (proto.Message, error) {
	reqData := request.(*doubletype2.NumberList)
	return &doubletype2.Result{Value: utils.Average(reqData.Entries)}, nil
}

var metricsAverage = utils.CollectGRPCServerMetrics(
	utils.Atoi(os.Getenv("SAMPLE_THREADS")), utils.Atoi(os.Getenv("TRIALS")), average,
	fmt.Sprintf(
		"results/gRPC/doubletype_%s_%d_threads", os.Getenv("VALUE"),
		utils.Atoi(os.Getenv("SAMPLE_THREADS")),
	)+"/server/%d.json",
)

// Average aRPC server function implementation
func (gs *ServerDefinition) Average(ctx context.Context, request *doubletype2.NumberList) (*doubletype2.Result, error) {
	respData := &doubletype2.Result{}

	response, err := metricsAverage(request)
	if err == nil {
		respData = response.(*doubletype2.Result)
	}

	return respData, err
}

func main() {
	listener, gRPCServer := utils.SetupGRPCServer()
	doubletype2.RegisterDoubleTypeServer(gRPCServer, &ServerDefinition{})
	utils.StartGRPCServer(listener, gRPCServer)
}
