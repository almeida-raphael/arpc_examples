package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	doubletype2 "github.com/almeida-raphael/arpc_examples/models/hprpc/doubletype"

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

	service := doubletype2.NewDoubletype(&aRPCController)
	err := aRPCController.StartClient()
	if err != nil {
		panic(err)
	}

	requestData := doubletype2.NumberList{Entries: utils.GenerateNumbers(value)}

	Average := func(req interfaces.Serializable) (interfaces.Serializable, error) {
		reqData := req.(*doubletype2.NumberList)
		response, err := service.Average(reqData)
		if err != nil {
			return nil, err
		}

		return response, nil
	}

	err = utils.RunClientRPCAndCollectMetrics(
		utils.Atoi(os.Getenv("SAMPLE_THREADS")), utils.Atoi(os.Getenv("TRIALS")), &requestData, Average,
		fmt.Sprintf(
			"results/aRPC/doubletype_%d_%d_threads/client/%d.json", value,
			utils.Atoi(os.Getenv("SAMPLE_THREADS")), time.Now().UnixNano(),
		),
	)
	if err != nil {
		log.Fatal(err)
	}
}
