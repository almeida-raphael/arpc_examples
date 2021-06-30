package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/almeida-raphael/arpc_examples/models/arpc/typeuint64"

	"github.com/almeida-raphael/arpc/interfaces"
	"github.com/almeida-raphael/arpc_examples/utils"
)

func main() {
	aRPCController := utils.SetupClient()

	service := typeuint64.NewTypeuint64(&aRPCController)
	err := aRPCController.StartClient()
	if err != nil {
		log.Fatal(err)
	}

	TypeUInt64 := func(req interfaces.Serializable) (interfaces.Serializable, error) {
		reqData := req.(*typeuint64.Request)
		response, err := service.TypeUInt64(reqData)
		if err != nil {
			return nil, err
		}

		return response, nil
	}

	err = utils.RunClientRPCAndCollectMetrics(
		utils.Atoi(os.Getenv("SAMPLE_THREADS")), utils.Atoi(os.Getenv("TRIALS")),
		&typeuint64.Request{Entries: uint64(rand.Int())}, TypeUInt64,
		fmt.Sprintf(
			"results/aRPC/type_uint64_%d_threads/client/%d.json",
			utils.Atoi(os.Getenv("SAMPLE_THREADS")), time.Now().UnixNano(),
		),
	)
	if err != nil {
		log.Fatal(err)
	}
}
