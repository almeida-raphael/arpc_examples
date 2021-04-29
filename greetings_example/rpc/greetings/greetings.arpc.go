package greetings

type Person struct{}
type SayHiRequest struct{}
type SayHiResponse struct{}

var _ = `
	type Person struct {
		title uint64
		name  text
	}
	
	type SayHiRequest struct {
		person Person
	}
	
	type SayHiResponse struct {
		response text
	}
`

type SayHi func(request *SayHiRequest)(*SayHiResponse, error)
