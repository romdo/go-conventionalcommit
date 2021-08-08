package conventionalcommit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var rawMessageTestCases = []struct {
	name       string
	bytes      []byte
	rawMessage *RawMessage
}{
	{
		name:  "nil",
		bytes: nil,
		rawMessage: &RawMessage{
			Lines:      Lines{},
			Paragraphs: []*Paragraph{},
		},
	},
	{
		name:  "empty",
		bytes: []byte(""),
		rawMessage: &RawMessage{
			Lines:      Lines{},
			Paragraphs: []*Paragraph{},
		},
	},
	{
		name:  "single space",
		bytes: []byte(" "),
		rawMessage: &RawMessage{
			Lines: Lines{
				{
					Number:  1,
					Content: []byte(" "),
					Break:   []byte{},
				},
			},
			Paragraphs: []*Paragraph{},
		},
	},
	{
		name:  "subject only",
		bytes: []byte("fix: a broken thing"),
		rawMessage: &RawMessage{
			Lines: Lines{
				{
					Number:  1,
					Content: []byte("fix: a broken thing"),
					Break:   []byte{},
				},
			},
			Paragraphs: []*Paragraph{
				{
					Lines: Lines{
						{
							Number:  1,
							Content: []byte("fix: a broken thing"),
							Break:   []byte{},
						},
					},
				},
			},
		},
	},
	{
		name:  "subject and body",
		bytes: []byte("fix: a broken thing\n\nIt is now fixed."),
		rawMessage: &RawMessage{
			Lines: Lines{
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
			Paragraphs: []*Paragraph{
				{
					Lines: Lines{
						{
							Number:  1,
							Content: []byte("fix: a broken thing"),
							Break:   []byte("\n"),
						},
					},
				},
				{
					Lines: Lines{
						{
							Number:  3,
							Content: []byte("It is now fixed."),
							Break:   []byte{},
						},
					},
				},
			},
		},
	},
	{
		name:  "subject and body with CRLF line breaks",
		bytes: []byte("fix: a broken thing\r\n\r\nIt is now fixed."),
		rawMessage: &RawMessage{
			Lines: Lines{
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
			Paragraphs: []*Paragraph{
				{
					Lines: Lines{
						{
							Number:  1,
							Content: []byte("fix: a broken thing"),
							Break:   []byte("\r\n"),
						},
					},
				},
				{
					Lines: Lines{
						{
							Number:  3,
							Content: []byte("It is now fixed."),
							Break:   []byte{},
						},
					},
				},
			},
		},
	},
	{
		name:  "subject and body with CR line breaks",
		bytes: []byte("fix: a broken thing\r\rIt is now fixed."),
		rawMessage: &RawMessage{
			Lines: Lines{
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
			Paragraphs: []*Paragraph{
				{
					Lines: Lines{
						{
							Number:  1,
							Content: []byte("fix: a broken thing"),
							Break:   []byte("\r"),
						},
					},
				},
				{
					Lines: Lines{
						{
							Number:  3,
							Content: []byte("It is now fixed."),
							Break:   []byte{},
						},
					},
				},
			},
		},
	},
	{
		name:  "separated by whitespace line",
		bytes: []byte("fix: a broken thing\n  \nIt is now fixed."),
		rawMessage: &RawMessage{
			Lines: Lines{
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
			Paragraphs: []*Paragraph{
				{
					Lines: Lines{
						{
							Number:  1,
							Content: []byte("fix: a broken thing"),
							Break:   []byte("\n"),
						},
					},
				},
				{
					Lines: Lines{
						{
							Number:  3,
							Content: []byte("It is now fixed."),
							Break:   []byte{},
						},
					},
				},
			},
		},
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
varius mi purus non odio. Pellentesque condimentum, magna ut suscipit hendrerit,
ipsum augue ornare nulla, non luctus diam neque sit amet urna. Curabitur
vulputate vestibulum lorem. Fusce sagittis, libero non molestie mollis, magna
orci ultrices dolor, at vulputate neque nulla lacinia eros. Sed id ligula quis
est convallis tempor. Curabitur lacinia pulvinar nibh. Nam a sapien.

Phasellus lacus. Nam euismod tellus id erat.`),
		rawMessage: &RawMessage{
			Lines: Lines{
				{
					Number:  1,
					Content: []byte("fix: something broken"),
					Break:   []byte("\n"),
				},
				{
					Number:  2,
					Content: []byte(""),
					Break:   []byte("\n"),
				},
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
				{
					Number:  9,
					Content: []byte(""),
					Break:   []byte("\n"),
				},
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
					Number: 14,
					Content: []byte(
						"varius mi purus non odio. Pellentesque " +
							"condimentum, magna ut suscipit hendrerit,",
					),
					Break: []byte("\n"),
				},
				{
					Number: 15,
					Content: []byte(
						"ipsum augue ornare nulla, non luctus diam neque " +
							"sit amet urna. Curabitur",
					),
					Break: []byte("\n"),
				},
				{
					Number: 16,
					Content: []byte(
						"vulputate vestibulum lorem. Fusce sagittis, " +
							"libero non molestie mollis, magna",
					),
					Break: []byte("\n"),
				},
				{
					Number: 17,
					Content: []byte(
						"orci ultrices dolor, at vulputate neque nulla " +
							"lacinia eros. Sed id ligula quis",
					),
					Break: []byte("\n"),
				},
				{
					Number: 18,
					Content: []byte(
						"est convallis tempor. Curabitur lacinia " +
							"pulvinar nibh. Nam a sapien.",
					),
					Break: []byte("\n"),
				},
				{
					Number:  19,
					Content: []byte(""),
					Break:   []byte("\n"),
				},
				{
					Number: 20,
					Content: []byte(
						"Phasellus lacus. Nam euismod tellus id erat.",
					),
					Break: []byte{},
				},
			},
			Paragraphs: []*Paragraph{
				{
					Lines: Lines{
						{
							Number:  1,
							Content: []byte("fix: something broken"),
							Break:   []byte("\n"),
						},
					},
				},
				{
					Lines: Lines{
						{
							Number: 3,
							Content: []byte(
								"Lorem ipsum dolor sit amet, " +
									"consectetuer adipiscing elit. Donec " +
									"hendrerit",
							),
							Break: []byte("\n"),
						},
						{
							Number: 4,
							Content: []byte(
								"tempor tellus. Donec pretium posuere " +
									"tellus. Proin quam nisl, tincidunt " +
									"et,",
							),
							Break: []byte("\n"),
						},
						{
							Number: 5,
							Content: []byte(
								"mattis eget, convallis nec, purus. Cum " +
									"sociis natoque penatibus et magnis " +
									"dis",
							),
							Break: []byte("\n"),
						},
						{
							Number: 6,
							Content: []byte(
								"parturient montes, nascetur ridiculous " +
									"mus. Nulla posuere. Donec vitae " +
									"dolor.",
							),
							Break: []byte("\n"),
						},
						{
							Number: 7,
							Content: []byte(
								"Nullam tristique diam non turpis. Cras " +
									"placerat accumsan nulla. Nullam " +
									"rutrum.",
							),
							Break: []byte("\n"),
						},
						{
							Number: 8,
							Content: []byte(
								"Nam vestibulum accumsan nisl.",
							),
							Break: []byte("\n"),
						},
					},
				},
				{
					Lines: Lines{
						{
							Number: 10,
							Content: []byte(
								"Nullam eu ante vel est convallis " +
									"dignissim. Fusce suscipit, wisi nec " +
									"facilisis",
							),
							Break: []byte("\n"),
						},
						{
							Number: 11,
							Content: []byte(
								"facilisis, est dui fermentum leo, quis " +
									"tempor ligula erat quis odio. Nunc " +
									"porta",
							),
							Break: []byte("\n"),
						},
						{
							Number: 12,
							Content: []byte(
								"vulputate tellus. Nunc rutrum turpis " +
									"sed pede. Sed bibendum. Aliquam " +
									"posuere.",
							),
							Break: []byte("\n"),
						},
						{
							Number: 13,
							Content: []byte(
								"Nunc aliquet, augue nec adipiscing " +
									"interdum, lacus tellus malesuada " +
									"massa, quis",
							),
							Break: []byte("\n"),
						},
						{
							Number: 14,
							Content: []byte(
								"varius mi purus non odio. Pellentesque " +
									"condimentum, magna ut suscipit " +
									"hendrerit,",
							),
							Break: []byte("\n"),
						},
						{
							Number: 15,
							Content: []byte(
								"ipsum augue ornare nulla, non luctus " +
									"diam neque sit amet urna. Curabitur",
							),
							Break: []byte("\n"),
						},
						{
							Number: 16,
							Content: []byte(
								"vulputate vestibulum lorem. Fusce " +
									"sagittis, libero non molestie " +
									"mollis, magna",
							),
							Break: []byte("\n"),
						},
						{
							Number: 17,
							Content: []byte(
								"orci ultrices dolor, at vulputate neque " +
									"nulla lacinia eros. Sed id ligula " +
									"quis",
							),
							Break: []byte("\n"),
						},
						{
							Number: 18,
							Content: []byte(
								"est convallis tempor. Curabitur lacinia " +
									"pulvinar nibh. Nam a sapien.",
							),
							Break: []byte("\n"),
						},
					},
				},
				{
					Lines: Lines{
						{
							Number: 20,
							Content: []byte(
								"Phasellus lacus. Nam euismod tellus id " +
									"erat.",
							),
							Break: []byte{},
						},
					},
				},
			},
		},
	},
}

func TestNewRawMessage(t *testing.T) {
	for _, tt := range rawMessageTestCases {
		t.Run(tt.name, func(t *testing.T) {
			got := NewRawMessage(tt.bytes)

			assert.Equal(t, tt.rawMessage, got)
		})
	}
}

func BenchmarkNewRawMessage(b *testing.B) {
	for _, tt := range rawMessageTestCases {
		b.Run(tt.name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				_ = NewRawMessage(tt.bytes)
			}
		})
	}
}

func TestRawMessage_Bytes(t *testing.T) {
	for _, tt := range rawMessageTestCases {
		if tt.bytes == nil {
			continue
		}
		t.Run(tt.name, func(t *testing.T) {
			got := tt.rawMessage.Bytes()

			assert.Equal(t, tt.bytes, got)
		})
	}
}

func BenchmarkRawMessage_Bytes(b *testing.B) {
	for _, tt := range rawMessageTestCases {
		if tt.bytes == nil {
			continue
		}
		b.Run(tt.name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				_ = tt.rawMessage.Bytes()
			}
		})
	}
}

func TestRawMessage_String(t *testing.T) {
	for _, tt := range rawMessageTestCases {
		if tt.bytes == nil {
			continue
		}
		t.Run(tt.name, func(t *testing.T) {
			got := tt.rawMessage.String()

			assert.Equal(t, string(tt.bytes), got)
		})
	}
}

func BenchmarkRawMessage_String(b *testing.B) {
	for _, tt := range rawMessageTestCases {
		if tt.bytes == nil {
			continue
		}
		b.Run(tt.name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				_ = tt.rawMessage.String()
			}
		})
	}
}
