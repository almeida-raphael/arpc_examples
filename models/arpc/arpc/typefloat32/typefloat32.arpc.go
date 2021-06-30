package typefloat32

// Request carries one data type
type Request struct {
	entries float32
}

// EmptyResult an empty result
type EmptyResult struct {}

// TypeFloat32 receives a request containing a data type and retuning an empty response
type TypeFloat32 func(request *Request) (*EmptyResult, error)
