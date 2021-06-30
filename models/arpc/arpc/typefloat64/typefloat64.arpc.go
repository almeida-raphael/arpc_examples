package typefloat64

// Request carries one data type
type Request struct {
	entries float64
}

// EmptyResult an empty result
type EmptyResult struct {}

// TypeFloat64 receives a request containing a data type and retuning an empty response
type TypeFloat64 func(request *Request) (*EmptyResult, error)
