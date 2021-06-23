package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/almeida-raphael/arpc/interfaces"
	"github.com/almeida-raphael/arpc_examples/hprpc/getrandomnumbers"
	"github.com/almeida-raphael/arpc_examples/utils"
)

func main() {
	valueStr := os.Getenv("VALUE")
	var value int32 = 1000
	if valueStr != "" {
		if v, err := strconv.Atoi(valueStr); err == nil {
			value = int32(v)
		} else {
			log.Fatal(err)
		}
	}

	aRPCController := utils.SetupClient()

	service := getrandomnumbers.NewGetrandomnumbers(&aRPCController)
	err := aRPCController.StartClient()
	if err != nil {
		log.Fatal(err)
	}

	requestData := getrandomnumbers.Amount{Value: value}

	GetNumbers := func(req interfaces.Serializable) (interfaces.Serializable, error) {
		reqData := req.(*getrandomnumbers.Amount)
		response, err := service.GetNumbers(reqData)
		if err != nil {
			return nil, err
		}

		return response, nil
	}

	err = utils.RunClientRPCAndCollectMetrics(
		20, 1000, &requestData, GetNumbers,
		fmt.Sprintf("results/aRPC/getrandomnumbers_%d/client/%d.json", value, time.Now().UnixNano()),
	)
	if err != nil {
		log.Fatal(err)
	}
}
