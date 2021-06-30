package typeall

// Request carries all data types
type Request struct {
	binary    binary
	bool      bool
	float32   float32
	float64   float64
	int32     int32
	int64     int64
	text      text
	uint32    uint32
	uint64    uint64
}

// EmptyResult an empty result
type EmptyResult struct {}

// TypeAll receives a request containing all data types and retuning an empty response
type TypeAll func(request *Request) (*EmptyResult, error)
