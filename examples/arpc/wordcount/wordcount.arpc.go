package wordcount

// Text input aRPC procedure text
type Text struct {
	data  text
}

// Entry frequency storage struct for a given word
type Entry struct {
	word text
	count uint64
}

// CountedWords aRPC word frequency response
type CountedWords struct {
	message []Entry
}

// CountWords aRPC procedure declaration
type CountWords func(request *Text)(*CountedWords, error)
