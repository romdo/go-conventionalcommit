package conventionalcommit

const (
	lf = 10 // linefeed ("\n") character
	cr = 13 // carriage return ("\r") character
)

// Line represents a single line of text defined as; A continuous sequence of
// bytes which do not contain a \r (carriage return) or \n (line-feed) byte.
type Line struct {
	// Line number within commit message, starting a 1 rather than 0, as
	// text viewed in a text editor starts on line 1, not line 0.
	Number int

	// Content is the raw bytes that make up the text content in the line.
	Content []byte

	// Break is the linebreak type used at the end of the line. It will be one
	// of "\n", "\r\n", "\r", or empty if it is the very last line.
	Break []byte
}

// Lines is a slice of *Line types with some helper methods attached.
type Lines []*Line

// NewLines breaks the given byte slice down into a slice of Line structs,
// allowing easier inspection and manipulation of content on a line-by-line
// basis.
func NewLines(content []byte) Lines {
	r := Lines{}

	if len(content) == 0 {
		return r
	}

	// List of start/end offsets for each line break.
	var breaks [][]int

	// Locate each line break within content.
	for i := 0; i < len(content); i++ {
		if content[i] == lf {
			breaks = append(breaks, []int{i, i + 1})
		} else if content[i] == cr {
			b := []int{i, i + 1}
			if i+1 < len(content) && content[i+1] == lf {
				b[1]++
				i++
			}
			breaks = append(breaks, b)
		}
	}

	// Return a single line if there are no line breaks.
	if len(breaks) == 0 {
		return Lines{{Number: 1, Content: content, Break: []byte{}}}
	}

	// Extract each line based on linebreak offsets.
	offset := 0
	for n, loc := range breaks {
		r = append(r, &Line{
			Number:  n + 1,
			Content: content[offset:loc[0]],
			Break:   content[loc[0]:loc[1]],
		})
		offset = loc[1]
	}

	// Extract final line
	r = append(r, &Line{
		Number:  len(breaks) + 1,
		Content: content[offset:],
		Break:   []byte{},
	})

	return r
}

// Bytes combines all Lines into a single byte slice, retaining the original
// line break types for each line.
func (s Lines) Bytes() []byte {
	// Pre-calculate capacity of result byte slice.
	size := 0
	for _, l := range s {
		size = size + len(l.Content) + len(l.Break)
	}

	b := make([]byte, 0, size)

	for _, l := range s {
		b = append(b, l.Content...)
		b = append(b, l.Break...)
	}

	return b
}

// Bytes combines all Lines into a single string, retaining the original line
// break types for each line.
func (s Lines) String() string {
	return string(s.Bytes())
}
