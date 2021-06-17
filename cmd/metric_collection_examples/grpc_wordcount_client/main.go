package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/almeida-raphael/arpc_examples/grpc/wordcount"

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
	client := wordcount.NewWordCountClient(conn)

	data := utils.LoadAsset()
	requestData := wordcount.Text{
		Data: string(data),
	}

	CountWords := func(req proto.Message) (proto.Message, error) {
		reqData := req.(*wordcount.Text)
		response, err := client.CountWords(context.Background(), reqData)
		if err != nil {
			return nil, err
		}

		return response, nil
	}

	err = utils.RunGRPCClientRPCAndCollectMetrics(
		20, 1000, &requestData, CountWords,
		fmt.Sprintf("results/wordcount_grpc/client/%d.json", time.Now().UnixNano()),
	)
	if err != nil {
		log.Fatal(err)
	}
}
