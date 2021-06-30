package main

import (
	"context"
	"fmt"
	"os"

	getrandomnumbers2 "github.com/almeida-raphael/arpc_examples/models/grpc/getrandomnumbers"

	"github.com/almeida-raphael/arpc_examples/utils"
	"google.golang.org/protobuf/proto"
)

// ServerDefinition struct to implement greetings aRPC service procedures
type ServerDefinition struct {
	getrandomnumbers2.UnimplementedGetRandomNumbersServer
}

// getNumbers gRPC server function implementation
func getNumbers(request proto.Message) (proto.Message, error) {
	reqData := request.(*getrandomnumbers2.Amount)
	return &getrandomnumbers2.NumbersList{Entries: utils.GenerateNumbers(int(reqData.Value))}, nil
}

var metricsGetNumbers = utils.CollectGRPCServerMetrics(
	utils.Atoi(os.Getenv("SAMPLE_THREADS")), utils.Atoi(os.Getenv("TRIALS")), getNumbers,
	fmt.Sprintf(
		"results/gRPC/getrandomnumbers_%s_%d_threads", os.Getenv("VALUE"),
		utils.Atoi(os.Getenv("SAMPLE_THREADS")),
	)+"/server/%d.json",
)

// GetNumbers aRPC server function implementation
func (gs *ServerDefinition) GetNumbers(ctx context.Context, request *getrandomnumbers2.Amount) (*getrandomnumbers2.NumbersList, error) {
	respData := &getrandomnumbers2.NumbersList{}

	response, err := metricsGetNumbers(request)
	if err == nil {
		respData = response.(*getrandomnumbers2.NumbersList)
	}

	return respData, err
}

func main() {
	listener, gRPCServer := utils.SetupGRPCServer()
	getrandomnumbers2.RegisterGetRandomNumbersServer(gRPCServer, &ServerDefinition{})
	utils.StartGRPCServer(listener, gRPCServer)
}
