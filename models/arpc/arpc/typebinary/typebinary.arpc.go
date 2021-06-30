package typebinary

// Request carries one data type
type Request struct {
	entries binary
}

// EmptyResult an empty result
type EmptyResult struct {}

// TypeBinary receives a request containing a data type and retuning an empty response
type TypeBinary func(request *Request) (*EmptyResult, error)
