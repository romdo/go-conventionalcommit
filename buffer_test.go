package conventionalcommit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var bufferTestCases = []struct {
	name       string
	bytes      []byte
	wantBuffer *Buffer
	wantHead   []int
	wantBody   []int
	wantFoot   []int
	wantLines  [2]int
}{
	{
		name:  "nil",
		bytes: nil,
		wantBuffer: &Buffer{
			lines: Lines{},
		},
		wantHead:  []int{},
		wantBody:  []int{},
		wantFoot:  []int{},
		wantLines: [2]int{0, 0},
	},
	{
		name:  "empty",
		bytes: []byte(""),
		wantBuffer: &Buffer{
			lines: Lines{},
		},
		wantHead:  []int{},
		wantBody:  []int{},
		wantFoot:  []int{},
		wantLines: [2]int{0, 0},
	},
	{
		name:  "single whitespace line",
		bytes: []byte(" "),
		wantBuffer: &Buffer{
			lines: Lines{
				{Number: 1, Content: []byte(" "), Break: []byte{}},
			},
		},
		wantHead:  []int{},
		wantBody:  []int{},
		wantFoot:  []int{},
		wantLines: [2]int{0, 0},
	},
	{
		name:  "multiple whitespace lines",
		bytes: []byte("\n\n  \n\n\t\n"),
		wantBuffer: &Buffer{
			lines: Lines{
				{Number: 1, Content: []byte(""), Break: []byte("\n")},
				{Number: 2, Content: []byte(""), Break: []byte("\n")},
				{Number: 3, Content: []byte("  "), Break: []byte("\n")},
				{Number: 4, Content: []byte(""), Break: []byte("\n")},
				{Number: 5, Content: []byte("\t"), Break: []byte("\n")},
				{Number: 6, Content: []byte(""), Break: []byte{}},
			},
		},
		wantHead:  []int{},
		wantBody:  []int{},
		wantFoot:  []int{},
		wantLines: [2]int{0, 0},
	},
	{
		name:  "single line",
		bytes: []byte("fix: a broken thing"),
		wantBuffer: &Buffer{
			headLen: 1,
			lines: Lines{
				{
					Number:  1,
					Content: []byte("fix: a broken thing"),
					Break:   []byte{},
				},
			},
		},
		wantHead:  []int{0},
		wantBody:  []int{},
		wantFoot:  []int{},
		wantLines: [2]int{0, 1},
	},
	{
		name:  "single line surrounded by whitespace",
		bytes: []byte("\n  \n\nfix: a broken thing\n\t\n"),
		wantBuffer: &Buffer{
			firstLine: 3,
			lastLine:  3,
			headLen:   1,
			lines: Lines{
				{Number: 1, Content: []byte(""), Break: []byte("\n")},
				{Number: 2, Content: []byte("  "), Break: []byte("\n")},
				{Number: 3, Content: []byte(""), Break: []byte("\n")},
				{
					Number:  4,
					Content: []byte("fix: a broken thing"),
					Break:   []byte("\n"),
				},
				{Number: 5, Content: []byte("\t"), Break: []byte("\n")},
				{Number: 6, Content: []byte(""), Break: []byte{}},
			},
		},
		wantHead:  []int{3},
		wantBody:  []int{},
		wantFoot:  []int{},
		wantLines: [2]int{3, 1},
	},
	{
		name:  "subject and body",
		bytes: []byte("fix: a broken thing\n\nIt is now fixed."),
		wantBuffer: &Buffer{
			firstLine: 0,
			lastLine:  2,
			headLen:   1,
			footLen:   0,
			lines: Lines{
				{
					Number:  1,
					Content: []byte("fix: a broken thing"),
					Break:   []byte("\n"),
				},
				{
					Number:  2,
					Content: []byte(""),
					Break:   []byte("\n"),
				},
				{
					Number:  3,
					Content: []byte("It is now fixed."),
					Break:   []byte{},
				},
			},
		},
		wantHead:  []int{0},
		wantBody:  []int{2},
		wantFoot:  []int{},
		wantLines: [2]int{0, 3},
	},
	{
		name: "subject and body with word footer token",
		bytes: []byte(`fix: a broken thing

It is now fixed.

Reviewed-by: John Carter`),
		wantBuffer: &Buffer{
			firstLine: 0,
			lastLine:  4,
			headLen:   1,
			footLen:   1,
			lines: Lines{
				{
					Number:  1,
					Content: []byte("fix: a broken thing"),
					Break:   []byte("\n"),
				},
				{
					Number:  2,
					Content: []byte(""),
					Break:   []byte("\n"),
				},
				{
					Number:  3,
					Content: []byte("It is now fixed."),
					Break:   []byte("\n"),
				},
				{
					Number:  4,
					Content: []byte(""),
					Break:   []byte("\n"),
				},
				{
					Number:  5,
					Content: []byte("Reviewed-by: John Carter"),
					Break:   []byte{},
				},
			},
		},
		wantHead:  []int{0},
		wantBody:  []int{2},
		wantFoot:  []int{4},
		wantLines: [2]int{0, 5},
	},
	{
		name: "subject and body with reference footer token",
		bytes: []byte(`fix: a broken thing

It is now fixed.

Fixes #39`),
		wantBuffer: &Buffer{
			firstLine: 0,
			lastLine:  4,
			headLen:   1,
			footLen:   1,
			lines: Lines{
				{
					Number:  1,
					Content: []byte("fix: a broken thing"),
					Break:   []byte("\n"),
				},
				{
					Number:  2,
					Content: []byte(""),
					Break:   []byte("\n"),
				},
				{
					Number:  3,
					Content: []byte("It is now fixed."),
					Break:   []byte("\n"),
				},
				{
					Number:  4,
					Content: []byte(""),
					Break:   []byte("\n"),
				},
				{
					Number:  5,
					Content: []byte("Fixes #39"),
					Break:   []byte{},
				},
			},
		},
		wantHead:  []int{0},
		wantBody:  []int{2},
		wantFoot:  []int{4},
		wantLines: [2]int{0, 5},
	},
	{
		name: "subject and body with BREAKING CHANGE footer",
		bytes: []byte(`refactor!: re-transpile the fugiator

This should improve performance.

BREAKING CHANGE: New argument is required, or BOOM!`),
		wantBuffer: &Buffer{
			firstLine: 0,
			lastLine:  4,
			headLen:   1,
			footLen:   1,
			lines: Lines{
				{
					Number:  1,
					Content: []byte("refactor!: re-transpile the fugiator"),
					Break:   []byte("\n"),
				},
				{
					Number:  2,
					Content: []byte(""),
					Break:   []byte("\n"),
				},
				{
					Number:  3,
					Content: []byte("This should improve performance."),
					Break:   []byte("\n"),
				},
				{
					Number:  4,
					Content: []byte(""),
					Break:   []byte("\n"),
				},
				{
					Number: 5,
					Content: []byte(
						"BREAKING CHANGE: New argument is required, or BOOM!",
					),
					Break: []byte{},
				},
			},
		},
		wantHead:  []int{0},
		wantBody:  []int{2},
		wantFoot:  []int{4},
		wantLines: [2]int{0, 5},
	},
	{
		name: "subject and body with BREAKING-CHANGE footer",
		bytes: []byte(`refactor!: re-transpile the fugiator

This should improve performance.

BREAKING-CHANGE: New argument is required, or BOOM!`),
		wantBuffer: &Buffer{
			firstLine: 0,
			lastLine:  4,
			headLen:   1,
			footLen:   1,
			lines: Lines{
				{
					Number:  1,
					Content: []byte("refactor!: re-transpile the fugiator"),
					Break:   []byte("\n"),
				},
				{
					Number:  2,
					Content: []byte(""),
					Break:   []byte("\n"),
				},
				{
					Number:  3,
					Content: []byte("This should improve performance."),
					Break:   []byte("\n"),
				},
				{
					Number:  4,
					Content: []byte(""),
					Break:   []byte("\n"),
				},
				{
					Number: 5,
					Content: []byte(
						"BREAKING-CHANGE: New argument is required, or BOOM!",
					),
					Break: []byte{},
				},
			},
		},
		wantHead:  []int{0},
		wantBody:  []int{2},
		wantFoot:  []int{4},
		wantLines: [2]int{0, 5},
	},
	{
		name: "subject and body with invalid footer token",
		bytes: []byte(`refactor!: re-transpile the fugiator

This should improve performance.

Reviewed by: John Carter`),
		wantBuffer: &Buffer{
			firstLine: 0,
			lastLine:  4,
			headLen:   1,
			footLen:   0,
			lines: Lines{
				{
					Number:  1,
					Content: []byte("refactor!: re-transpile the fugiator"),
					Break:   []byte("\n"),
				},
				{
					Number:  2,
					Content: []byte(""),
					Break:   []byte("\n"),
				},
				{
					Number:  3,
					Content: []byte("This should improve performance."),
					Break:   []byte("\n"),
				},
				{
					Number:  4,
					Content: []byte(""),
					Break:   []byte("\n"),
				},
				{
					Number:  5,
					Content: []byte("Reviewed by: John Carter"),
					Break:   []byte{},
				},
			},
		},
		wantHead:  []int{0},
		wantBody:  []int{2, 3, 4},
		wantFoot:  []int{},
		wantLines: [2]int{0, 5},
	},
	{
		name: "subject and body with valid footer token on second line",
		bytes: []byte(`refactor!: re-transpile the fugiator

This should improve performance.

the invalid footer starts here
Reviewed-by: John Carter`),
		wantBuffer: &Buffer{
			firstLine: 0,
			lastLine:  5,
			headLen:   1,
			footLen:   0,
			lines: Lines{
				{
					Number:  1,
					Content: []byte("refactor!: re-transpile the fugiator"),
					Break:   []byte("\n"),
				},
				{
					Number:  2,
					Content: []byte(""),
					Break:   []byte("\n"),
				},
				{
					Number:  3,
					Content: []byte("This should improve performance."),
					Break:   []byte("\n"),
				},
				{
					Number:  4,
					Content: []byte(""),
					Break:   []byte("\n"),
				},
				{
					Number:  5,
					Content: []byte("the invalid footer starts here"),
					Break:   []byte("\n"),
				},
				{
					Number:  6,
					Content: []byte("Reviewed-by: John Carter"),
					Break:   []byte{},
				},
			},
		},
		wantHead:  []int{0},
		wantBody:  []int{2, 3, 4, 5},
		wantFoot:  []int{},
		wantLines: [2]int{0, 6},
	},
	{
		name:  "subject and body with CRLF line breaks",
		bytes: []byte("fix: a broken thing\r\n\r\nIt is now fixed."),
		wantBuffer: &Buffer{
			firstLine: 0,
			lastLine:  2,
			headLen:   1,
			footLen:   0,
			lines: Lines{
				{
					Number:  1,
					Content: []byte("fix: a broken thing"),
					Break:   []byte("\r\n"),
				},
				{
					Number:  2,
					Content: []byte(""),
					Break:   []byte("\r\n"),
				},
				{
					Number:  3,
					Content: []byte("It is now fixed."),
					Break:   []byte{},
				},
			},
		},
		wantHead:  []int{0},
		wantBody:  []int{2},
		wantFoot:  []int{},
		wantLines: [2]int{0, 3},
	},
	{
		name:  "subject and body with CR line breaks",
		bytes: []byte("fix: a broken thing\r\rIt is now fixed."),
		wantBuffer: &Buffer{
			firstLine: 0,
			lastLine:  2,
			headLen:   1,
			footLen:   0,
			lines: Lines{
				{
					Number:  1,
					Content: []byte("fix: a broken thing"),
					Break:   []byte("\r"),
				},
				{
					Number:  2,
					Content: []byte(""),
					Break:   []byte("\r"),
				},
				{
					Number:  3,
					Content: []byte("It is now fixed."),
					Break:   []byte{},
				},
			},
		},
		wantHead:  []int{0},
		wantBody:  []int{2},
		wantFoot:  []int{},
		wantLines: [2]int{0, 3},
	},
	{
		name:  "separated by whitespace line",
		bytes: []byte("fix: a broken thing\n  \nIt is now fixed."),
		wantBuffer: &Buffer{
			firstLine: 0,
			lastLine:  2,
			headLen:   1,
			footLen:   0,
			lines: Lines{
				{
					Number:  1,
					Content: []byte("fix: a broken thing"),
					Break:   []byte("\n"),
				},
				{
					Number:  2,
					Content: []byte("  "),
					Break:   []byte("\n"),
				},
				{
					Number:  3,
					Content: []byte("It is now fixed."),
					Break:   []byte{},
				},
			},
		},
		wantHead:  []int{0},
		wantBody:  []int{2},
		wantFoot:  []int{},
		wantLines: [2]int{0, 3},
	},
	{
		name: "multi-line head and body",
		bytes: []byte(`

foo
bar

foz
baz

hello
world

`),
		wantBuffer: &Buffer{
			firstLine: 2,
			lastLine:  9,
			headLen:   2,
			footLen:   0,
			lines: Lines{
				{Number: 1, Content: []byte(""), Break: []byte("\n")},
				{Number: 2, Content: []byte(""), Break: []byte("\n")},
				{Number: 3, Content: []byte("foo"), Break: []byte("\n")},
				{Number: 4, Content: []byte("bar"), Break: []byte("\n")},
				{Number: 5, Content: []byte(""), Break: []byte("\n")},
				{Number: 6, Content: []byte("foz"), Break: []byte("\n")},
				{Number: 7, Content: []byte("baz"), Break: []byte("\n")},
				{Number: 8, Content: []byte(""), Break: []byte("\n")},
				{Number: 9, Content: []byte("hello"), Break: []byte("\n")},
				{Number: 10, Content: []byte("world"), Break: []byte("\n")},
				{Number: 11, Content: []byte(""), Break: []byte("\n")},
				{Number: 12, Content: []byte(""), Break: []byte{}},
			},
		},
		wantHead:  []int{2, 3},
		wantBody:  []int{5, 6, 7, 8, 9},
		wantFoot:  []int{},
		wantLines: [2]int{2, 8},
	},
	{
		name: "body surrounded by whitespace lines",
		bytes: []byte(`

foo
bar



foz
baz



hello
world


`),
		wantBuffer: &Buffer{
			firstLine: 2,
			lastLine:  13,
			headLen:   2,
			footLen:   0,
			lines: Lines{
				{Number: 1, Content: []byte(""), Break: []byte("\n")},
				{Number: 2, Content: []byte(""), Break: []byte("\n")},
				{Number: 3, Content: []byte("foo"), Break: []byte("\n")},
				{Number: 4, Content: []byte("bar"), Break: []byte("\n")},
				{Number: 5, Content: []byte(""), Break: []byte("\n")},
				{Number: 6, Content: []byte(""), Break: []byte("\n")},
				{Number: 7, Content: []byte(""), Break: []byte("\n")},
				{Number: 8, Content: []byte("foz"), Break: []byte("\n")},
				{Number: 9, Content: []byte("baz"), Break: []byte("\n")},
				{Number: 10, Content: []byte(""), Break: []byte("\n")},
				{Number: 11, Content: []byte(""), Break: []byte("\n")},
				{Number: 12, Content: []byte(""), Break: []byte("\n")},
				{Number: 13, Content: []byte("hello"), Break: []byte("\n")},
				{Number: 14, Content: []byte("world"), Break: []byte("\n")},
				{Number: 15, Content: []byte(""), Break: []byte("\n")},
				{Number: 16, Content: []byte(""), Break: []byte("\n")},
				{Number: 17, Content: []byte(""), Break: []byte{}},
			},
		},
		wantHead:  []int{2, 3},
		wantBody:  []int{7, 8, 9, 10, 11, 12, 13},
		wantFoot:  []int{},
		wantLines: [2]int{2, 12},
	},
	{
		name: "whitespace-only body",
		bytes: []byte(`

foo
bar




Approved-by: John Smith

`),
		wantBuffer: &Buffer{
			firstLine: 2,
			lastLine:  8,
			headLen:   2,
			footLen:   1,
			lines: Lines{
				{Number: 1, Content: []byte(""), Break: []byte("\n")},
				{Number: 2, Content: []byte(""), Break: []byte("\n")},
				{Number: 3, Content: []byte("foo"), Break: []byte("\n")},
				{Number: 4, Content: []byte("bar"), Break: []byte("\n")},
				{Number: 5, Content: []byte(""), Break: []byte("\n")},
				{Number: 6, Content: []byte(""), Break: []byte("\n")},
				{Number: 7, Content: []byte(""), Break: []byte("\n")},
				{Number: 8, Content: []byte(""), Break: []byte("\n")},
				{
					Number:  9,
					Content: []byte("Approved-by: John Smith"),
					Break:   []byte("\n"),
				},
				{Number: 10, Content: []byte(""), Break: []byte("\n")},
				{Number: 11, Content: []byte(""), Break: []byte{}},
			},
		},
		wantHead:  []int{2, 3},
		wantBody:  []int{},
		wantFoot:  []int{8},
		wantLines: [2]int{2, 7},
	},
	{
		name: "subject and body surrounded by whitespace",
		bytes: []byte(
			"\n  \nfix: a broken thing\n\nIt is now fixed.\n  \n\n",
		),
		wantBuffer: &Buffer{
			firstLine: 2,
			lastLine:  4,
			headLen:   1,
			footLen:   0,
			lines: Lines{
				{
					Number:  1,
					Content: []byte(""),
					Break:   []byte("\n"),
				},
				{
					Number:  2,
					Content: []byte("  "),
					Break:   []byte("\n"),
				},
				{
					Number:  3,
					Content: []byte("fix: a broken thing"),
					Break:   []byte("\n"),
				},
				{
					Number:  4,
					Content: []byte(""),
					Break:   []byte("\n"),
				},
				{
					Number:  5,
					Content: []byte("It is now fixed."),
					Break:   []byte("\n"),
				},
				{
					Number:  6,
					Content: []byte("  "),
					Break:   []byte("\n"),
				},
				{
					Number:  7,
					Content: []byte(""),
					Break:   []byte("\n"),
				},
				{
					Number:  8,
					Content: []byte(""),
					Break:   []byte{},
				},
			},
		},
		wantHead:  []int{2},
		wantBody:  []int{4},
		wantFoot:  []int{},
		wantLines: [2]int{2, 3},
	},
	{
		name: "subject and long body",
		bytes: []byte(`fix: something broken

Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Donec hendrerit
tempor tellus. Donec pretium posuere tellus. Proin quam nisl, tincidunt et,
mattis eget, convallis nec, purus. Cum sociis natoque penatibus et magnis dis
parturient montes, nascetur ridiculous mus. Nulla posuere. Donec vitae dolor.
Nullam tristique diam non turpis. Cras placerat accumsan nulla. Nullam rutrum.
Nam vestibulum accumsan nisl.

Nullam eu ante vel est convallis dignissim. Fusce suscipit, wisi nec facilisis
facilisis, est dui fermentum leo, quis tempor ligula erat quis odio. Nunc porta
vulputate tellus. Nunc rutrum turpis sed pede. Sed bibendum. Aliquam posuere.
Nunc aliquet, augue nec adipiscing interdum, lacus tellus malesuada massa, quis
varius mi purus non odio.

Phasellus lacus. Nam euismod tellus id erat. Pellentesque condimentum, magna ut
suscipit hendrerit, ipsum augue ornare nulla, non luctus diam neque sit amet
urna. Curabitur vulputate vestibulum lorem. Fusce sagittis, libero non molestie
mollis, magna orci ultrices dolor, at vulputate neque nulla lacinia eros. Sed id
ligula quis est convallis tempor. Curabitur lacinia pulvinar nibh. Nam a
sapien.`),
		wantBuffer: &Buffer{
			firstLine: 0,
			lastLine:  20,
			headLen:   1,
			footLen:   0,
			lines: Lines{
				{
					Number:  1,
					Content: []byte("fix: something broken"),
					Break:   []byte("\n"),
				},
				{Number: 2, Content: []byte(""), Break: []byte("\n")},
				{
					Number: 3,
					Content: []byte(
						"Lorem ipsum dolor sit amet, consectetuer " +
							"adipiscing elit. Donec hendrerit"),
					Break: []byte("\n"),
				},
				{
					Number: 4,
					Content: []byte(
						"tempor tellus. Donec pretium posuere tellus. " +
							"Proin quam nisl, tincidunt et,"),
					Break: []byte("\n"),
				},
				{
					Number: 5,
					Content: []byte(
						"mattis eget, convallis nec, purus. Cum sociis " +
							"natoque penatibus et magnis dis"),
					Break: []byte("\n"),
				},
				{
					Number: 6,
					Content: []byte(
						"parturient montes, nascetur ridiculous mus. " +
							"Nulla posuere. Donec vitae dolor."),
					Break: []byte("\n"),
				},
				{
					Number: 7,
					Content: []byte(
						"Nullam tristique diam non turpis. Cras placerat " +
							"accumsan nulla. Nullam rutrum."),
					Break: []byte("\n"),
				},
				{
					Number: 8,
					Content: []byte(
						"Nam vestibulum accumsan nisl."),
					Break: []byte("\n"),
				},
				{Number: 9, Content: []byte(""), Break: []byte("\n")},
				{
					Number: 10,
					Content: []byte(
						"Nullam eu ante vel est convallis dignissim. " +
							"Fusce suscipit, wisi nec facilisis",
					),
					Break: []byte("\n"),
				},
				{
					Number: 11,
					Content: []byte(
						"facilisis, est dui fermentum leo, quis tempor " +
							"ligula erat quis odio. Nunc porta",
					),
					Break: []byte("\n"),
				},
				{
					Number: 12,
					Content: []byte(
						"vulputate tellus. Nunc rutrum turpis sed pede. " +
							"Sed bibendum. Aliquam posuere.",
					),
					Break: []byte("\n"),
				},
				{
					Number: 13,
					Content: []byte(
						"Nunc aliquet, augue nec adipiscing interdum, " +
							"lacus tellus malesuada massa, quis",
					),
					Break: []byte("\n"),
				},
				{
					Number:  14,
					Content: []byte("varius mi purus non odio."),
					Break:   []byte("\n"),
				},
				{Number: 15, Content: []byte(""), Break: []byte("\n")},
				{
					Number: 16,
					Content: []byte("Phasellus lacus. Nam euismod tellus id " +
						"erat. Pellentesque condimentum, magna ut"),
					Break: []byte("\n"),
				},
				{
					Number: 17,
					Content: []byte("suscipit hendrerit, ipsum augue ornare " +
						"nulla, non luctus diam neque sit amet"),
					Break: []byte("\n"),
				},
				{
					Number: 18,
					Content: []byte("urna. Curabitur vulputate vestibulum " +
						"lorem. Fusce sagittis, libero non molestie"),
					Break: []byte("\n"),
				},
				{
					Number: 19,
					Content: []byte("mollis, magna orci ultrices dolor, at " +
						"vulputate neque nulla lacinia eros. Sed id"),
					Break: []byte("\n"),
				},
				{
					Number: 20,
					Content: []byte("ligula quis est convallis tempor. " +
						"Curabitur lacinia pulvinar nibh. Nam a"),
					Break: []byte("\n"),
				},
				{
					Number:  21,
					Content: []byte("sapien."),
					Break:   []byte{},
				},
			},
		},
		wantHead: []int{0},
		wantBody: []int{
			2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
		},
		wantFoot:  []int{},
		wantLines: [2]int{0, 21},
	},
}

