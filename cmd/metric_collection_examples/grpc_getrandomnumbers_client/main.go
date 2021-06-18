package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/almeida-raphael/arpc_examples/grpc/getrandomnumbers"

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
	client := getrandomnumbers.NewGetRandomNumbersClient(conn)

	requestData := getrandomnumbers.Amount{Value: value}

	GetNumbers := func(req proto.Message) (proto.Message, error) {
		reqData := req.(*getrandomnumbers.Amount)
		response, err := client.GetNumbers(context.Background(), reqData)
		if err != nil {
			return nil, err
		}

		return response, nil
	}

	err = utils.RunGRPCClientRPCAndCollectMetrics(
		20, 1000, &requestData, GetNumbers,
		fmt.Sprintf("results/getrandomnumbers_grpc_%d/client/%d.json", value, time.Now().UnixNano()),
	)
	if err != nil {
		log.Fatal(err)
	}
}
