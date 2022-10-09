package instruct

import (
	"reflect"
	"testing"
)

func TestReportData_JSON(t *testing.T) {
	type fields struct {
		Version int
		Type    ReportType
		Last    int64
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		{
			name: "",
			fields: fields{
				Version: 2,
				Type:    3,
				Last:    4,
			},
			want: []byte(`{"version":2,"type":3,"last":4}`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := ReportData{
				Version: tt.fields.Version,
				Type:    tt.fields.Type,
				Last:    tt.fields.Last,
			}
			if got := d.JSON(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JSON() = %v, want %v", string(got), tt.want)
			}
		})
	}
}
