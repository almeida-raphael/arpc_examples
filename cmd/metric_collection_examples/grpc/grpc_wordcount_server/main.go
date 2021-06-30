package main

import (
	"context"
	"os"
	"strings"

	wordcount2 "github.com/almeida-raphael/arpc_examples/models/grpc/wordcount"

	"github.com/almeida-raphael/arpc_examples/utils"
	"google.golang.org/protobuf/proto"
)

// ServerDefinition struct to implement greetings aRPC service procedures
type ServerDefinition struct {
	wordcount2.UnimplementedWordCountServer
}

// countWords aRPC server function implementation
func countWords(request proto.Message) (proto.Message, error) {
	reqData := request.(*wordcount2.Text)

	wordFrequency := make(map[string]uint64)
	for _, word := range strings.Fields(reqData.Data) {
		wordFrequency[word]++
	}

	wordFrequencyList := make([]*wordcount2.CountedWords_Entry, 0, len(wordFrequency))
	for word, count := range wordFrequency {
		wordFrequencyList = append(wordFrequencyList, &wordcount2.CountedWords_Entry{
			Word:  word,
			Count: count,
		})
	}

	response := wordcount2.CountedWords{
		Entries: wordFrequencyList,
	}

	return &response, nil
}

var metricsCountWords = utils.CollectGRPCServerMetrics(
	utils.Atoi(os.Getenv("SAMPLE_THREADS")), utils.Atoi(os.Getenv("TRIALS")), countWords,
	"results/gRPC/wordcount/server/%d.json",
)

// CountWords gRPC server function implementation
func (gs *ServerDefinition) CountWords(ctx context.Context, request *wordcount2.Text) (*wordcount2.CountedWords, error) {
	respData := &wordcount2.CountedWords{}

	response, err := metricsCountWords(request)
	if err == nil {
		respData = response.(*wordcount2.CountedWords)
	}

	return respData, err
}

func main() {
	listener, gRPCServer := utils.SetupGRPCServer()
	wordcount2.RegisterWordCountServer(gRPCServer, &ServerDefinition{})
	utils.StartGRPCServer(listener, gRPCServer)
}
