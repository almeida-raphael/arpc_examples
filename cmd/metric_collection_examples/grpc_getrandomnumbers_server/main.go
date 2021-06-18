package main

import (
	"context"
	"fmt"
	"os"

	"github.com/almeida-raphael/arpc_examples/grpc/getrandomnumbers"

	"github.com/almeida-raphael/arpc_examples/utils"
	"google.golang.org/protobuf/proto"
)

// ServerDefinition struct to implement greetings aRPC service procedures
type ServerDefinition struct {
	getrandomnumbers.UnimplementedGetRandomNumbersServer
}

// getNumbers gRPC server function implementation
func getNumbers(request proto.Message) (proto.Message, error) {
	reqData := request.(*getrandomnumbers.Amount)
	return &getrandomnumbers.NumbersList{Entries: utils.GenerateNumbers(int(reqData.Value))}, nil
}

var metricsGetNumbers = utils.CollectGRPCServerMetrics(
	20, 1000, getNumbers,
	fmt.Sprintf("results/getrandomnumbers_grpc_%s", os.Getenv("VALUE"))+"/server/%d.json",
)

// GetNumbers aRPC server function implementation
func (gs *ServerDefinition) GetNumbers(ctx context.Context, request *getrandomnumbers.Amount) (*getrandomnumbers.NumbersList, error) {
	respData := &getrandomnumbers.NumbersList{}

	response, err := metricsGetNumbers(request)
	if err == nil {
		respData = response.(*getrandomnumbers.NumbersList)
	}

	return respData, err
}

func main() {
	listener, gRPCServer := utils.SetupGRPCServer()
	getrandomnumbers.RegisterGetRandomNumbersServer(gRPCServer, &ServerDefinition{})
	utils.StartGRPCServer(listener, gRPCServer)
}
