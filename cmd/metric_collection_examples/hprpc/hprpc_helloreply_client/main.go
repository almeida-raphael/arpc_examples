package main

import (
	"fmt"
	"os"
	"time"

	helloreply2 "github.com/almeida-raphael/arpc_examples/models/hprpc/helloreply"

	"github.com/almeida-raphael/arpc/interfaces"
	"github.com/almeida-raphael/arpc_examples/utils"
)

func main() {
	aRPCController := utils.SetupClient()

	service := helloreply2.NewHelloreply(&aRPCController)
	err := aRPCController.StartClient()
	if err != nil {
		panic(err)
	}

	SayHello := func(req interfaces.Serializable) (interfaces.Serializable, error) {
		reqData := req.(*helloreply2.Text)
		response, err := service.SayHello(reqData)
		if err != nil {
			return nil, err
		}

		return response, nil
	}

	err = utils.RunClientRPCAndCollectMetrics(
		utils.Atoi(os.Getenv("SAMPLE_THREADS")), utils.Atoi(os.Getenv("TRIALS")),
		&helloreply2.Text{Data: utils.GenerateString(1000)}, SayHello,
		fmt.Sprintf(
			"results/aRPC/helloreply/client/%d.json",
			time.Now().UnixNano(),
		),
	)
	if err != nil {
		panic(err)
	}
}
