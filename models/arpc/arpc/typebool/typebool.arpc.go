package typebool

// Request carries one data type
type Request struct {
	entries bool
}

// EmptyResult an empty result
type EmptyResult struct {}

// TypeBool receives a request containing a data type and retuning an empty response
type TypeBool func(request *Request) (*EmptyResult, error)
