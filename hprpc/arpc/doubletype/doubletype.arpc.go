package doubletype

// NumberList a list of numbers to be averaged
type NumberList struct {
	entries  []int32
}

// Result the average result for the number list
type Result struct {
	value  int32
}

// Average takes the average of a list of 32 bit numbers
type Average func(request *NumberList)(*Result, error)
