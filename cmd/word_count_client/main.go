package main

import (
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/almeida-raphael/arpc/channel"
	"github.com/almeida-raphael/arpc/controller"
	arpcerrors "github.com/almeida-raphael/arpc/errors"
	"github.com/almeida-raphael/arpc_examples/examples/word_count"
	"github.com/almeida-raphael/arpc_examples/utils"
	"io/ioutil"
	"os"
	"time"
)

func main(){
	rootCAPath := os.Getenv("CA_FILE")
	serverAddress := os.Getenv("SERVER_ADDRESS")
	assetPath := os.Getenv("ASSET_PATH")

	rootCA, err := utils.LoadCA(rootCAPath)
	if err != nil {
		panic(err)
	}

	tlsConfig := tls.Config{
		RootCAs:                     rootCA,
		NextProtos:                  []string{"quic-arcp"},
	}

	aRPCController := controller.NewRPCController(channel.NewQUICChannel(
		serverAddress,
		7653,
		&tlsConfig,
		nil,
	))
	wordCountService := word_count.NewWordCount(&aRPCController)

	err = aRPCController.StartClient()
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadFile(assetPath)
	if err != nil {
		panic(err)
	}

	requestData := word_count.Text{
		Data: string(data),
	}

	rpcStartTime := time.Now()
	wordCountResponse, err := wordCountService.CountWords(&requestData)
	rpcElapsedTime := time.Since(rpcStartTime)

	if err != nil {
		if errors.Is(err, &arpcerrors.Remote{}) {
			fmt.Printf("Remote Error: %v", err)
		}else{
			panic(err)
		}
	}else{
		fmt.Printf("Response:\n")
		for _, entry := range wordCountResponse.Message{
			fmt.Printf("%s: %d\n", entry.Word, entry.Count)
		}
	}

	requestSerializationTime, requestDeserializationTime, requestSize, err := utils.TestSerialization(&requestData)
	if err != nil {
		panic(err)
	}
	responseSerializationTime, responseDeserializationTime, responseSize, err := utils.TestSerialization(
		wordCountResponse,
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf(`
		Metrics:
		Request Serialization Time: %vs
		Request Deserialization Time: %vs
		Serialized Request Size: %db

		Response Serialization Time: %vs
		Response Deserialization Time: %vs
		Deerialized Response Size: %db

		aRPC Transport Time: %vs
		aRPC Total Time: %vs
		`,
		(*requestSerializationTime).Seconds(),
		(*requestDeserializationTime).Seconds(),
		requestSize,

		(*responseSerializationTime).Seconds(),
		(*responseDeserializationTime).Seconds(),
		responseSize,

		(rpcElapsedTime - *requestSerializationTime - *requestDeserializationTime -
			*responseSerializationTime - *responseDeserializationTime).Seconds(),
		rpcElapsedTime.Seconds(),
	)
}

