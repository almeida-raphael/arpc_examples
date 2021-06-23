package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/almeida-raphael/arpc_examples/grpc/helloreply"

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
	client := helloreply.NewHelloReplyClient(conn)

	SayHello := func(req proto.Message) (proto.Message, error) {
		reqData := req.(*helloreply.Text)
		response, err := client.SayHello(context.Background(), reqData)
		if err != nil {
			return nil, err
		}

		return response, nil
	}

	err = utils.RunGRPCClientRPCAndCollectMetrics(
		20, 1000, &helloreply.Text{Data: utils.GenerateString(1000)}, SayHello,
		fmt.Sprintf("results/gRPC/helloreply/client/%d.json", time.Now().UnixNano()),
	)
	if err != nil {
		log.Fatal(err)
	}
}
