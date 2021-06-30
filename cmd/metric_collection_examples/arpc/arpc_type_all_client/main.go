package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/almeida-raphael/arpc_examples/models/arpc/typeall"

	"github.com/almeida-raphael/arpc/interfaces"
	"github.com/almeida-raphael/arpc_examples/utils"
)

func main() {
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

	err = utils.RunClientRPCAndCollectMetrics(
		utils.Atoi(os.Getenv("SAMPLE_THREADS")), utils.Atoi(os.Getenv("TRIALS")),
		&typeall.Request{
			Binary:    []byte(utils.GenerateString(1024)),
			Bool:      true,
			Float32:   rand.Float32(),
			Float64:   rand.Float64(),
			Int32:     int32(rand.Uint32()),
			Int64:     int64(rand.Uint64()),
			Text:      utils.GenerateString(1024),
			Timestamp: time.Now(),
			Uint8:     uint8(rand.Int()),
			Uint16:    uint16(rand.Int()),
			Uint32:    rand.Uint32(),
			Uint64:    rand.Uint64(),
		}, TypeAll,
		fmt.Sprintf(
			"results/aRPC/type_all_%d_threads/client/%d.json",
			utils.Atoi(os.Getenv("SAMPLE_THREADS")), time.Now().UnixNano(),
		),
	)
	if err != nil {
		log.Fatal(err)
	}
}
