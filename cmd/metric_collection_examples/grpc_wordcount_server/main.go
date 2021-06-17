package main

import (
	"context"
	"strings"

	"github.com/almeida-raphael/arpc_examples/grpc/wordcount"

	"github.com/almeida-raphael/arpc_examples/utils"
	"google.golang.org/protobuf/proto"
)

// ServerDefinition struct to implement greetings aRPC service procedures
type ServerDefinition struct {
	wordcount.UnimplementedWordCountServer
}

// countWords aRPC server function implementation
func countWords(request proto.Message) (proto.Message, error) {
	reqData := request.(*wordcount.Text)

	wordFrequency := make(map[string]uint64)
	for _, word := range strings.Fields(reqData.Data) {
		wordFrequency[word]++
	}

	wordFrequencyList := make([]*wordcount.CountedWords_Entry, 0, len(wordFrequency))
	for word, count := range wordFrequency {
		wordFrequencyList = append(wordFrequencyList, &wordcount.CountedWords_Entry{
			Word:  word,
			Count: count,
		})
	}

	response := wordcount.CountedWords{
		Entries: wordFrequencyList,
	}

	return &response, nil
}

var metricsCountWords = utils.CollectGRPCServerMetrics(
	20, 1000, countWords,
	"results/wordcount_grpc/server/%d.json",
)

// CountWords gRPC server function implementation
func (gs *ServerDefinition) CountWords(ctx context.Context, request *wordcount.Text) (*wordcount.CountedWords, error) {
	respData := &wordcount.CountedWords{}

	response, err := metricsCountWords(request)
	if err == nil {
		respData = response.(*wordcount.CountedWords)
	}

	return respData, err
}

func main() {
	listener, gRPCServer := utils.SetupGRPCServer()
	wordcount.RegisterWordCountServer(gRPCServer, &ServerDefinition{})
	utils.StartGRPCServer(listener, gRPCServer)
}
