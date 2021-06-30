package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/almeida-raphael/arpc_examples/models/arpc/typebinary"

	"github.com/almeida-raphael/arpc/interfaces"
	"github.com/almeida-raphael/arpc_examples/utils"
)

func main() {
	aRPCController := utils.SetupClient()

	service := typebinary.NewTypebinary(&aRPCController)
	err := aRPCController.StartClient()
	if err != nil {
		log.Fatal(err)
	}

	TypeBinary := func(req interfaces.Serializable) (interfaces.Serializable, error) {
		reqData := req.(*typebinary.Request)
		response, err := service.TypeBinary(reqData)
		if err != nil {
			return nil, err
		}

		return response, nil
	}

	err = utils.RunClientRPCAndCollectMetrics(
		utils.Atoi(os.Getenv("SAMPLE_THREADS")), utils.Atoi(os.Getenv("TRIALS")),
		&typebinary.Request{Entries: []byte(utils.GenerateString(1024))}, TypeBinary,
		fmt.Sprintf(
			"results/aRPC/type_binary_%d_threads/client/%d.json",
			utils.Atoi(os.Getenv("SAMPLE_THREADS")), time.Now().UnixNano(),
		),
	)
	if err != nil {
		log.Fatal(err)
	}
}
