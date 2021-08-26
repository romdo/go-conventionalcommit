package conventionalcommit

import (
	"bytes"
	"strings"
)

const (
	lf   = 10 // ASCII linefeed ("\n") character.
	cr   = 13 // ASCII carriage return ("\r") character.
	hash = 35 // ASCII hash ("#") character.

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

// Empty returns true if line content has a length of zero.
func (s *Line) Empty() bool {
	return len(s.Content) == 0
}

// Blank returns true if line content has a length of zero after leading and
// trailing white space has been trimmed.
func (s *Line) Blank() bool {
	return len(bytes.TrimSpace(s.Content)) == 0
}

// Comment returns true if line content is a commit comment, where the first
// non-whitespace character in the line is a hash (#).
func (s *Line) Comment() bool {
	trimmed := bytes.TrimSpace(s.Content)

	if len(trimmed) == 0 {
		return false
	}

	return trimmed[0] == hash
}

// Lines is a slice of *Line types with some helper methods attached.
type Lines []*Line

// NewLines breaks the given byte slice down into a slice of Line structs,
// allowing easier inspection and manipulation of content on a line-by-line
// basis.
func NewLines(content []byte) Lines {
	r := Lines{}
	cLen := len(content)

	if cLen == 0 {
		return r
	}

	// List of start/end offsets for each line break.
	var breaks [][]int

	// Locate each line break within content.
	for i := 0; i < cLen; i++ {
		switch content[i] {
		case lf:
			breaks = append(breaks, []int{i, i + 1})
		case cr:
			b := []int{i, i + 1}
			if i+1 < cLen && content[i+1] == lf {
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

// FirstTextIndex returns the line offset of the first line which contains any
// non-whitespace characters.
func (s Lines) FirstTextIndex() int {
	for i, line := range s {
		if !line.Blank() {
			return i
		}
	}

	return -1
}

// LastTextIndex returns the line offset of the last line which contains any
// non-whitespace characters.
func (s Lines) LastTextIndex() int {
	for i := len(s) - 1; i >= 0; i-- {
		if !s[i].Blank() {
			return i
		}
	}

	return -1
}

// Trim returns a new Lines instance where all leading and trailing whitespace
// lines have been removed, based on index values from FirstTextIndex() and
// LastTextIndex().
//
// If there are no lines with non-whitespace characters, a empty Lines type is
// returned.
func (s Lines) Trim() Lines {
	start := s.FirstTextIndex()
	if start == -1 {
		return Lines{}
	}

	return s[start : s.LastTextIndex()+1]
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

func (s Lines) Join(sep string) string {
	r := make([]string, 0, len(s))
	for _, line := range s {
		r = append(r, string(line.Content))
	}

	return strings.Join(r, sep)
}
