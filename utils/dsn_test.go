package utils

import (
	"testing"
)

func TestOpenDSN(t *testing.T) {
	type args struct {
		t    DSNType
		path string
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   bool
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				t:    0,
				path: "/mnt/d/test",
				name: "dsnfile",
			},
			want:    "file:\\mnt\\d\\test\\dsnfile_warp?cache=shared&_journal=WAL&_fk=1",
			want1:   false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := OpenDSN(tt.args.t, tt.args.path, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("OpenDSN() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("OpenDSN() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("OpenDSN() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
