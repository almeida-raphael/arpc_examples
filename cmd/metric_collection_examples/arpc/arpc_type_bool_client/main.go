package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/almeida-raphael/arpc_examples/models/arpc/typebool"

	"github.com/almeida-raphael/arpc/interfaces"
	"github.com/almeida-raphael/arpc_examples/utils"
)

func main() {
	aRPCController := utils.SetupClient()

	service := typebool.NewTypebool(&aRPCController)
	err := aRPCController.StartClient()
	if err != nil {
		log.Fatal(err)
	}

	TypeBool := func(req interfaces.Serializable) (interfaces.Serializable, error) {
		reqData := req.(*typebool.Request)
		response, err := service.TypeBool(reqData)
		if err != nil {
			return nil, err
		}

		return response, nil
	}

	err = utils.RunClientRPCAndCollectMetrics(
		utils.Atoi(os.Getenv("SAMPLE_THREADS")), utils.Atoi(os.Getenv("TRIALS")),
		&typebool.Request{Entries: true}, TypeBool,
		fmt.Sprintf(
			"results/aRPC/type_bool_%d_threads/client/%d.json",
			utils.Atoi(os.Getenv("SAMPLE_THREADS")), time.Now().UnixNano(),
		),
	)
	if err != nil {
		log.Fatal(err)
	}
}