func TestNewBuffer(t *testing.T) {
	for _, tt := range bufferTestCases {
		t.Run(tt.name, func(t *testing.T) {
			got := NewBuffer(tt.bytes)

			assert.Equal(t, tt.wantBuffer, got)
		})
	}
}

func BenchmarkNewBuffer(b *testing.B) {
	for _, tt := range bufferTestCases {
		b.Run(tt.name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				_ = NewBuffer(tt.bytes)
			}
		})
	}
}

func TestBuffer_Head(t *testing.T) {
	for _, tt := range bufferTestCases {
		t.Run(tt.name, func(t *testing.T) {
			want := Lines{}
			for _, i := range tt.wantHead {
				want = append(want, tt.wantBuffer.lines[i])
			}

			got := tt.wantBuffer.Head()

			assert.Equal(t, want, got)
		})
	}
}

func BenchmarkBuffer_Head(b *testing.B) {
	for _, tt := range bufferTestCases {
		if tt.bytes == nil {
			continue
		}
		b.Run(tt.name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				_ = tt.wantBuffer.Head()
			}
		})
	}
}

func TestBuffer_Body(t *testing.T) {
	for _, tt := range bufferTestCases {
		t.Run(tt.name, func(t *testing.T) {
			want := Lines{}
			for _, i := range tt.wantBody {
				want = append(want, tt.wantBuffer.lines[i])
			}

			got := tt.wantBuffer.Body()

			assert.Equal(t, want, got)
		})
	}
}

