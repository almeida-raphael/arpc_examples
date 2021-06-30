package typeint32

// Request carries one data type
type Request struct {
	entries int32
}

// EmptyResult an empty result
type EmptyResult struct {}

// TypeInt32 receives a request containing a data type and retuning an empty response
type TypeInt32 func(request *Request) (*EmptyResult, error)
