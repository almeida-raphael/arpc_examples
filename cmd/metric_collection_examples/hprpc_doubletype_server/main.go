package main

import (
	"github.com/almeida-raphael/arpc/interfaces"
	"github.com/almeida-raphael/arpc_examples/hprpc/doubletype"
	"github.com/almeida-raphael/arpc_examples/utils"
)

// ServerDefinition struct to implement greetings aRPC service procedures
type ServerDefinition struct {}

// getNumbers aRPC server function implementation
func average(request interfaces.Serializable)(interfaces.Serializable, error){
	reqData := request.(*doubletype.NumberList)
	return &doubletype.Result{Value: utils.Average(reqData.Entries)}, nil
}

var metricsAverage = utils.CollectServerMetrics(
	20, 1000, average,
	"results/doubletype/server/%d.json",
)

// Average aRPC server function implementation
func (gs *ServerDefinition)Average(request *doubletype.NumberList)(*doubletype.Result, error){
	var reqData interfaces.Serializable = request
	respData := &doubletype.Result{}

	response, err := metricsAverage(reqData)
	if err == nil {
		respData = response.(*doubletype.Result)
	}

	return respData, err
}

func main(){
	aRPCController := utils.SetupServer()
	doubletype.RegisterDoubletypeServer(aRPCController, &ServerDefinition{})
	utils.StartServer(aRPCController)
}

