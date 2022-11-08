package instruct

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tikafog/of/buffers/instruct"
)

func TestNewInstruct(t *testing.T) {
	type args struct {
		p    Type
		data any
	}
	tests := []struct {
		name string
		args args
		want any
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				p: instruct.TypeResource,
				data: &ResourceData{
					List: []string{
						"a", "b", "c",
					},
				},
			},
			want: nil,
		},
		{
			name: "",
			args: args{
				p: instruct.TypeResource,
				data: &ReportData{
					Version: 5,
					Type:    0,
					Last:    5,
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.args.data.(type) {
			case *ResourceData:
				got := NewInstruct[ResourceData]()
				got.SetDataSource(tt.args.data.(*ResourceData))
				assert.Equal(t, got.Type(), instruct.TypeResource)
				ri, err := ParseJSONInstruct(got.JSON())
				if err != nil {
					t.Fatal(err)
					return
				}
				t.Logf("Decode1: %+v", string(ri.Data()))
				i, err := ParseJSONInstruct(got.JSON())
				if err != nil {
					t.Fatal(err)
					return
				}
				ri3, ok := CastInstruct[ResourceData](i)
				if ok {
					t.Logf("Decode2: %+v", string(ri3.Data()))
				}
				ri2, err := ParseJSONInstruct(i.JSON())
				if err != nil {
					t.Fatal(err)
				}
				t.Logf("Decode3: %+v", string(ri2.Data()))

				var inst metaInstruct
				_ = json.Unmarshal(i.JSON(), &inst)
				buf := inst.Bytes()

				c, err := ParseInstruct(buf)
				if err != nil {
					t.Fatal(c)
				}

				ri4, ok := CastInstruct[ResourceData](c)
				if !ok {
					p := reflect.TypeOf(c)
					t.Logf("Decode4: Type: (%v)", p)
				}
				//ri3, _ := ParseInstruct[ResourceData](c)
				t.Logf("Decode4: %+v", string(ri4.Data()))
			case *ReportData:
				got := NewInstruct[ReportData]()
				got.SetDataSource(tt.args.data.(*ReportData))
				assert.Equal(t, got.Type(), instruct.TypeReport)
				ri, err := ParseJSONInstruct(got.JSON())
				if err != nil {
					t.Fatal(err)
					return
				}

				t.Logf("Decode1: %+v", string(ri.Data()))
				i, err := ParseJSONInstruct(got.JSON())
				if err != nil {
					t.Fatal(err)
					return
				}
				ri3, ok := CastInstruct[ReportData](i)
				if ok {
					t.Logf("Decode2: %+v", string(ri3.Data()))
				}
				ri2, err := ParseJSONInstruct(i.JSON())
				if err != nil {
					t.Fatal(err)
					return
				}
				t.Logf("Decode3: %+v", string(ri2.Data()))

				var inst metaInstruct
				_ = json.Unmarshal(i.JSON(), &inst)
				buf := inst.Bytes()

				c, err := ParseInstruct(buf)
				if err != nil {
					t.Fatal(c)
				}

				ri4, ok := CastInstruct[ReportData](c)
				if !ok {
					p := reflect.TypeOf(c)
					t.Logf("Decode4: Type: (%v)", p)
				}
				//ri3, _ := ParseInstruct[ResourceData](c)
				t.Logf("Decode4: %+v", string(ri4.Data()))
			}

		})
	}
}

//
//func TestNewInstruct1(t *testing.T) {
//	tests := []struct {
//		name string
//		want *Instruct[CorrectData]
//	}{
//		{
//			name: "",
//			want: NewInstruct[CorrectData](),
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got := NewInstruct[CorrectData]()
//			assert.Equal(t, got.Type, tt.want.Type)
//		})
//	}
//}
