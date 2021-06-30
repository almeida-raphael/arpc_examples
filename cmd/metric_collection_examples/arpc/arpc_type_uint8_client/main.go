package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/almeida-raphael/arpc_examples/models/arpc/typeuint8"

	"github.com/almeida-raphael/arpc/interfaces"
	"github.com/almeida-raphael/arpc_examples/utils"
)

func main() {
	aRPCController := utils.SetupClient()

	service := typeuint8.NewTypeuint8(&aRPCController)
	err := aRPCController.StartClient()
	if err != nil {
		log.Fatal(err)
	}

	TypeUInt8 := func(req interfaces.Serializable) (interfaces.Serializable, error) {
		reqData := req.(*typeuint8.Request)
		response, err := service.TypeUInt8(reqData)
		if err != nil {
			return nil, err
		}

		return response, nil
	}

	err = utils.RunClientRPCAndCollectMetrics(
		utils.Atoi(os.Getenv("SAMPLE_THREADS")), utils.Atoi(os.Getenv("TRIALS")),
		&typeuint8.Request{Entries: uint8(rand.Int())}, TypeUInt8,
		fmt.Sprintf(
			"results/aRPC/type_uint8_%d_threads/client/%d.json",
			utils.Atoi(os.Getenv("SAMPLE_THREADS")), time.Now().UnixNano(),
		),
	)
	if err != nil {
		log.Fatal(err)
	}
}
