package conventionalcommit

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var (
	Err             = errors.New("conventionalcommit")
	ErrEmptyMessage = fmt.Errorf("%w: empty message", Err)
)

// HeaderToken will match a Conventional Commit formatted subject line, to
// extract type, scope, breaking change (bool), and description.
//
// It is intentionally VERY forgiving so as to be able to extract the various
// parts even when things aren't quite right.
var HeaderToken = regexp.MustCompile(
	`^([^\(\)\r\n]*?)(\((.*?)\)\s*)?(!)?(\s*\:)\s(.*)$`,
)

// FooterToken will match against all variations of Conventional Commit footer
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
var FooterToken = regexp.MustCompile(
	`^([\w-]+|BREAKING[\s-]CHANGE)(?:\s*(:)\s+|\s+(#))(.+)$`,
)

// Message represents a Conventional Commit message in a structured way.
type Message struct {
	// Type indicates what kind of a change the commit message describes.
	Type string

	// Scope indicates the context/component/area that the change affects.
	Scope string

	// Description is the primary description for the commit.
	Description string

	// Body is the main text body of the commit message. Effectively all text
	// between the subject line, and any footers if present.
	Body string

	// Footers are all footers which are not references or breaking changes.
	Footers []*Footer

	// References are all footers defined with a reference style token, for
	// example:
	//
	//  Fixes #42
	References []*Reference

	// Breaking is set to true if the message subject included the "!" breaking
	// change indicator.
	Breaking bool

	// BreakingChanges includes the descriptions from all BREAKING CHANGE
	// footers.
	BreakingChanges []string
}

func NewMessage(buf *Buffer) (*Message, error) {
	msg := &Message{}
	count := buf.LineCount()

	if count == 0 {
		return nil, ErrEmptyMessage
	}

	msg.Description = buf.Head().Join("\n")
	if m := HeaderToken.FindStringSubmatch(msg.Description); len(m) > 0 {
		msg.Type = strings.TrimSpace(m[1])
		msg.Scope = strings.TrimSpace(m[3])
		msg.Breaking = m[4] == "!"
		msg.Description = m[6]
	}

	msg.Body = buf.Body().Join("\n")

	if foot := buf.Foot(); len(foot) > 0 {
		footers := parseFooters(foot)

		for _, f := range footers {
			name := string(f.name)
			value := string(f.value)

			switch {
			case f.ref:
				msg.References = append(msg.References, &Reference{
					Name:  name,
					Value: value,
				})
			case name == "BREAKING CHANGE" || name == "BREAKING-CHANGE":
				msg.BreakingChanges = append(msg.BreakingChanges, value)
			default:
				msg.Footers = append(msg.Footers, &Footer{
					Name:  name,
					Value: value,
				})
			}
		}
	}

	return msg, nil
}

func (s *Message) IsBreakingChange() bool {
	return s.Breaking || len(s.BreakingChanges) > 0
}

func parseFooters(lines Lines) []*rawFooter {
	var footers []*rawFooter
	footer := &rawFooter{}
	for _, line := range lines {
		if m := FooterToken.FindSubmatch(line.Content); m != nil {
			if len(footer.name) > 0 {
				footers = append(footers, footer)
			}

			footer = &rawFooter{}
			if len(m[3]) > 0 {
				footer.ref = true
				footer.value = []byte{hash}
			}
			footer.name = m[1]
			footer.value = append(footer.value, m[4]...)
		} else if len(footer.name) > 0 {
			footer.value = append(footer.value, lf)
			footer.value = append(footer.value, line.Content...)
		}
	}

	if len(footer.name) > 0 {
		footers = append(footers, footer)
	}

	return footers
}

type rawFooter struct {
	name  []byte
	value []byte
	ref   bool
}

type Footer struct {
	Name  string
	Value string
}

type Reference struct {
	Name  string
	Value string
}
