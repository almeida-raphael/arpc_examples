package typeuint16

// Request carries one data type
type Request struct {
	entries uint16
}

// EmptyResult an empty result
type EmptyResult struct {}

// TypeUInt16 receives a request containing a data type and retuning an empty response
type TypeUInt16 func(request *Request) (*EmptyResult, error)
