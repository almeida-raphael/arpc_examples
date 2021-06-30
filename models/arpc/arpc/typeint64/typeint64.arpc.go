package typeint64

// Request carries one data type
type Request struct {
	entries int64
}

// EmptyResult an empty result
type EmptyResult struct {}

// TypeInt64 receives a request containing a data type and retuning an empty response
type TypeInt64 func(request *Request) (*EmptyResult, error)
