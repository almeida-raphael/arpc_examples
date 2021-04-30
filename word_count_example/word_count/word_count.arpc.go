package word_count

import (
	"context"
	"github.com/almeida-raphael/arpc/controller"
	"github.com/almeida-raphael/arpc/helpers"
)

var serviceID = helpers.Hash("word_count_client")

//////////////////////////////////////////////////////// CLIENT ////////////////////////////////////////////////////////

type WordCount struct {
	controller *controller.RPC
}

func NewWordCount(controller *controller.RPC) WordCount {
	return WordCount{
		controller: controller,
	}
}

func (wordCount *WordCount)CountWords(request *Text, ctx ...context.Context)(*CountedWords, error){
	if ctx == nil || len(ctx) == 0{
		ctx = []context.Context{context.Background()}
	}
	response := CountedWords{}
	err := wordCount.controller.SendRPCCall(ctx[0], serviceID, 0, request, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

//////////////////////////////////////////////////////// SERVER ////////////////////////////////////////////////////////

type WordCountServer interface {
	CountWords(request *Text)(*CountedWords, error)
}

func bindCountWords(server WordCountServer)(
	func(msg []byte)([]byte, error),
) {
	return func(msg []byte)([]byte, error){
		request := Text{}
		err := request.UnmarshalBinary(msg)
		if err != nil {
			return nil, err
		}

		response, err := server.CountWords(&request)
		if err != nil {
			return nil, err
		}

		responseBytes, err := response.MarshalBinary()
		if err != nil {
			return nil, err
		}

		return responseBytes, nil
	}
}

////////////////////////////////////////////////////// REGISTRAR ///////////////////////////////////////////////////////

func RegisterWordCountServer(controller controller.RPC, srv WordCountServer){
	controller.RegisterService(
		serviceID,
		map[uint16]func(message []byte)([]byte, error){
			0: bindCountWords(srv),
		},
	)
}
