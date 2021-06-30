package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/almeida-raphael/arpc_examples/models/arpc/typetimestamp"

	"github.com/almeida-raphael/arpc/interfaces"
	"github.com/almeida-raphael/arpc_examples/utils"
)

func main() {
	aRPCController := utils.SetupClient()

	service := typetimestamp.NewTypetimestamp(&aRPCController)
	err := aRPCController.StartClient()
	if err != nil {
		log.Fatal(err)
	}

	TypeTimestamp := func(req interfaces.Serializable) (interfaces.Serializable, error) {
		reqData := req.(*typetimestamp.Request)
		response, err := service.TypeTimestamp(reqData)
		if err != nil {
			return nil, err
		}

		return response, nil
	}

	err = utils.RunClientRPCAndCollectMetrics(
		utils.Atoi(os.Getenv("SAMPLE_THREADS")), utils.Atoi(os.Getenv("TRIALS")),
		&typetimestamp.Request{Entries: time.Now()}, TypeTimestamp,
		fmt.Sprintf(
			"results/aRPC/type_timestamp_%d_threads/client/%d.json",
			utils.Atoi(os.Getenv("SAMPLE_THREADS")), time.Now().UnixNano(),
		),
	)
	if err != nil {
		log.Fatal(err)
	}
}
