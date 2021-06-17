package main

import (
	"fmt"
	"log"
	"time"

	"github.com/almeida-raphael/arpc/interfaces"
	"github.com/almeida-raphael/arpc_examples/hprpc/doubletype"
	"github.com/almeida-raphael/arpc_examples/utils"
)

func main() {
	aRPCController := utils.SetupClient()

	service := doubletype.NewDoubletype(&aRPCController)
	err := aRPCController.StartClient()
	if err != nil {
		panic(err)
	}

	requestData := doubletype.NumberList{Entries: utils.GenerateNumbers(1000)}

	Average := func(req interfaces.Serializable) (interfaces.Serializable, error) {
		reqData := req.(*doubletype.NumberList)
		response, err := service.Average(reqData)
		if err != nil {
			return nil, err
		}

		return response, nil
	}

	err = utils.RunClientRPCAndCollectMetrics(
		20, 1000, &requestData, Average,
		fmt.Sprintf("results/doubletype/client/%d.json", time.Now().UnixNano()),
	)
	if err != nil {
		log.Fatal(err)
	}
}
