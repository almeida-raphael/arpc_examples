package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/almeida-raphael/arpc_examples/models/grpc/typeall"

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
	client := typeall.NewTypeAllClient(conn)

	TypeAll := func(req proto.Message) (proto.Message, error) {
		reqData := req.(*typeall.Request)
		response, err := client.TypeAll(context.Background(), reqData)
		if err != nil {
			return nil, err
		}

		return response, nil
	}

	requestData := &typeall.Request{
		TypesAll: make([]*typeall.TypesAll, value),
	}

	for idx := 0; idx < value; idx++ {
		requestData.TypesAll[idx] = &typeall.TypesAll{
			Binary:  []byte(utils.GenerateString(1024)),
			Bool:    true,
			Float32: rand.Float32(),
			Float64: rand.Float64(),
			Int32:   int32(rand.Uint32()),
			Int64:   int64(rand.Uint64()),
			Text:    utils.GenerateString(1024),
			Uint32:  rand.Uint32(),
			Uint64:  rand.Uint64(),
		}
	}

	err = utils.RunGRPCClientRPCAndCollectMetrics(
		utils.Atoi(os.Getenv("SAMPLE_THREADS")), utils.Atoi(os.Getenv("TRIALS")),
		requestData, TypeAll,
		fmt.Sprintf(
			"results/gRPC/type_all_%d_%d_threads/client/%d.json",
			value, utils.Atoi(os.Getenv("SAMPLE_THREADS")), time.Now().UnixNano(),
		),
	)
	if err != nil {
		log.Fatal(err)
	}
}
