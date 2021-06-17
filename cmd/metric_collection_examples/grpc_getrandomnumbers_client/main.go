package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/almeida-raphael/arpc_examples/grpc/getrandomnumbers"

	"google.golang.org/protobuf/proto"

	"github.com/almeida-raphael/arpc_examples/utils"
)

func main() {
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

	requestData := getrandomnumbers.Amount{Value: 1000}

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
		fmt.Sprintf("results/getrandomnumbers_grpc/client/%d.json", time.Now().UnixNano()),
	)
	if err != nil {
		log.Fatal(err)
	}
}
