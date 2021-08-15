package conventionalcommit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMessage_IsBreakingChange(t *testing.T) {
	type fields struct {
		Breaking        bool
		BreakingChanges []string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "false breaking flag, no change texts",
			fields: fields{
				Breaking:        false,
				BreakingChanges: []string{},
			},
			want: false,
		},
		{
			name: "true breaking flag, no change texts",
			fields: fields{
				Breaking:        true,
				BreakingChanges: []string{},
			},
			want: true,
		},
		{
			name: "false breaking flag, 1 change texts",
			fields: fields{
				Breaking:        false,
				BreakingChanges: []string{"be careful"},
			},
			want: true,
		},
		{
			name: "true breaking flag, 1 change texts",
			fields: fields{
				Breaking:        true,
				BreakingChanges: []string{"be careful"},
			},
			want: true,
		},
		{
			name: "false breaking flag, 3 change texts",
			fields: fields{
				Breaking:        false,
				BreakingChanges: []string{"be careful", "oops", "ouch"},
			},
			want: true,
		},
		{
			name: "true breaking flag, 3 change texts",
			fields: fields{
				Breaking:        true,
				BreakingChanges: []string{"be careful", "oops", "ouch"},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := &Message{
				Breaking:        tt.fields.Breaking,
				BreakingChanges: tt.fields.BreakingChanges,
			}

			got := msg.IsBreakingChange()

			assert.Equal(t, tt.want, got)
		})
	}
}
