package conventionalcommit

// RawMessage represents a commit message in a more structured form than a
// simple string or byte slice. This makes it easier to process a message for
// the purposes of extracting detailed information, linting, and formatting.
type RawMessage struct {
	// Lines is a list of all individual lines of text in the commit message,
	// which also includes the original line number, making it easy to pass a
	// single Line around while still knowing where in the original commit
	// message it belongs.
	Lines Lines

	// Paragraphs is a list of textual paragraphs in the commit message. A
	// paragraph is defined as any continuous sequence of lines which are not
	// empty or consist of only whitespace.
	Paragraphs []*Paragraph
}

// NewRawMessage returns a RawMessage, with the given commit message broken down
// into individual lines of text, with sequential non-empty lines grouped into
// paragraphs.
func NewRawMessage(message []byte) *RawMessage {
	r := &RawMessage{
		Lines:      Lines{},
		Paragraphs: []*Paragraph{},
	}

	if len(message) == 0 {
		return r
	}

	r.Lines = NewLines(message)
	r.Paragraphs = NewParagraphs(r.Lines)

	return r
}

// Bytes renders the RawMessage back into a byte slice which is identical to the
// original input byte slice given to NewRawMessage. This includes retaining the
// original line break types for each line.
func (s *RawMessage) Bytes() []byte {
	return s.Lines.Bytes()
}

// String renders the RawMessage back into a string which is identical to the
// original input byte slice given to NewRawMessage. This includes retaining the
// original line break types for each line.
func (s *RawMessage) String() string {
	return s.Lines.String()
}
