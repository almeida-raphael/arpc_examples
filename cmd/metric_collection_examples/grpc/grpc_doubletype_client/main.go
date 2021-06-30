package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	doubletype2 "github.com/almeida-raphael/arpc_examples/models/grpc/doubletype"

	"google.golang.org/protobuf/proto"

	"github.com/almeida-raphael/arpc_examples/utils"
)

func main() {
	valueStr := os.Getenv("VALUE")
	var value = 1000
	if valueStr != "" {
		if v, err := strconv.Atoi(valueStr); err == nil {
			value = v
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
	client := doubletype2.NewDoubleTypeClient(conn)

	requestData := doubletype2.NumberList{Entries: utils.GenerateNumbers(value)}

	Average := func(req proto.Message) (proto.Message, error) {
		reqData := req.(*doubletype2.NumberList)
		response, err := client.Average(context.Background(), reqData)
		if err != nil {
			return nil, err
		}

		return response, nil
	}

	err = utils.RunGRPCClientRPCAndCollectMetrics(
		utils.Atoi(os.Getenv("SAMPLE_THREADS")), utils.Atoi(os.Getenv("TRIALS")), &requestData, Average,
		fmt.Sprintf("results/gRPC/doubletype_%d_%d_threads/client/%d.json", value,
			utils.Atoi(os.Getenv("SAMPLE_THREADS")), time.Now().UnixNano()),
	)
	if err != nil {
		log.Fatal(err)
	}
}
