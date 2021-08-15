package conventionalcommit

import (
	"regexp"
)

// footerToken will match against all variations of Conventional Commit footer
// formats.
//
// Examples of valid footer tokens:
//
//  Approved-by: John Carter
//  ReviewdBy: Noctis
//  Fixes #49
//  Reverts #SOL-42
//  BREAKING CHANGE: Flux capacitor no longer exists.
//  BREAKING-CHANGE: Time will flow backwads
//
// Examples of invalid footer tokens:
//
//  Approved-by:
//  Approved-by:John Carter
//  Approved by: John Carter
//    ReviewdBy: Noctis
//  Fixes#49
//  Fixes #
//  Fixes 49
//  BREAKING CHANGE:Flux capacitor no longer exists.
//  Breaking Change: Flux capacitor no longer exists.
//  Breaking-Change: Time will flow backwads
//
var footerToken = regexp.MustCompile(
	`^(?:([\w-]+)\s+(#.+)|([\w-]+|BREAKING[\s-]CHANGE):\s+(.+))$`,
)

// Buffer represents a commit message in a more structured form than a simple
// string or byte slice. This makes it easier to process a message for the
// purposes of extracting detailed information, linting, and formatting.
//
// The commit message is conceptually broken down into two three separate
// sections:
//
// - Head section holds the commit message subject/description, along with type
//   and scope for conventional commits. The head section should only ever be a
//   single line according to git convention, but Buffer supports multi-line
//   headers so they can be parsed and handled as needed.
//
// - Body section holds the rest of the message. Except if the last paragraph
//   starts with a footer token, then the last paragraph is omitted from the
//   body section.
//
// - Foot section holds conventional commit footers. It is always the last
//   paragraph of a commit message, and is only considered to be the foot
//   section if the first line of the paragraph beings with a footer token.
//
// Each section is returned as a Lines type, which provides per-line access to
// the text within the section.
type Buffer struct {
	// firstLine is the lines offset for the first line which contains any
	// non-whitespace character.
	firstLine int

	// lastLine is the lines offset for the last line which contains any
	// non-whitespace character.
	lastLine int

	// headLen is the number of lines that the headLen section (first paragraph)
	// spans.
	headLen int

	// footLen is the number of lines that the footLen section (last paragraph)
	// spans.
	footLen int

	// lines is a list of all individual lines of text in the commit message,
	// which also includes the original line number, making it easy to pass a
	// single Line around while still knowing where in the original commit
	// message it belongs.
	lines Lines
}

// NewBuffer returns a Buffer, with the given commit message broken down into
// individual lines of text, with sequential non-empty lines grouped into
// paragraphs.
func NewBuffer(message []byte) *Buffer {
	buf := &Buffer{
		lines: Lines{},
	}

	if len(message) == 0 {
		return buf
	}

	buf.lines = NewLines(message)
	// Find fist non-whitespace line.
	if i := buf.lines.FirstTextIndex(); i > -1 {
		buf.firstLine = i
	}

	// Find last non-whitespace line.
	if i := buf.lines.LastTextIndex(); i > -1 {
		buf.lastLine = i
	}

	// Determine number of lines in first paragraph (head section).
	for i := buf.firstLine; i <= buf.lastLine; i++ {
		if buf.lines[i].Blank() {
			break
		}
		buf.headLen++
	}

	// Determine number of lines in the last paragraph.
	lastLen := 0
	for i := buf.lastLine; i > buf.firstLine+buf.headLen; i-- {
		if buf.lines[i].Blank() {
			break
		}
		lastLen++
	}

	// If last paragraph starts with a Convention Commit footer token, it is the
	// foot section, otherwise it is part of the body.
	if lastLen > 0 {
		line := buf.lines[buf.lastLine-lastLen+1]
		if footerToken.Match(line.Content) {
			buf.footLen = lastLen
		}
	}

	return buf
}

// Head returns the first paragraph, defined as the first group of sequential
// lines which contain any non-whitespace characters.
func (s *Buffer) Head() Lines {
	return s.lines[s.firstLine : s.firstLine+s.headLen]
}

// Body returns all lines between the first and last paragraphs. If the body is
// surrounded by multiple empty lines, they will be removed, ensuring first and
// last line of body is not a blank whitespace line.
func (s *Buffer) Body() Lines {
	if s.firstLine == s.lastLine {
		return Lines{}
	}

	first := s.firstLine + s.headLen + 1
	last := s.lastLine + 1

	if s.footLen > 0 {
		last -= s.footLen
	}

	return s.lines[first:last].Trim()
}

// Head returns the last paragraph, defined as the last group of sequential
// lines which contain any non-whitespace characters.
func (s *Buffer) Foot() Lines {
	if s.footLen == 0 {
		return Lines{}
	}

	return s.lines[s.lastLine-s.footLen+1 : s.lastLine+1]
}

// Lines returns all lines with any blank lines from the beginning and end of
// the buffer removed. Effectively all lines from the first to the last line
// which contain any non-whitespace characters.
func (s *Buffer) Lines() Lines {
	if s.lastLine+1 > len(s.lines) || (s.lastLine == 0 && s.lines[0].Blank()) {
		return Lines{}
	}

	return s.lines[s.firstLine : s.lastLine+1]
}

func (s *Buffer) LineCount() int {
	if s.headLen == 0 {
		return 0
	}

	return (s.lastLine + 1) - s.firstLine
}

// Bytes renders the Buffer back into a byte slice, without any leading or
// trailing whitespace lines. Leading whitespace on the first line which
// contains non-whitespace characters is retained. It is only whole lines
// consisting of only whitespace which are excluded.
func (s *Buffer) Bytes() []byte {
	return s.Lines().Bytes()
}

// String renders the Buffer back into a string, without any leading or trailing
// whitespace lines. Leading whitespace on the first line which contains
// non-whitespace characters is retained. It is only whole lines consisting of
// only whitespace which are excluded.
func (s *Buffer) String() string {
	return s.Lines().String()
}

// BytesRaw renders the Buffer back into a byte slice which is identical to the
// original input byte slice given to NewBuffer. This includes retaining the
// original line break types for each line.
func (s *Buffer) BytesRaw() []byte {
	return s.lines.Bytes()
}

// StringRaw renders the Buffer back into a string which is identical to the
// original input byte slice given to NewBuffer. This includes retaining the
// original line break types for each line.
func (s *Buffer) StringRaw() string {
	return s.lines.String()
}
