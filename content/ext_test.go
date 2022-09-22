package content

import (
	"reflect"
	"testing"

	"github.com/tikafog/of/buffers/content"

	"github.com/bxcodec/faker/v4"
)

func randomExtNode() ExtNode {
	var ext ExtNode
	err := faker.FakeData(&ext)
	checkErr(err)
	//makeExt, errors := MakeExt(&ext)
	//checkErr(errors)
	//return makeExt
	return ext
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func TestMakeExt(t *testing.T) {
	type args struct {
		v ExtNode
	}
	tests := []struct {
		name    string
		args    args
		want    ExtNode
		wantErr bool
	}{
		{
			name: "",
			args: args{
				v: randomExtNode(),
			},
			//want:    ,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MakeExt(&tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("MakeExt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			tt.want = tt.args.v
			c := NewContentWithType(content.TypeCore)
			c.SetExts(got).SetMessage(&Message{})
			json, err := c.JSON()
			checkErr(err)

			parseC, err := ParseJSONContent(json)
			checkErr(err)
			var extNode ExtNode
			for _, ext := range parseC.Exts {
				err := ParseExt(ext, &extNode)
				checkErr(err)
				if !reflect.DeepEqual(extNode, tt.want) {
					t.Errorf("MakeExt1() got = %+v, want %+v", extNode, tt.want)
				}
			}
			bytes := c.FinishBytes()
			retC, err := ParseContent(bytes)
			if err != nil {
				t.Errorf("%+v\n", err)
			}
			for _, ext := range retC.Exts {
				err := ParseExt(ext, &extNode)
				checkErr(err)
				if !reflect.DeepEqual(extNode, tt.want) {
					t.Errorf("MakeExt2() got = %+v, want %+v", extNode, tt.want)
				}
			}

			t.Logf("%+v\n", extNode)
		})
	}
}
