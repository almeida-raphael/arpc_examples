package helloreply

// Text a text message
type Text struct {
	data  text
}

// SayHello sends a message and waits for a return message that contains ”Hello ” prefixed on the message
type SayHello func(request *Text)(*Text, error)
