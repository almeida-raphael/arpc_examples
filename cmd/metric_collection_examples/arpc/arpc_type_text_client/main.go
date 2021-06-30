package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/almeida-raphael/arpc_examples/models/arpc/typetext"

	"github.com/almeida-raphael/arpc/interfaces"
	"github.com/almeida-raphael/arpc_examples/utils"
)

func main() {
	aRPCController := utils.SetupClient()

	service := typetext.NewTypetext(&aRPCController)
	err := aRPCController.StartClient()
	if err != nil {
		log.Fatal(err)
	}

	TypeText := func(req interfaces.Serializable) (interfaces.Serializable, error) {
		reqData := req.(*typetext.Request)
		response, err := service.TypeText(reqData)
		if err != nil {
			return nil, err
		}

		return response, nil
	}

	err = utils.RunClientRPCAndCollectMetrics(
		utils.Atoi(os.Getenv("SAMPLE_THREADS")), utils.Atoi(os.Getenv("TRIALS")),
		&typetext.Request{Entries: utils.GenerateString(1024)}, TypeText,
		fmt.Sprintf(
			"results/aRPC/type_text_%d_threads/client/%d.json",
			utils.Atoi(os.Getenv("SAMPLE_THREADS")), time.Now().UnixNano(),
		),
	)
	if err != nil {
		log.Fatal(err)
	}
}
