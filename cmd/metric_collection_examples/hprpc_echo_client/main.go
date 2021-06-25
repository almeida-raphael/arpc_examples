package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/almeida-raphael/arpc/interfaces"
	"github.com/almeida-raphael/arpc_examples/hprpc/echo"
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

	service := echo.NewEcho(&aRPCController)
	err := aRPCController.StartClient()
	if err != nil {
		log.Fatal(err)
	}

	Yell := func(req interfaces.Serializable) (interfaces.Serializable, error) {
		reqData := req.(*echo.Numbers)
		response, err := service.Yell(reqData)
		if err != nil {
			return nil, err
		}

		return response, nil
	}

	err = utils.RunClientRPCAndCollectMetrics(
		utils.Atoi(os.Getenv("SAMPLE_THREADS")), 100, &echo.Numbers{Entries: utils.GenerateNumbers(value)}, Yell,
		fmt.Sprintf("results/aRPC/largedata_%d/client/%d.json", value, time.Now().UnixNano()),
	)
	if err != nil {
		log.Fatal(err)
	}
}
