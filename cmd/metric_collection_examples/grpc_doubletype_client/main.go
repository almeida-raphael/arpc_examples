package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"google.golang.org/protobuf/proto"

	"github.com/almeida-raphael/arpc_examples/grpc/doubletype"
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
	client := doubletype.NewDoubleTypeClient(conn)

	requestData := doubletype.NumberList{Entries: utils.GenerateNumbers(value)}

	Average := func(req proto.Message) (proto.Message, error) {
		reqData := req.(*doubletype.NumberList)
		response, err := client.Average(context.Background(), reqData)
		if err != nil {
			return nil, err
		}

		return response, nil
	}

	err = utils.RunGRPCClientRPCAndCollectMetrics(
		20, 1000, &requestData, Average,
		fmt.Sprintf("results/doubletype_grpc_%d/client/%d.json", value, time.Now().UnixNano()),
	)
	if err != nil {
		log.Fatal(err)
	}
}
