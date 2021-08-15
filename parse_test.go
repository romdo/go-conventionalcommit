package conventionalcommit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name    string
		message []byte
		want    *Message
		wantErr string
	}{
		{
			name:    "empty",
			message: []byte{},
			wantErr: "conventionalcommit: empty message",
		},
		{
			name:    "description only",
			message: []byte("change a thing"),
			want: &Message{
				Description: "change a thing",
			},
		},
		{
			name: "description and body",
			message: []byte(`change a thing

more stuff
and more`,
			),
			want: &Message{
				Description: "change a thing",
				Body:        "more stuff\nand more",
			},
		},
		{
			name:    "type and description",
			message: []byte("feat: change a thing"),
			want: &Message{
				Type:        "feat",
				Description: "change a thing",
			},
		},
		{
			name: "type, description and body",
			message: []byte(
				"feat: change a thing\n\nmore stuff\nand more",
			),
			want: &Message{
				Type:        "feat",
				Description: "change a thing",
				Body:        "more stuff\nand more",
			},
		},
		{
			name:    "type, scope and description",
			message: []byte("feat(token): change a thing"),
			want: &Message{
				Type:        "feat",
				Scope:       "token",
				Description: "change a thing",
			},
		},
		{
			name: "type, scope, description and body",
			message: []byte(
				`feat(token): change a thing

more stuff
and more`,
			),
			want: &Message{
				Type:        "feat",
				Scope:       "token",
				Description: "change a thing",
				Body:        "more stuff\nand more",
			},
		},
		{
			name: "breaking change in subject line",
			message: []byte(
				`feat!: change a thing

more stuff
and more`,
			),
			want: &Message{
				Type:        "feat",
				Description: "change a thing",
				Body:        "more stuff\nand more",
				Breaking:    true,
			},
		},
		{
			name: "breaking change in subject line with scope",
			message: []byte(
				`feat(token)!: change a thing

more stuff
and more`,
			),
			want: &Message{
				Type:        "feat",
				Scope:       "token",
				Description: "change a thing",
				Body:        "more stuff\nand more",
				Breaking:    true,
			},
		},

		{
			name: "BREAKING CHANGE footer",
			message: []byte(
				`feat: change a thing

BREAKING CHANGE: will blow up
`,
			),
			want: &Message{
				Type:            "feat",
				Description:     "change a thing",
				BreakingChanges: []string{"will blow up"},
			},
		},
		{
			name: "BREAKING-CHANGE footer",
			message: []byte(
				`feat(token): change a thing

BREAKING-CHANGE: maybe not
`,
			),
			want: &Message{
				Type:            "feat",
				Scope:           "token",
				Description:     "change a thing",
				BreakingChanges: []string{"maybe not"},
			},
		},
		{
			name: "reference footer",
			message: []byte(
				`feat: change a thing

Fixes #349
`,
			),
			want: &Message{
				Type:        "feat",
				Description: "change a thing",
				References: []*Reference{
					{Name: "Fixes", Value: "#349"},
				},
			},
		},
		{
			name: "reference (alt) footer",
			message: []byte(
				`feat: change a thing

Reverts #SOL-934
`,
			),
			want: &Message{
				Type:        "feat",
				Description: "change a thing",
				References: []*Reference{
					{Name: "Reverts", Value: "#SOL-934"},
				},
			},
		},
		{
			name: "token footer",
			message: []byte(
				`feat: change a thing

Approved-by: John Carter
`,
			),
			want: &Message{
				Type:        "feat",
				Description: "change a thing",
				Footers: []*Footer{
					{Name: "Approved-by", Value: "John Carter"},
				},
			},
		},
		{
			name: "token (alt) footer",
			message: []byte(
				`feat: change a thing

ReviewedBy: Noctis
`,
			),
			want: &Message{
				Type:        "feat",
				Description: "change a thing",
				Footers: []*Footer{
					{Name: "ReviewedBy", Value: "Noctis"},
				},
			},
		},

		{
			name: "BREAKING CHANGE footer with body",
			message: []byte(
				`feat: change a thing

more stuff
and more

BREAKING CHANGE: will blow up
`,
			),
			want: &Message{
				Type:            "feat",
				Description:     "change a thing",
				Body:            "more stuff\nand more",
				BreakingChanges: []string{"will blow up"},
			},
		},
		{
			name: "BREAKING-CHANGE footer with body",
			message: []byte(
				`feat(token): change a thing

more stuff
and more

BREAKING-CHANGE: maybe not
`,
			),
			want: &Message{
				Type:            "feat",
				Scope:           "token",
				Description:     "change a thing",
				Body:            "more stuff\nand more",
				BreakingChanges: []string{"maybe not"},
			},
		},
		{
			name: "reference footer with body",
			message: []byte(
				`feat: change a thing

more stuff
and more

Fixes #349
`,
			),
			want: &Message{
				Type:        "feat",
				Description: "change a thing",
				Body:        "more stuff\nand more",
				References: []*Reference{
					{Name: "Fixes", Value: "#349"},
				},
			},
		},
		{
			name: "reference (alt) footer with body",
			message: []byte(
				`feat: change a thing

more stuff
and more

Reverts #SOL-934
`,
			),
			want: &Message{
				Type:        "feat",
				Description: "change a thing",
				Body:        "more stuff\nand more",
				References: []*Reference{
					{Name: "Reverts", Value: "#SOL-934"},
				},
			},
		},
		{
			name: "token footer with body",
			message: []byte(
				`feat: change a thing

more stuff
and more

Approved-by: John Carter
`,
			),
			want: &Message{
				Type:        "feat",
				Description: "change a thing",
				Body:        "more stuff\nand more",
				Footers: []*Footer{
					{Name: "Approved-by", Value: "John Carter"},
				},
			},
		},
		{
			name: "token (alt) footer with body",
			message: []byte(
				`feat: change a thing

more stuff
and more

ReviewedBy: Noctis
`,
			),
			want: &Message{
				Type:        "feat",
				Description: "change a thing",
				Body:        "more stuff\nand more",
				Footers: []*Footer{
					{Name: "ReviewedBy", Value: "Noctis"},
				},
			},
		},
		{
			name: "type, scope, description, body and footers",
			message: []byte(
				`feat(token): change a thing

more stuff
and more

BREAKING CHANGE: will blow up
BREAKING-CHANGE: maybe not
Fixes #349
Reverts #SOL-934
Approved-by: John Carter
ReviewedBy: Noctis
`,
			),
			want: &Message{
				Type:        "feat",
				Scope:       "token",
				Description: "change a thing",
				Body:        "more stuff\nand more",
				Footers: []*Footer{
					{Name: "Approved-by", Value: "John Carter"},
					{Name: "ReviewedBy", Value: "Noctis"},
				},
				References: []*Reference{
					{Name: "Fixes", Value: "#349"},
					{Name: "Reverts", Value: "#SOL-934"},
				},
				BreakingChanges: []string{"will blow up", "maybe not"},
			},
		},
		{
			name: "multi-line footers",
			message: []byte(
				`feat(token): change a thing

Some stuff

BREAKING CHANGE: Nam euismod tellus id erat.  Cum sociis natoque penatibus
et magnis dis parturient montes, nascetur ridiculous mus.
Approved-by: John Carter
and Noctis
Fixes #SOL-349 and also
#SOL-9440
`,
			),
			want: &Message{
				Type:        "feat",
				Scope:       "token",
				Description: "change a thing",
				Body:        "Some stuff",
				Footers: []*Footer{
					{Name: "Approved-by", Value: "John Carter\nand Noctis"},
				},
				References: []*Reference{
					{Name: "Fixes", Value: "#SOL-349 and also\n#SOL-9440"},
				},
				BreakingChanges: []string{
					`Nam euismod tellus id erat.  Cum sociis natoque penatibus
et magnis dis parturient montes, nascetur ridiculous mus.`,
				},
			},
		},
		{
			name: "indented footer",
			message: []byte(
				`feat(token): change a thing

Some stuff

    Approved-by: John Carter
`,
			),
			want: &Message{
				Type:        "feat",
				Scope:       "token",
				Description: "change a thing",
				Body:        "Some stuff\n\n    Approved-by: John Carter",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.message)

			if tt.wantErr != "" {
				assert.EqualError(t, err, tt.wantErr)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
