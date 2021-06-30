package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	getrandomnumbers2 "github.com/almeida-raphael/arpc_examples/models/hprpc/getrandomnumbers"

	"github.com/almeida-raphael/arpc/interfaces"
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
	service := getrandomnumbers2.NewGetrandomnumbers(&aRPCController)
	err := aRPCController.StartClient()
	if err != nil {
		log.Fatal(err)
	}

	requestData := getrandomnumbers2.Amount{Value: value}

	GetNumbers := func(req interfaces.Serializable) (interfaces.Serializable, error) {
		reqData := req.(*getrandomnumbers2.Amount)
		response, err := service.GetNumbers(reqData)
		if err != nil {
			return nil, err
		}

		return response, nil
	}

	err = utils.RunClientRPCAndCollectMetrics(
		utils.Atoi(os.Getenv("SAMPLE_THREADS")), utils.Atoi(os.Getenv("TRIALS")), &requestData, GetNumbers,
		fmt.Sprintf(
			"results/aRPC/getrandomnumbers_%d_%d_threads/client/%d.json", value,
			utils.Atoi(os.Getenv("SAMPLE_THREADS")), time.Now().UnixNano(),
		),
	)
	if err != nil {
		log.Fatal(err)
	}
}
