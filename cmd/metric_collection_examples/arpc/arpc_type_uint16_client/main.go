package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/almeida-raphael/arpc_examples/models/arpc/typeuint16"

	"github.com/almeida-raphael/arpc/interfaces"
	"github.com/almeida-raphael/arpc_examples/utils"
)

func main() {
	aRPCController := utils.SetupClient()

	service := typeuint16.NewTypeuint16(&aRPCController)
	err := aRPCController.StartClient()
	if err != nil {
		log.Fatal(err)
	}

	TypeUInt16 := func(req interfaces.Serializable) (interfaces.Serializable, error) {
		reqData := req.(*typeuint16.Request)
		response, err := service.TypeUInt16(reqData)
		if err != nil {
			return nil, err
		}

		return response, nil
	}

	err = utils.RunClientRPCAndCollectMetrics(
		utils.Atoi(os.Getenv("SAMPLE_THREADS")), utils.Atoi(os.Getenv("TRIALS")),
		&typeuint16.Request{Entries: uint16(rand.Int())}, TypeUInt16,
		fmt.Sprintf(
			"results/aRPC/type_uint16_%d_threads/client/%d.json",
			utils.Atoi(os.Getenv("SAMPLE_THREADS")), time.Now().UnixNano(),
		),
	)
	if err != nil {
		log.Fatal(err)
	}
}
