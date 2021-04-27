package greetings

import (
	"github.com/almeida-raphael/arpc/controller"
	"github.com/almeida-raphael/arpc/headers"
	"github.com/almeida-raphael/arpc/helpers"
)

var serviceID = helpers.Hash("greetings")

//////////////////////////////////////////////////////// CLIENT ////////////////////////////////////////////////////////

type Greetings struct {
	controller controller.RPC
}

func NewGreetings(controller controller.RPC) Greetings {
	return Greetings{
		controller: controller,
	}
}

func (greetings *Greetings)SayHi(request *SayHiRequest)(*SayHiResponse, error){
	response := SayHiResponse{}
	err := greetings.controller.SendRPC(headers.Call, serviceID, 0, request, &response)
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
	bind := func(msg []byte)([]byte, error){
		request := SayHiRequest{}
		err := request.UnmarshalBinary(msg)
		if err != nil {
			return nil, err
		}

		response, err := server.SayHi(&request)

		responseBytes, err := helpers.SerializeWithHeaders(headers.Result, serviceID, 0, response)
		if err != nil {
			return nil, err
		}

		return responseBytes, nil
	}

	return bind
}

////////////////////////////////////////////////////// REGISTRAR ///////////////////////////////////////////////////////

func RegisterGreetingsServer(controller controller.RPC, srv GreetingsServer)error{
	err := controller.RegisterService(
		serviceID,
		map[uint16]func(message []byte)([]byte, error){
			0: bindSayHi(srv),
		},
	)
	if err != nil {
		return err
	}

	return nil
}
