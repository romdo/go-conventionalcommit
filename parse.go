package conventionalcommit

// Parse parses a conventional commit message and returns it as a *Message
// struct.
func Parse(message []byte) (*Message, error) {
	buffer := NewBuffer(message)

	return NewMessage(buffer)
}