func BenchmarkBuffer_Body(b *testing.B) {
	for _, tt := range bufferTestCases {
		if tt.bytes == nil {
			continue
		}
		b.Run(tt.name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				_ = tt.wantBuffer.Body()
			}
		})
	}
}

func TestBuffer_Foot(t *testing.T) {
	for _, tt := range bufferTestCases {
		t.Run(tt.name, func(t *testing.T) {
			want := Lines{}
			for _, i := range tt.wantFoot {
				want = append(want, tt.wantBuffer.lines[i])
			}

			got := tt.wantBuffer.Foot()

			assert.Equal(t, want, got)
		})
	}
}

func BenchmarkBuffer_Foot(b *testing.B) {
	for _, tt := range bufferTestCases {
		if tt.bytes == nil {
			continue
		}
		b.Run(tt.name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				_ = tt.wantBuffer.Foot()
			}
		})
	}
}

func TestBuffer_Lines(t *testing.T) {
	for _, tt := range bufferTestCases {
		t.Run(tt.name, func(t *testing.T) {
			start := tt.wantLines[0]
			end := tt.wantLines[0] + tt.wantLines[1]
			want := tt.wantBuffer.lines[start:end]

			got := tt.wantBuffer.Lines()

			assert.Equal(t, want, got)
		})
	}
}

