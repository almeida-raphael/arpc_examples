package typetimestamp

// Request carries one data type
type Request struct {
	entries timestamp
}

// EmptyResult an empty result
type EmptyResult struct {}

// TypeTimestamp receives a request containing a data type and retuning an empty response
type TypeTimestamp func(request *Request) (*EmptyResult, error)
