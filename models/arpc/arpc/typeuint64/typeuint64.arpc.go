package typeuint64

// Request carries one data type
type Request struct {
	entries uint64
}

// EmptyResult an empty result
type EmptyResult struct {}

// TypeUInt64 receives a request containing a data type and retuning an empty response
type TypeUInt64 func(request *Request) (*EmptyResult, error)