func BenchmarkBuffer_Lines(b *testing.B) {
	for _, tt := range bufferTestCases {
		if tt.bytes == nil {
			continue
		}
		b.Run(tt.name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				_ = tt.wantBuffer.Lines()
			}
		})
	}
}

func TestBuffer_Bytes(t *testing.T) {
	for _, tt := range bufferTestCases {
		if tt.bytes == nil {
			continue
		}
		t.Run(tt.name, func(t *testing.T) {
			start := tt.wantLines[0]
			end := tt.wantLines[0] + tt.wantLines[1]
			want := tt.wantBuffer.lines[start:end].Bytes()

			got := tt.wantBuffer.Bytes()

			assert.Equal(t, want, got)
		})
	}
}

func BenchmarkMessage_Bytes(b *testing.B) {
	for _, tt := range bufferTestCases {
		if tt.bytes == nil {
			continue
		}
		b.Run(tt.name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				_ = tt.wantBuffer.Bytes()
			}
		})
	}
}

func TestBuffer_String(t *testing.T) {
	for _, tt := range bufferTestCases {
		if tt.bytes == nil {
			continue
		}
		t.Run(tt.name, func(t *testing.T) {
			start := tt.wantLines[0]
			end := tt.wantLines[0] + tt.wantLines[1]
			want := tt.wantBuffer.lines[start:end].String()

			got := tt.wantBuffer.String()

			assert.Equal(t, want, got)
		})
	}
}

