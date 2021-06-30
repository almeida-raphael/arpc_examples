package typeuint32

// Request carries one data type
type Request struct {
	entries uint32
}

// EmptyResult an empty result
type EmptyResult struct {}

// TypeUInt32 receives a request containing a data type and retuning an empty response
type TypeUInt32 func(request *Request) (*EmptyResult, error)
