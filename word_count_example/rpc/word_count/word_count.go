package word_count

type Text struct{}
type CountedWords struct{}
type SayHiResponse struct{}

var _ = `
	type Text struct {
		data  text
	}
	
	type Entry struct {
		word text
		count uint64
	}
	
	type CountedWords struct {
		person []Entry
	}
`

type CountWords func(request *Text)(*CountedWords, error)
