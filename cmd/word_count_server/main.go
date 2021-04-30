package main

import (
	"context"
	"crypto/tls"
	"github.com/almeida-raphael/arpc/channel"
	"github.com/almeida-raphael/arpc/controller"
	"github.com/almeida-raphael/arpc_examples/utils"
	"github.com/almeida-raphael/arpc_examples/word_count_example/word_count"
	"os"
	"strings"
)


type WordCountServerDefinition struct {}

func (gs *WordCountServerDefinition)CountWords(request *word_count.Text)(*word_count.CountedWords, error){
	wordFrequency := make(map[string]uint64)
	for _, word := range strings.Fields(request.Data){
		if _, ok := wordFrequency[word]; ok{
			wordFrequency[word] += 1
		}else{
			wordFrequency[word] = 1
		}
	}

	var wordFrequencyList []*word_count.Entry
	for word, count := range wordFrequency{
		wordFrequencyList = append(wordFrequencyList, &word_count.Entry{
			Word:  word,
			Count: count,
		})
	}

	return &word_count.CountedWords{
		Message: wordFrequencyList,
	}, nil
}

func main(){
	certFilePath := os.Getenv("CERT_FILE")
	keyFilePath := os.Getenv("KEY_FILE")
	serverAddress := os.Getenv("SERVER_ADDRESS")

	certificates, err := utils.LoadCertificates(certFilePath, keyFilePath)
	if err != nil {
		panic(err)
	}

	tlsConfig := tls.Config{
		Certificates:                []tls.Certificate{*certificates},
		NextProtos:                  []string{"quic-arcp"},
	}

	aRPCController := controller.NewRPCController(channel.NewQUICChannel(
		serverAddress,
		7653,
		&tlsConfig,
		nil,
	))

	word_count.RegisterWordCountServer(aRPCController, &WordCountServerDefinition{})

	err = aRPCController.StartServer(context.Background())
	if err != nil {
		panic(err)
	}
}

