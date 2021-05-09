package greetings

// Person info struct
type Person struct {
	title uint64
	name  text
}

// SayHiRequest aRPC input struct
type SayHiRequest struct {
	person Person
}

// SayHiResponse aRPC output struct
type SayHiResponse struct {
	response text
}

// SayHi aRPC procedure declaration
type SayHi func(request *SayHiRequest) (*SayHiResponse, error)
