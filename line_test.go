package conventionalcommit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLine_Empty(t *testing.T) {
	tests := []struct {
		name string
		line *Line
		want bool
	}{
		{
			name: "nil",
			line: &Line{},
			want: true,
		},
		{
			name: "empty",
			line: &Line{
				Number:  1,
				Content: []byte(""),
				Break:   []byte{},
			},
			want: true,
		},
		{
			name: "space only",
			line: &Line{
				Number:  1,
				Content: []byte("  "),
				Break:   []byte{},
			},
			want: false,
		},
		{
			name: "tab only",
			line: &Line{
				Number:  1,
				Content: []byte("\t\t"),
				Break:   []byte{},
			},
			want: false,
		},
		{
			name: "spaces and tabs",
			line: &Line{
				Number:  1,
				Content: []byte(" \t "),
				Break:   []byte{},
			},
			want: false,
		},
		{
			name: "text",
			line: &Line{
				Number:  1,
				Content: []byte("foobar"),
				Break:   []byte{},
			},
			want: false,
		},
		{
			name: "text with surrounding white space",
			line: &Line{
				Number:  1,
				Content: []byte(" foobar  "),
				Break:   []byte{},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.line.Empty()

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestLine_Blank(t *testing.T) {
	tests := []struct {
		name string
		line *Line
		want bool
	}{
		{
			name: "nil",
			line: &Line{},
			want: true,
		},
		{
			name: "empty",
			line: &Line{
				Number:  1,
				Content: []byte(""),
				Break:   []byte{},
			},
			want: true,
		},
		{
			name: "space only",
			line: &Line{
				Number:  1,
				Content: []byte("  "),
				Break:   []byte{},
			},
			want: true,
		},
		{
			name: "tab only",
			line: &Line{
				Number:  1,
				Content: []byte("\t\t"),
				Break:   []byte{},
			},
			want: true,
		},
		{
			name: "spaces and tabs",
			line: &Line{
				Number:  1,
				Content: []byte(" \t "),
				Break:   []byte{},
			},
			want: true,
		},
		{
			name: "text",
			line: &Line{
				Number:  1,
				Content: []byte("foobar"),
				Break:   []byte{},
			},
			want: false,
		},
		{
			name: "text with surrounding white space",
			line: &Line{
				Number:  1,
				Content: []byte(" foobar  "),
				Break:   []byte{},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.line.Blank()

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestLine_Comment(t *testing.T) {
	tests := []struct {
		name string
		line *Line
		want bool
	}{
		{
			name: "nil",
			line: &Line{},
			want: false,
		},
		{
			name: "empty",
			line: &Line{
				Number:  1,
				Content: []byte(""),
				Break:   []byte{},
			},
			want: false,
		},
		{
			name: "space only",
			line: &Line{
				Number:  1,
				Content: []byte("  "),
				Break:   []byte{},
			},
			want: false,
		},
		{
			name: "tab only",
			line: &Line{
				Number:  1,
				Content: []byte("\t\t"),
				Break:   []byte{},
			},
			want: false,
		},
		{
			name: "spaces and tabs",
			line: &Line{
				Number:  1,
				Content: []byte(" \t "),
				Break:   []byte{},
			},
			want: false,
		},
		{
			name: "text",
			line: &Line{
				Number:  1,
				Content: []byte("foobar"),
				Break:   []byte{},
			},
			want: false,
		},
		{
			name: "beings with hash",
			line: &Line{
				Number:  1,
				Content: []byte("# foo bar"),
				Break:   []byte{},
			},
			want: true,
		},
		{
			name: "beings with hash after whitespace",
			line: &Line{
				Number:  1,
				Content: []byte("  \t  # foo bar"),
				Break:   []byte{},
			},
			want: true,
		},
		{
			name: "has hash in middle of text",
			line: &Line{
				Number:  1,
				Content: []byte("  foo # bar"),
				Break:   []byte{},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.line.Comment()

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestNewLines(t *testing.T) {
	tests := []struct {
		name    string
		content []byte
		want    Lines
	}{
		{
			name:    "nil",
			content: nil,
			want:    Lines{},
		},
		{
			name:    "empty",
			content: []byte{},
			want:    Lines{},
		},
		{
			name:    "single line without trailing linebreak",
			content: []byte("hello world"),
			want: Lines{
				{
					Number:  1,
					Content: []byte("hello world"),
					Break:   []byte{},
				},
			},
		},
		{
			name:    "single line with trailing LF",
			content: []byte("hello world\n"),
			want: Lines{
				{
					Number:  1,
					Content: []byte("hello world"),
					Break:   []byte("\n"),
				},
				{
					Number:  2,
					Content: []byte(""),
					Break:   []byte{},
				},
			},
		},
		{
			name:    "single line with trailing CRLF",
			content: []byte("hello world\r\n"),
			want: Lines{
				{
					Number:  1,
					Content: []byte("hello world"),
					Break:   []byte("\r\n"),
				},
				{
					Number:  2,
					Content: []byte(""),
					Break:   []byte{},
				},
			},
		},
		{
			name:    "single line with trailing CR",
			content: []byte("hello world\r"),
			want: Lines{
				{
					Number:  1,
					Content: []byte("hello world"),
					Break:   []byte("\r"),
				},
				{
					Number:  2,
					Content: []byte(""),
					Break:   []byte{},
				},
			},
		},
		{
			name:    "multiple lines separated by LF",
			content: []byte("hello world\nfoo\nbar"),
			want: Lines{
				{
					Number:  1,
					Content: []byte("hello world"),
					Break:   []byte("\n"),
				},
				{
					Number:  2,
					Content: []byte("foo"),
					Break:   []byte("\n"),
				},
				{
					Number:  3,
					Content: []byte("bar"),
					Break:   []byte{},
				},
			},
		},
		{
			name:    "multiple lines separated by LF with trailing LF",
			content: []byte("hello world\nfoo\nbar\n"),
			want: Lines{
				{
					Number:  1,
					Content: []byte("hello world"),
					Break:   []byte("\n"),
				},
				{
					Number:  2,
					Content: []byte("foo"),
					Break:   []byte("\n"),
				},
				{
					Number:  3,
					Content: []byte("bar"),
					Break:   []byte("\n"),
				},
				{
					Number:  4,
					Content: []byte(""),
					Break:   []byte{},
				},
			},
		},
		{
			name:    "multiple lines separated by CRLF",
			content: []byte("hello world\r\nfoo\r\nbar"),
			want: Lines{
				{
					Number:  1,
					Content: []byte("hello world"),
					Break:   []byte("\r\n"),
				},
				{
					Number:  2,
					Content: []byte("foo"),
					Break:   []byte("\r\n"),
				},
				{
					Number:  3,
					Content: []byte("bar"),
					Break:   []byte{},
				},
			},
		},
		{
			name:    "multiple lines separated by CRLF with trailing CRLF",
			content: []byte("hello world\r\nfoo\r\nbar\r\n"),
			want: Lines{
				{
					Number:  1,
					Content: []byte("hello world"),
					Break:   []byte("\r\n"),
				},
				{
					Number:  2,
					Content: []byte("foo"),
					Break:   []byte("\r\n"),
				},
				{
					Number:  3,
					Content: []byte("bar"),
					Break:   []byte("\r\n"),
				},
				{
					Number:  4,
					Content: []byte(""),
					Break:   []byte{},
				},
			},
		},
		{
			name:    "multiple lines separated by CR",
			content: []byte("hello world\rfoo\rbar"),
			want: Lines{
				{
					Number:  1,
					Content: []byte("hello world"),
					Break:   []byte("\r"),
				},
				{
					Number:  2,
					Content: []byte("foo"),
					Break:   []byte("\r"),
				},
				{
					Number:  3,
					Content: []byte("bar"),
					Break:   []byte{},
				},
			},
		},
		{
			name:    "multiple lines separated by CR with trailing CR",
			content: []byte("hello world\rfoo\rbar\r"),
			want: Lines{
				{
					Number:  1,
					Content: []byte("hello world"),
					Break:   []byte("\r"),
				},
				{
					Number:  2,
					Content: []byte("foo"),
					Break:   []byte("\r"),
				},
				{
					Number:  3,
					Content: []byte("bar"),
					Break:   []byte("\r"),
				},
				{
					Number:  4,
					Content: []byte(""),
					Break:   []byte{},
				},
			},
		},
		{
			name:    "multiple lines separated by mixed break types",
			content: []byte("hello\nworld\r\nfoo\rbar"),
			want: Lines{
				{
					Number:  1,
					Content: []byte("hello"),
					Break:   []byte("\n"),
				},
				{
					Number:  2,
					Content: []byte("world"),
					Break:   []byte("\r\n"),
				},
				{
					Number:  3,
					Content: []byte("foo"),
					Break:   []byte("\r"),
				},
				{
					Number:  4,
					Content: []byte("bar"),
					Break:   []byte{},
				},
			},
		},
		{
			name: "multiple lines separated by mixed break types with " +
				"trailing LF",
			content: []byte("hello\nworld\r\nfoo\rbar\n"),
			want: Lines{
				{
					Number:  1,
					Content: []byte("hello"),
					Break:   []byte("\n"),
				},
				{
					Number:  2,
					Content: []byte("world"),
					Break:   []byte("\r\n"),
				},
				{
					Number:  3,
					Content: []byte("foo"),
					Break:   []byte("\r"),
				},
				{
					Number:  4,
					Content: []byte("bar"),
					Break:   []byte("\n"),
				},
				{
					Number:  5,
					Content: []byte(""),
					Break:   []byte{},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewLines(tt.content)

			assert.Equal(t, tt.want, got)
		})
	}
}

var linesTestCases = []struct {
	name           string
	lines          Lines
	bytes          []byte
	firstTextIndex int
	lastTextIndex  int
}{
	{
		name:           "no lines",
		lines:          Lines{},
		bytes:          []byte(""),
		firstTextIndex: -1,
		lastTextIndex:  -1,
	},
	{
		name: "empty line",
		lines: Lines{
			{
				Number:  1,
				Content: []byte(""),
			},
		},
		bytes:          []byte(""),
		firstTextIndex: -1,
		lastTextIndex:  -1,
	},
	{
		name: "whitespace line",
		lines: Lines{
			{
				Number:  1,
				Content: []byte("  "),
			},
		},
		bytes:          []byte("  "),
		firstTextIndex: -1,
		lastTextIndex:  -1,
	},
	{
		name: "multiple whitespace lines",
		lines: Lines{
			{
				Number:  1,
				Content: []byte("  "),
				Break:   []byte("\n"),
			},
			{
				Number:  2,
				Content: []byte("\t"),
				Break:   []byte("\n"),
			},
			{
				Number:  3,
				Content: []byte("   "),
				Break:   []byte{},
			},
		},
		bytes:          []byte("  \n\t\n   "),
		firstTextIndex: -1,
		lastTextIndex:  -1,
	},
	{
		name: "single line",
		lines: Lines{
			{
				Number:  1,
				Content: []byte("hello world"),
			},
		},
		bytes:          []byte("hello world"),
		firstTextIndex: 0,
		lastTextIndex:  0,
	},
	{
		name: "single line with trailing LF",
		lines: Lines{
			{
				Number:  1,
				Content: []byte("hello world"),
				Break:   []byte("\n"),
			},
			{
				Number:  2,
				Content: []byte(""),
				Break:   []byte{},
			},
		},
		bytes:          []byte("hello world\n"),
		firstTextIndex: 0,
		lastTextIndex:  0,
	},
	{
		name: "single line with trailing CRLF",
		lines: Lines{
			{
				Number:  1,
				Content: []byte("hello world"),
				Break:   []byte("\r\n"),
			},
			{
				Number:  2,
				Content: []byte(""),
				Break:   []byte{},
			},
		},
		bytes:          []byte("hello world\r\n"),
		firstTextIndex: 0,
		lastTextIndex:  0,
	},
	{
		name: "single line with trailing CR",
		lines: Lines{
			{
				Number:  1,
				Content: []byte("hello world"),
				Break:   []byte("\r"),
			},
			{
				Number:  2,
				Content: []byte(""),
				Break:   []byte{},
			},
		},
		bytes:          []byte("hello world\r"),
		firstTextIndex: 0,
		lastTextIndex:  0,
	},
	{
		name: "multi-line separated by LF",
		lines: Lines{
			{
				Number:  1,
				Content: []byte(""),
				Break:   []byte("\n"),
			},
			{
				Number:  2,
				Content: []byte("Aliquam feugiat tellus ut neque."),
				Break:   []byte("\n"),
			},
			{
				Number:  3,
				Content: []byte("Sed bibendum."),
				Break:   []byte("\n"),
			},
			{
				Number:  4,
				Content: []byte("Nullam libero mauris, consequat."),
				Break:   []byte("\n"),
			},
			{
				Number:  5,
				Content: []byte(""),
				Break:   []byte("\n"),
			},
			{
				Number:  6,
				Content: []byte("Integer placerat tristique nisl."),
				Break:   []byte("\n"),
			},
			{
				Number:  7,
				Content: []byte("Etiam vel neque nec dui bibendum."),
				Break:   []byte("\n"),
			},
			{
				Number:  8,
				Content: []byte(""),
				Break:   []byte("\n"),
			},
			{
				Number:  9,
				Content: []byte(""),
				Break:   []byte("\n"),
			},
			{
				Number:  10,
				Content: []byte("Nullam libero mauris, dictum id, arcu."),
				Break:   []byte("\n"),
			},
			{
				Number:  11,
				Content: []byte(""),
				Break:   []byte{},
			},
		},
		bytes: []byte(
			"\n" +
				"Aliquam feugiat tellus ut neque.\n" +
				"Sed bibendum.\n" +
				"Nullam libero mauris, consequat.\n" +
				"\n" +
				"Integer placerat tristique nisl.\n" +
				"Etiam vel neque nec dui bibendum.\n" +
				"\n" +
				"\n" +
				"Nullam libero mauris, dictum id, arcu.\n",
		),
		firstTextIndex: 1,
		lastTextIndex:  9,
	},
	{
		name: "multi-line separated by CRLF",
		lines: Lines{
			{
				Number:  1,
				Content: []byte(""),
				Break:   []byte("\r\n"),
			},
			{
				Number:  2,
				Content: []byte("Aliquam feugiat tellus ut neque."),
				Break:   []byte("\r\n"),
			},
			{
				Number:  3,
				Content: []byte("Sed bibendum."),
				Break:   []byte("\r\n"),
			},
			{
				Number:  4,
				Content: []byte("Nullam libero mauris, consequat."),
				Break:   []byte("\r\n"),
			},
			{
				Number:  5,
				Content: []byte(""),
				Break:   []byte("\r\n"),
			},
			{
				Number:  6,
				Content: []byte("Integer placerat tristique nisl."),
				Break:   []byte("\r\n"),
			},
			{
				Number:  7,
				Content: []byte("Etiam vel neque nec dui bibendum."),
				Break:   []byte("\r\n"),
			},
			{
				Number:  8,
				Content: []byte(""),
				Break:   []byte("\r\n"),
			},
			{
				Number:  9,
				Content: []byte(""),
				Break:   []byte("\r\n"),
			},
			{
				Number:  10,
				Content: []byte("Nullam libero mauris, dictum id, arcu."),
				Break:   []byte("\r\n"),
			},
			{
				Number:  11,
				Content: []byte(""),
				Break:   []byte{},
			},
		},
		bytes: []byte(
			"\r\n" +
				"Aliquam feugiat tellus ut neque.\r\n" +
				"Sed bibendum.\r\n" +
				"Nullam libero mauris, consequat.\r\n" +
				"\r\n" +
				"Integer placerat tristique nisl.\r\n" +
				"Etiam vel neque nec dui bibendum.\r\n" +
				"\r\n" +
				"\r\n" +
				"Nullam libero mauris, dictum id, arcu.\r\n",
		),
		firstTextIndex: 1,
		lastTextIndex:  9,
	},
	{
		name: "multi-line separated by CR",
		lines: Lines{
			{
				Number:  1,
				Content: []byte(""),
				Break:   []byte("\r"),
			},
			{
				Number:  2,
				Content: []byte("Aliquam feugiat tellus ut neque."),
				Break:   []byte("\r"),
			},
			{
				Number:  3,
				Content: []byte("Sed bibendum."),
				Break:   []byte("\r"),
			},
			{
				Number:  4,
				Content: []byte("Nullam libero mauris, consequat."),
				Break:   []byte("\r"),
			},
			{
				Number:  5,
				Content: []byte(""),
				Break:   []byte("\r"),
			},
			{
				Number:  6,
				Content: []byte("Integer placerat tristique nisl."),
				Break:   []byte("\r"),
			},
			{
				Number:  7,
				Content: []byte("Etiam vel neque nec dui bibendum."),
				Break:   []byte("\r"),
			},
			{
				Number:  8,
				Content: []byte(""),
				Break:   []byte("\r"),
			},
			{
				Number:  9,
				Content: []byte(""),
				Break:   []byte("\r"),
			},
			{
				Number:  10,
				Content: []byte("Nullam libero mauris, dictum id, arcu."),
				Break:   []byte("\r"),
			},
			{
				Number:  11,
				Content: []byte(""),
				Break:   []byte{},
			},
		},
		bytes: []byte(
			"\r" +
				"Aliquam feugiat tellus ut neque.\r" +
				"Sed bibendum.\r" +
				"Nullam libero mauris, consequat.\r" +
				"\r" +
				"Integer placerat tristique nisl.\r" +
				"Etiam vel neque nec dui bibendum.\r" +
				"\r" +
				"\r" +
				"Nullam libero mauris, dictum id, arcu.\r",
		),
		firstTextIndex: 1,
		lastTextIndex:  9,
	},
}

func TestLines_FirstTextIndex(t *testing.T) {
	for _, tt := range linesTestCases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.lines.FirstTextIndex()

			assert.Equal(t, tt.firstTextIndex, got)
		})
	}
}

func BenchmarkLines_FirstTextIndex(b *testing.B) {
	for _, tt := range linesTestCases {
		b.Run(tt.name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				_ = tt.lines.FirstTextIndex()
			}
		})
	}
}

func TestLines_LastTextIndex(t *testing.T) {
	for _, tt := range linesTestCases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.lines.LastTextIndex()

			assert.Equal(t, tt.lastTextIndex, got)
		})
	}
}

func BenchmarkLines_LastTextIndex(b *testing.B) {
	for _, tt := range linesTestCases {
		b.Run(tt.name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				_ = tt.lines.LastTextIndex()
			}
		})
	}
}

func TestLines_Trim(t *testing.T) {
	for _, tt := range linesTestCases {
		t.Run(tt.name, func(t *testing.T) {
			want := Lines{}
			if tt.firstTextIndex != -1 {
				want = tt.lines[tt.firstTextIndex : tt.lastTextIndex+1]
			}

			got := tt.lines.Trim()

			assert.Equal(t, want, got)
		})
	}
}

func BenchmarkLines_Trim(b *testing.B) {
	for _, tt := range linesTestCases {
		b.Run(tt.name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				_ = tt.lines.Trim()
			}
		})
	}
}

func TestLines_Bytes(t *testing.T) {
	for _, tt := range linesTestCases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.lines.Bytes()

			assert.Equal(t, tt.bytes, got)
		})
	}
}

func BenchmarkLines_Bytes(b *testing.B) {
	for _, tt := range linesTestCases {
		b.Run(tt.name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				_ = tt.lines.Bytes()
			}
		})
	}
}

func TestLines_String(t *testing.T) {
	for _, tt := range linesTestCases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.lines.String()

			assert.Equal(t, string(tt.bytes), got)
		})
	}
}

func BenchmarkLines_String(b *testing.B) {
	for _, tt := range linesTestCases {
		b.Run(tt.name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				_ = tt.lines.String()
			}
		})
	}
}
