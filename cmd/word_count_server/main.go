package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/almeida-raphael/arpc/channel"
	"github.com/almeida-raphael/arpc/controller"
	"github.com/almeida-raphael/arpc_examples/examples/wordcount"
	"github.com/almeida-raphael/arpc_examples/utils"
	"os"
	"strings"
	"time"
)


// WordCountServerDefinition struct to implement wordcount aRPC service procedures
type WordCountServerDefinition struct {}

// CountWords aRPC WordCount.CountWords function implementation
func (gs *WordCountServerDefinition)CountWords(request *wordcount.Text)(*wordcount.CountedWords, error){
	processingTime := time.Now()
	wordFrequency := make(map[string]uint64)
	for _, word := range strings.Fields(request.Data){
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
	fmt.Printf("Processing Time: %vs", time.Since(processingTime).Seconds())

	return &response, nil
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

	wordcount.RegisterWordcountServer(aRPCController, &WordCountServerDefinition{})

	err = aRPCController.StartServer(context.Background())
	if err != nil {
		panic(err)
	}
}

