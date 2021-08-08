package conventionalcommit

import "bytes"

// Paragraph represents a textual paragraph defined as; A continuous sequence of
// textual lines which are not empty or and do not consist of only whitespace.
type Paragraph struct {
	// Lines is a list of lines which collectively form a paragraph.
	Lines Lines
}

func NewParagraphs(lines Lines) []*Paragraph {
	r := []*Paragraph{}

	paragraph := &Paragraph{Lines: Lines{}}
	for _, line := range lines {
		if len(bytes.TrimSpace(line.Content)) > 0 {
			paragraph.Lines = append(paragraph.Lines, line)
		} else if len(paragraph.Lines) > 0 {
			r = append(r, paragraph)
			paragraph = &Paragraph{Lines: Lines{}}
		}
	}

	if len(paragraph.Lines) > 0 {
		r = append(r, paragraph)
	}

	return r
}
