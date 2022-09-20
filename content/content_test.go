package content

import (
	"reflect"
	"testing"

	"github.com/tikafog/of/buffers/content"
)

func TestNewContentWithType(t *testing.T) {
	type args struct {
		tp content.Type
	}
	tests := []struct {
		name string
		args args
		want *Content
	}{
		{
			name: "",
			args: args{
				tp: 0,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewContentWithType(tt.args.tp); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewContentWithType() = %v, want %v", got, tt.want)
			}
		})
	}
}
