package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	echo2 "github.com/almeida-raphael/arpc_examples/models/grpc/echo"

	"google.golang.org/protobuf/proto"

	"github.com/almeida-raphael/arpc_examples/utils"
)

func main() {
	valueStr := os.Getenv("VALUE")
	value := 1000
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
	client := echo2.NewEchoClient(conn)

	Yell := func(req proto.Message) (proto.Message, error) {
		reqData := req.(*echo2.Numbers)
		response, err := client.Yell(context.Background(), reqData)
		if err != nil {
			return nil, err
		}

		return response, nil
	}

	err = utils.RunGRPCClientRPCAndCollectMetrics(
		utils.Atoi(os.Getenv("SAMPLE_THREADS")), utils.Atoi(os.Getenv("TRIALS")),
		&echo2.Numbers{Entries: utils.GenerateNumbers(value)}, Yell,
		fmt.Sprintf(
			"results/gRPC/largedata_%d_%d_threads/client/%d.json", value,
			utils.Atoi(os.Getenv("SAMPLE_THREADS")), time.Now().UnixNano()),
	)
	if err != nil {
		log.Fatal(err)
	}
}
