package greetings

import (
	"context"
	"github.com/almeida-raphael/arpc/controller"
	"github.com/almeida-raphael/arpc/helpers"
)

var serviceID = helpers.Hash("greetings_client")

//////////////////////////////////////////////////////// CLIENT ////////////////////////////////////////////////////////

type Greetings struct {
	controller *controller.RPC
}

func NewGreetings(controller *controller.RPC) Greetings {
	return Greetings{
		controller: controller,
	}
}

func (greetings *Greetings)SayHi(request *SayHiRequest, ctx ...context.Context)(*SayHiResponse, error){
	if ctx == nil || len(ctx) == 0{
		ctx = []context.Context{context.Background()}
	}
	response := SayHiResponse{}
	err := greetings.controller.SendRPCCall(ctx[0], serviceID, 0, request, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

//////////////////////////////////////////////////////// SERVER ////////////////////////////////////////////////////////

type GreetingsServer interface {
	SayHi(request *SayHiRequest)(*SayHiResponse, error)
}

func bindSayHi(server GreetingsServer)(
	func(msg []byte)([]byte, error),
) {
	return func(msg []byte)([]byte, error){
		request := SayHiRequest{}
		err := request.UnmarshalBinary(msg)
		if err != nil {
			return nil, err
		}

		response, err := server.SayHi(&request)
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

func RegisterGreetingsServer(controller controller.RPC, srv GreetingsServer){
	controller.RegisterService(
		serviceID,
		map[uint16]func(message []byte)([]byte, error){
			0: bindSayHi(srv),
		},
	)
}
