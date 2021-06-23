package main

import (
	"fmt"
	"time"

	"github.com/almeida-raphael/arpc/interfaces"
	"github.com/almeida-raphael/arpc_examples/examples/wordcount"
	"github.com/almeida-raphael/arpc_examples/utils"
)

func main() {
	aRPCController := utils.SetupClient()

	service := wordcount.NewWordcount(&aRPCController)
	err := aRPCController.StartClient()
	if err != nil {
		panic(err)
	}

	data := utils.LoadAsset()
	requestData := wordcount.Text{
		Data: string(data),
	}

	CountWords := func(req interfaces.Serializable) (interfaces.Serializable, error) {
		reqData := req.(*wordcount.Text)
		response, err := service.CountWords(reqData)
		if err != nil {
			return nil, err
		}

		return response, nil
	}

	if !utils.HandleRemoteError(err) {
		fmt.Printf("Response:\n")
	}

	err = utils.RunClientRPCAndCollectMetrics(
		20, 1000, &requestData, CountWords,
		fmt.Sprintf("results/aRPC/wordcount/client/%d.json", time.Now().UnixNano()),
	)
}
