package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	getrandomnumbers2 "github.com/almeida-raphael/arpc_examples/models/grpc/getrandomnumbers"

	"google.golang.org/protobuf/proto"

	"github.com/almeida-raphael/arpc_examples/utils"
)

func main() {
	valueStr := os.Getenv("VALUE")
	var value int32 = 1000
	if valueStr != "" {
		if v, err := strconv.Atoi(valueStr); err == nil {
			value = int32(v)
		} else {
			log.Fatal(err)
		}
	}

	conn, err := utils.SetupGRPCClient()
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	client := getrandomnumbers2.NewGetRandomNumbersClient(conn)

	GetNumbers := func(req proto.Message) (proto.Message, error) {
		reqData := req.(*getrandomnumbers2.Amount)
		response, err := client.GetNumbers(context.Background(), reqData)
		if err != nil {
			return nil, err
		}

		return response, nil
	}

	err = utils.RunGRPCClientRPCAndCollectMetrics(
		utils.Atoi(os.Getenv("SAMPLE_THREADS")), utils.Atoi(os.Getenv("TRIALS")),
		func() proto.Message { return &getrandomnumbers2.Amount{Value: value} }, GetNumbers,
		fmt.Sprintf(
			"results/gRPC/getrandomnumbers_%d_%d_threads/client/%d.json",
			value, utils.Atoi(os.Getenv("SAMPLE_THREADS")), time.Now().UnixNano(),
		),
	)
	if err != nil {
		log.Fatal(err)
	}
}
