package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/almeida-raphael/arpc/interfaces"
	"github.com/almeida-raphael/arpc_examples/hprpc/doubletype"
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

	service := doubletype.NewDoubletype(&aRPCController)
	err := aRPCController.StartClient()
	if err != nil {
		panic(err)
	}

	requestData := doubletype.NumberList{Entries: utils.GenerateNumbers(value)}

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
		fmt.Sprintf("results/doubletype_%d/client/%d.json", value, time.Now().UnixNano()),
	)
	if err != nil {
		log.Fatal(err)
	}
}
