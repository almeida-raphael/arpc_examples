package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	echo2 "github.com/almeida-raphael/arpc_examples/models/hprpc/echo"

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

	service := echo2.NewEcho(&aRPCController)
	err := aRPCController.StartClient()
	if err != nil {
		log.Fatal(err)
	}

	Yell := func(req interfaces.Serializable) (interfaces.Serializable, error) {
		reqData := req.(*echo2.Numbers)
		response, err := service.Yell(reqData)
		if err != nil {
			return nil, err
		}

		return response, nil
	}

	err = utils.RunClientRPCAndCollectMetrics(
		utils.Atoi(os.Getenv("SAMPLE_THREADS")), utils.Atoi(os.Getenv("TRIALS")),
		&echo2.Numbers{Entries: utils.GenerateNumbers(value)}, Yell,
		fmt.Sprintf(
			"results/aRPC/largedata_%d_%d_threads/client/%d.json", value,
			utils.Atoi(os.Getenv("SAMPLE_THREADS")), time.Now().UnixNano(),
		),
	)
	if err != nil {
		log.Fatal(err)
	}
}
