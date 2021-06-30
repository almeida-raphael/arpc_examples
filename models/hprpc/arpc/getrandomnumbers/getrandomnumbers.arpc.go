package getrandomnumbers

// Amount the amount of numbers to be generated
type Amount struct {
	value  int32
}

// NumberList a list of the generated numbers
type NumberList struct {
	entries  []int32
}

// GetNumbers Generates a list of random numbers
type GetNumbers func(amount *Amount)(*NumberList, error)
