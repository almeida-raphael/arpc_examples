package main

import (
	"fmt"
	"github.com/almeida-raphael/arpc/interfaces"
	"github.com/almeida-raphael/arpc_examples/hprpc/getrandomnumbers"
	"github.com/almeida-raphael/arpc_examples/utils"
	"time"
)


func main(){
	aRPCController := utils.SetupClient()

	service := getrandomnumbers.NewGetrandomnumbers(&aRPCController)
	err := aRPCController.StartClient()
	if err != nil {
		panic(err)
	}

	requestData := getrandomnumbers.Amount{Value: 1000}

	GetNumbers := func(req interfaces.Serializable)(interfaces.Serializable, error){
		reqData := req.(*getrandomnumbers.Amount)
		response, err := service.GetNumbers(reqData)
		if err != nil {
			return nil, err
		}

		return response, nil
	}

	err = utils.RunClientRPCAndCollectMetrics(
		20, 1000, &requestData, GetNumbers,
		fmt.Sprintf("results/getrandomnumbers/client/%d.json", time.Now().UnixNano()),
	)
	if err != nil {
		panic(err)
	}
}

