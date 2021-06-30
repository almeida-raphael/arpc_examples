package echo

// Numbers a numbers list
type Numbers struct {
	entries []int32
}

// Yell sends a message and waits for a return message that contains the same message
type Yell func(request *Numbers) (*Numbers, error)
