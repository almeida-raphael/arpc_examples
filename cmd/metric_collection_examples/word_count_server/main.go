package main

import (
	"github.com/almeida-raphael/arpc/interfaces"
	"github.com/almeida-raphael/arpc_examples/examples/wordcount"
	"github.com/almeida-raphael/arpc_examples/utils"
	"strings"
)


// WordCountServerDefinition struct to implement wordcount aRPC service procedures
type WordCountServerDefinition struct {}

// countWords aRPC server function implementation
func countWords(request interfaces.Serializable)(interfaces.Serializable, error){
	reqData := request.(*wordcount.Text)

	wordFrequency := make(map[string]uint64)
	for _, word := range strings.Fields(reqData.Data){
		wordFrequency[word] ++
	}

	wordFrequencyList := make([]*wordcount.Entry, 0, len(wordFrequency))
	for word, count := range wordFrequency{
		wordFrequencyList = append(wordFrequencyList, &wordcount.Entry{
			Word:  word,
			Count: count,
		})
	}

	response := wordcount.CountedWords{
		Message: wordFrequencyList,
	}

	return &response, nil
}

var metricsCountWords = utils.CollectServerMetrics(
	20, 1000, countWords,
	"results/wordcount/server/%d.json",
)

// CountWords aRPC WordCount.CountWords function implementation
func (gs *WordCountServerDefinition)CountWords(request *wordcount.Text)(*wordcount.CountedWords, error){
	var reqData interfaces.Serializable = request
	respData := &wordcount.CountedWords{}

	response, err := metricsCountWords(reqData)
	if err == nil {
		respData = response.(*wordcount.CountedWords)
	}

	return respData, err
}

func main(){
	aRPCController := utils.SetupServer()
	wordcount.RegisterWordcountServer(aRPCController, &WordCountServerDefinition{})
	utils.StartServer(aRPCController)
}

