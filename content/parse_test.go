package content

import (
	"reflect"
	"testing"
)

func TestParseContent(t *testing.T) {
	type args struct {
		bytes []byte
	}
	tests := []struct {
		name     string
		args     args
		wantRetC *Content
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				bytes: NewContent(0).JSON(),
			},
			wantRetC: nil,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRetC, err := ParseContent(tt.args.bytes)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseContent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRetC, tt.wantRetC) {
				t.Errorf("ParseContent() gotRetC = %v, want %v", gotRetC, tt.wantRetC)
			}
		})
	}
}
