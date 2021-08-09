package conventionalcommit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewParagraphs(t *testing.T) {
	tests := []struct {
		name  string
		lines Lines
		want  []*Paragraph
	}{
		{
			name:  "nil",
			lines: nil,
			want:  []*Paragraph{},
		},
		{
			name:  "no lines",
			lines: Lines{},
			want:  []*Paragraph{},
		},
		{
			name: "single empty line",
			lines: Lines{
				{
					Number:  1,
					Content: []byte{},
					Break:   []byte{},
				},
			},
			want: []*Paragraph{},
		},
		{
			name: "multiple empty lines",
			lines: Lines{
				{
					Number:  1,
					Content: []byte{},
					Break:   []byte("\n"),
				},
				{
					Number:  2,
					Content: []byte{},
					Break:   []byte("\n"),
				},
				{
					Number:  3,
					Content: []byte{},
					Break:   []byte{},
				},
			},
			want: []*Paragraph{},
		},
		{
			name: "single whitespace line",
			lines: Lines{
				{
					Number:  1,
					Content: []byte("\t  "),
					Break:   []byte{},
				},
			},
			want: []*Paragraph{},
		},
		{
			name: "multiple whitespace lines",
			lines: Lines{
				{
					Number:  1,
					Content: []byte{},
					Break:   []byte("\t  "),
				},
				{
					Number:  2,
					Content: []byte{},
					Break:   []byte("\t  "),
				},
				{
					Number:  3,
					Content: []byte("\t  "),
					Break:   []byte{},
				},
			},
			want: []*Paragraph{},
		},
		{
			name: "single line",
			lines: Lines{
				{
					Number:  1,
					Content: []byte("hello world"),
					Break:   []byte{},
				},
			},
			want: []*Paragraph{
				{
					Lines: Lines{
						{
							Number:  1,
							Content: []byte("hello world"),
							Break:   []byte{},
						},
					},
				},
			},
		},
		{
			name: "multiple lines",
			lines: Lines{
				{
					Number:  1,
					Content: []byte("hello world"),
					Break:   []byte("\n"),
				},
				{
					Number:  2,
					Content: []byte("foo bar"),
					Break:   []byte{},
				},
			},
			want: []*Paragraph{
				{
					Lines: Lines{
						{
							Number:  1,
							Content: []byte("hello world"),
							Break:   []byte("\n"),
						},
						{
							Number:  2,
							Content: []byte("foo bar"),
							Break:   []byte{},
						},
					},
				},
			},
		},
		{
			name: "multiple lines with trailing line break",
			lines: Lines{
				{
					Number:  1,
					Content: []byte("hello world"),
					Break:   []byte("\n"),
				},
				{
					Number:  2,
					Content: []byte("foo bar"),
					Break:   []byte("\n"),
				},
				{
					Number:  3,
					Content: []byte(""),
					Break:   []byte{},
				},
			},
			want: []*Paragraph{
				{
					Lines: Lines{
						{
							Number:  1,
							Content: []byte("hello world"),
							Break:   []byte("\n"),
						},
						{
							Number:  2,
							Content: []byte("foo bar"),
							Break:   []byte("\n"),
						},
					},
				},
			},
		},
		{
			name: "multiple paragraphs with excess blank lines",
			lines: Lines{
				{
					Number:  1,
					Content: []byte(""),
					Break:   []byte("\n"),
				},
				{
					Number:  2,
					Content: []byte("\t  "),
					Break:   []byte("\r\n"),
				},
				{
					Number:  3,
					Content: []byte("Aliquam feugiat tellus ut neque."),
					Break:   []byte("\r"),
				},
				{
					Number:  4,
					Content: []byte("Sed bibendum."),
					Break:   []byte("\r"),
				},
				{
					Number:  5,
					Content: []byte("Nullam libero mauris, consequat."),
					Break:   []byte("\n"),
				},
				{
					Number:  6,
					Content: []byte(""),
					Break:   []byte("\n"),
				},
				{
					Number:  7,
					Content: []byte("Integer placerat tristique nisl."),
					Break:   []byte("\n"),
				},
				{
					Number:  8,
					Content: []byte("Etiam vel neque nec dui bibendum."),
					Break:   []byte("\n"),
				},
				{
					Number:  9,
					Content: []byte(""),
					Break:   []byte("\n"),
				},
				{
					Number:  10,
					Content: []byte("  "),
					Break:   []byte("\n"),
				},
				{
					Number:  11,
					Content: []byte("\t\t"),
					Break:   []byte("\n"),
				},
				{
					Number:  12,
					Content: []byte(""),
					Break:   []byte("\n"),
				},
				{
					Number:  13,
					Content: []byte("Donec hendrerit tempor tellus."),
					Break:   []byte("\n"),
				},
				{
					Number:  14,
					Content: []byte("In id erat non orci commodo lobortis."),
					Break:   []byte("\n"),
				},
				{
					Number:  15,
					Content: []byte(""),
					Break:   []byte("\n"),
				},
				{
					Number:  16,
					Content: []byte("  "),
					Break:   []byte("\n"),
				},
				{
					Number:  17,
					Content: []byte("\t\t"),
					Break:   []byte("\n"),
				},
				{
					Number:  18,
					Content: []byte(""),
					Break:   []byte("\n"),
				},
				{
					Number:  18,
					Content: []byte(""),
					Break:   []byte{},
				},
			},
			want: []*Paragraph{
				{
					Lines: Lines{
						{
							Number:  3,
							Content: []byte("Aliquam feugiat tellus ut neque."),
							Break:   []byte("\r"),
						},
						{
							Number:  4,
							Content: []byte("Sed bibendum."),
							Break:   []byte("\r"),
						},
						{
							Number:  5,
							Content: []byte("Nullam libero mauris, consequat."),
							Break:   []byte("\n"),
						},
					},
				},
				{
					Lines: Lines{
						{
							Number:  7,
							Content: []byte("Integer placerat tristique nisl."),
							Break:   []byte("\n"),
						},
						{
							Number: 8,
							Content: []byte(
								"Etiam vel neque nec dui bibendum.",
							),
							Break: []byte("\n"),
						},
					},
				},
				{
					Lines: Lines{
						{
							Number:  13,
							Content: []byte("Donec hendrerit tempor tellus."),
							Break:   []byte("\n"),
						},
						{
							Number: 14,
							Content: []byte(
								"In id erat non orci commodo lobortis.",
							),
							Break: []byte("\n"),
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewParagraphs(tt.lines)

			assert.Equal(t, tt.want, got)
		})
	}
}
