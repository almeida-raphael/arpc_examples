package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/almeida-raphael/arpc_examples/models/arpc/typeint32"

	"github.com/almeida-raphael/arpc/interfaces"
	"github.com/almeida-raphael/arpc_examples/utils"
)

func main() {
	aRPCController := utils.SetupClient()

	service := typeint32.NewTypeint32(&aRPCController)
	err := aRPCController.StartClient()
	if err != nil {
		log.Fatal(err)
	}

	TypeInt32 := func(req interfaces.Serializable) (interfaces.Serializable, error) {
		reqData := req.(*typeint32.Request)
		response, err := service.TypeInt32(reqData)
		if err != nil {
			return nil, err
		}

		return response, nil
	}

	err = utils.RunClientRPCAndCollectMetrics(
		utils.Atoi(os.Getenv("SAMPLE_THREADS")), utils.Atoi(os.Getenv("TRIALS")),
		&typeint32.Request{Entries: int32(rand.Int())}, TypeInt32,
		fmt.Sprintf(
			"results/aRPC/type_int32_%d_threads/client/%d.json",
			utils.Atoi(os.Getenv("SAMPLE_THREADS")), time.Now().UnixNano(),
		),
	)
	if err != nil {
		log.Fatal(err)
	}
}
