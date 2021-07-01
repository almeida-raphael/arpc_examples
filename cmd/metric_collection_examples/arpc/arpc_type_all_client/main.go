package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/almeida-raphael/arpc_examples/models/arpc/typeall"

	"github.com/almeida-raphael/arpc/interfaces"
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

	aRPCController := utils.SetupClient()

	service := typeall.NewTypeall(&aRPCController)
	err := aRPCController.StartClient()
	if err != nil {
		log.Fatal(err)
	}

	TypeAll := func(req interfaces.Serializable) (interfaces.Serializable, error) {
		reqData := req.(*typeall.Request)
		response, err := service.TypeAll(reqData)
		if err != nil {
			return nil, err
		}

		return response, nil
	}

	requests := &typeall.Request{
		TypesAll: make([]*typeall.TypesAll, value),
	}

	for idx := 0; idx < value; idx++ {
		requests.TypesAll[idx] = &typeall.TypesAll{
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

	err = utils.RunClientRPCAndCollectMetrics(
		utils.Atoi(os.Getenv("SAMPLE_THREADS")), utils.Atoi(os.Getenv("TRIALS")),
		requests, TypeAll,
		fmt.Sprintf(
			"results/aRPC/type_all_%d_%d_threads/client/%d.json",
			value, utils.Atoi(os.Getenv("SAMPLE_THREADS")), time.Now().UnixNano(),
		),
	)
	if err != nil {
		log.Fatal(err)
	}
}
