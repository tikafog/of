package errors

import (
	"testing"
)

func TestMakeErrIndex(t *testing.T) {
	type args struct {
		prefix uint32
		index  uint32
	}
	tests := []struct {
		name string
		args args
		want ErrIndex
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				prefix: ErrIndexCorePrefix,
				index:  1,
			},
			want: ErrFatherNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakeErrIndex(tt.args.prefix, tt.args.index); got != tt.want {
				t.Errorf("MakeErrIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
