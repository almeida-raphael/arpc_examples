package main

import (
	"fmt"
	"github.com/almeida-raphael/arpc/interfaces"
	"github.com/almeida-raphael/arpc_examples/hprpc/helloreply"
	"github.com/almeida-raphael/arpc_examples/utils"
	"time"
)


func main(){
	aRPCController := utils.SetupClient()

	service := helloreply.NewHelloreply(&aRPCController)
	err := aRPCController.StartClient()
	if err != nil {
		panic(err)
	}

	SayHello := func(req interfaces.Serializable)(interfaces.Serializable, error){
		reqData := req.(*helloreply.Text)
		response, err := service.SayHello(reqData)
		if err != nil {
			return nil, err
		}

		return response, nil
	}

	err = utils.RunClientRPCAndCollectMetrics(
		20, 1000, &helloreply.Text{Data: utils.GenerateString(1000)}, SayHello,
		fmt.Sprintf("results/helloreply/client/%d.json", time.Now().UnixNano()),
	)
	if err != nil {
		panic(err)
	}
}