func BenchmarkMessage_String(b *testing.B) {
	for _, tt := range bufferTestCases {
		if tt.bytes == nil {
			continue
		}
		b.Run(tt.name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				_ = tt.wantBuffer.String()
			}
		})
	}
}

func TestBuffer_BytesRaw(t *testing.T) {
	for _, tt := range bufferTestCases {
		if tt.bytes == nil {
			continue
		}
		t.Run(tt.name, func(t *testing.T) {
			got := tt.wantBuffer.BytesRaw()

			assert.Equal(t, tt.bytes, got)
		})
	}
}

func BenchmarkBuffer_BytesRaw(b *testing.B) {
	for _, tt := range bufferTestCases {
		if tt.bytes == nil {
			continue
		}
		b.Run(tt.name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				_ = tt.wantBuffer.BytesRaw()
			}
		})
	}
}

func TestBuffer_StringRaw(t *testing.T) {
	for _, tt := range bufferTestCases {
		if tt.bytes == nil {
			continue
		}
		t.Run(tt.name, func(t *testing.T) {
			got := tt.wantBuffer.StringRaw()

			assert.Equal(t, string(tt.bytes), got)
		})
	}
}

func BenchmarkBuffer_StringRaw(b *testing.B) {
	for _, tt := range bufferTestCases {
		if tt.bytes == nil {
			continue
		}
		b.Run(tt.name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				_ = tt.wantBuffer.StringRaw()
			}
		})
	}
}
