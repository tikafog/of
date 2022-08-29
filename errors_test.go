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
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				err: "test error",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := StackError(tt.args.err); (err != nil) != tt.wantErr {
				t.Errorf("StackError() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
