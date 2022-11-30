package of

import (
	"reflect"
	"testing"
)

func TestNewKeyArg(t *testing.T) {
	type args struct {
		arg any
	}
	tests := []struct {
		name string
		args args
		want KeyArg
	}{
		{
			name: "",
			args: args{
				struct {
					v string
				}{v: "abc"},
			},
			want: KeyArg{
				reflect.TypeOf(struct {
					v string
				}{v: "abc"}),
				struct {
					v string
				}{v: "abc"},
			},
		},
		{
			name: "",
			args: args{
				string("abc"),
			},
			want: KeyArg{
				Key:   reflect.TypeOf(string("")),
				Value: "abc",
			},
		},
		{
			name: "",
			args: args{
				[]byte("abc"),
			},
			want: KeyArg{
				Key:   reflect.TypeOf([]byte("abc")),
				Value: []byte("abc"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTypeArg(tt.args.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewKeyArg() = %v, want %v", got, tt.want)
			}
		})
	}
}
