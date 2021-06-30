package typetext

// Request carries one data type
type Request struct {
	entries text
}

// EmptyResult an empty result
type EmptyResult struct {}

// TypeText receives a request containing a data type and retuning an empty response
type TypeText func(request *Request) (*EmptyResult, error)
