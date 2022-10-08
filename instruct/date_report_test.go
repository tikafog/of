package instruct

import (
	"reflect"
	"testing"
)

func TestReportData_JSON(t *testing.T) {
	type fields struct {
		VersionData DataVersion
		Type        ReportType
		Last        int64
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		{
			name: "",
			fields: fields{
				VersionData: DataVersion{
					Version: 2,
				},
				Type: 3,
				Last: 4,
			},
			want: []byte(`{"version":2,"type":3,"last":4}`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := ReportData{
				DataVersion: tt.fields.VersionData,
				Type:        tt.fields.Type,
				Last:        tt.fields.Last,
			}
			if got := d.JSON(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JSON() = %v, want %v", string(got), tt.want)
			}
		})
	}
}
