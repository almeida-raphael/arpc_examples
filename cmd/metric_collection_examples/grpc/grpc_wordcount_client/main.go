package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	wordcount2 "github.com/almeida-raphael/arpc_examples/models/grpc/wordcount"

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
	client := wordcount2.NewWordCountClient(conn)

	data := utils.LoadAsset()
	requestData := wordcount2.Text{
		Data: string(data),
	}

	CountWords := func(req proto.Message) (proto.Message, error) {
		reqData := req.(*wordcount2.Text)
		response, err := client.CountWords(context.Background(), reqData)
		if err != nil {
			return nil, err
		}

		return response, nil
	}

	err = utils.RunGRPCClientRPCAndCollectMetrics(
		utils.Atoi(os.Getenv("SAMPLE_THREADS")), utils.Atoi(os.Getenv("TRIALS")), &requestData,
		CountWords, fmt.Sprintf("results/gRPC/wordcount/client/%d.json", time.Now().UnixNano()),
	)
	if err != nil {
		log.Fatal(err)
	}
}
