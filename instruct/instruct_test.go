package instruct

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tikafog/of/buffers/instruct"
	"github.com/tikafog/of/utils"
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
				got.SetData(tt.args.data.(*ResourceData))
				assert.Equal(t, got.Type, instruct.TypeResource)
				var ri Instruct[ResourceData]
				err := json.Unmarshal(got.JSON(), &ri)
				if err != nil {
					t.Fatal(err)
					return
				}
				t.Logf("Decode1: %+v", ri.Data)
				_, i, err := ParseJSONInstruct(got.JSON())
				if err != nil {
					t.Fatal(err)
					return
				}
				switch v := i.(type) {
				case *Instruct[ResourceData]:
					t.Logf("Decode2: %+v", v.Data)
				default:
					t.Fatal("type", reflect.TypeOf(i))
				}
				data := utils.Must(json.Marshal(i))

				var ri2 Instruct[ResourceData]
				_ = json.Unmarshal(data, &ri2)
				t.Logf("Decode3: %+v", ri2.Data)

				var inst metaInstruct
				_ = json.Unmarshal(data, &inst)
				buf := inst.Bytes()

				_, c, err := ParseInstruct(buf)
				if err != nil {
					t.Fatal(c)
				}

				ri3, ok := (c).(*Instruct[ResourceData])
				if !ok {
					p := reflect.TypeOf(c)
					t.Logf("Decode4: Type: (%v)", p)
				}
				//ri3, _ := ParseInstruct[ResourceData](c)
				t.Logf("Decode4: %+v", ri3.Data)
			case *ReportData:
				got := NewInstruct[ReportData]()
				got.SetData(tt.args.data.(*ReportData))
				assert.Equal(t, got.Type, instruct.TypeReport)
				var ri Instruct[ReportData]
				err := json.Unmarshal(got.JSON(), &ri)
				if err != nil {
					t.Fatal(err)
					return
				}

				t.Logf("Decode1: %+v", ri.Data)
				_, i, err := ParseJSONInstruct(got.JSON())
				if err != nil {
					t.Fatal(err)
					return
				}
				switch v := i.(type) {
				case *Instruct[ReportData]:
					t.Logf("Decode2: %+v", v.Data)
				default:
					t.Fatal("type", reflect.TypeOf(i))
				}
				data := utils.Must(json.Marshal(i))

				var ri2 Instruct[ReportData]
				_ = json.Unmarshal(data, &ri2)
				t.Logf("Decode3: %+v", ri2.Data)

				var inst metaInstruct
				_ = json.Unmarshal(data, &inst)
				buf := inst.Bytes()

				_, c, err := ParseInstruct(buf)
				if err != nil {
					t.Fatal(c)
				}

				ri3, ok := (c).(*Instruct[ReportData])
				if !ok {
					p := reflect.TypeOf(c)
					t.Logf("Decode4: Type: (%v)", p)
				}
				//ri3, _ := ParseInstruct[ResourceData](c)
				t.Logf("Decode4: %+v", ri3.Data)
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
