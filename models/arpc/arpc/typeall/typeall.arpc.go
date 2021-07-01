package typeall

// TypeAll carries all grpc and Colfer supported data types
type TypesAll struct {
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

// Request carries a list of typesAll
type Request struct {
	typesAll []TypesAll
}

// EmptyResult an empty result
type EmptyResult struct {}

// TypeAll receives a request containing all data types and retuning an empty response
type TypeAll func(request *Request) (*EmptyResult, error)
