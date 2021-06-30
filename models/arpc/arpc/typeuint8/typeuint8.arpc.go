package typeuint8

// Request carries one data type
type Request struct {
	entries uint8
}

// EmptyResult an empty result
type EmptyResult struct {}

// TypeUInt8 receives a request containing a data type and retuning an empty response
type TypeUInt8 func(request *Request) (*EmptyResult, error)
