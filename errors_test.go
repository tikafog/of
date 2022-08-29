package of

import "testing"

func TestStackError(t *testing.T) {
	type args struct {
		err string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    string
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				err: "test error",
			},
			wantErr: true,
			want:    "test error",
		},
		{
			name: "",
			args: args{
				err: "test error",
			},
			wantErr: true,
			want:    "test error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := StackError(tt.args.err)
			if (err != nil) != tt.wantErr {
				t.Errorf("StackError() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !ErrorMessageIs(tt.want, err) {
				t.Errorf("StackError() error = %v, wantErr %v", err, tt.want)
			}
		})
	}
}

func TestErrorMessageIs(t *testing.T) {
	type args struct {
		str string
		err error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				str: "test1",
				err: WrapError(nil, "test1"),
			},
			want: true,
		},
		{
			name: "",
			args: args{
				str: "test1",
				err: WrapError(nil, "test2"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ErrorMessageIs(tt.args.str, tt.args.err); got != tt.want {
				t.Errorf("ErrorMessageIs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrorIndexIs(t *testing.T) {
	type args struct {
		i   Err
		err error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				i:   1,
				err: WrapIndexError(nil, 1),
			},
			want: true,
		},
		{
			name: "",
			args: args{
				i:   1,
				err: WrapIndexError(nil, 2),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ErrorIndexIs(tt.args.i, tt.args.err); got != tt.want {
				t.Errorf("ErrorIndexIs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrorIs(t *testing.T) {
	type args struct {
		err    error
		target error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				err:    WrapIndexError(nil, 0),
				target: WrapError(nil, "core is nil"),
			},
			want: true,
		},
		{
			name: "",
			args: args{
				err:    WrapIndexError(nil, 1),
				target: WrapError(nil, "core is nil"),
			},
			want: false,
		},
		{
			name: "",
			args: args{
				err:    StackError("core is nil"),
				target: WrapError(nil, "core is nil"),
			},
			want: true,
		},
		{
			name: "",
			args: args{
				err:    StackError("core is nil1"),
				target: WrapError(nil, "core is nil"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ErrorIs(tt.args.err, tt.args.target); got != tt.want {
				t.Errorf("ErrorIs() = %v, want %v", got, tt.want)
			}
		})
	}
}
