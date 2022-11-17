package content

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/tikafog/of/buffers/content"
	"github.com/tikafog/of/utils"
)

func TestNewContentWithType(t *testing.T) {
	type args struct {
		tp  content.Type
		msg *Message
	}
	nodata := utils.Must(json.Marshal("nodata"))
	fmt.Println("nodata is:", string(nodata))
	tests := []struct {
		name string
		args args
		want *Content
	}{
		{
			name: "",
			args: args{
				tp:  0,
				msg: NewMessageV2([]byte("nodata")),
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewContentWithType(tt.args.tp).SetMessage(tt.args.msg)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewContentWithType() = %v, want %v", got, tt.want)
			}
			t.Logf("jsonv:%+v", string(got.JSON()))
			oldc := []byte(`{"version":"0.0.1","message":{"version":1,"length":6,"data":"bm9kYXRh"}}`)
			jsonContent1, err := ParseJSONContent(oldc)
			if err != nil {
				t.Fatal(err)
				return
			}
			//var bytes []byte
			//_ = json.Unmarshal(jsonContent.Message.Data, &bytes)
			//t.Logf("json1:%+v", string(jsonContent1.JSON()))
			if jsonContent1.Message != nil {
				t.Logf("data:%+v", string(jsonContent1.Message.Data))
			}
			//t.Log(string(utils.Must(jsonContent.JSON())))

			oldc2 := []byte(`{"version":"0.0.1","message":{"version":2,"length":8,"data":"nodata"}}`)
			jsonContent2, err2 := ParseJSONContent(oldc2)
			if err2 != nil {
				t.Fatal(err2)
				return
			}
			t.Logf("json2:%+v", string(jsonContent2.JSON()))
			if jsonContent2.Message != nil {
				t.Logf("data:%+v", string(jsonContent2.Message.Data))
			}
			oldc3 := []byte(`{"version":"0.0.1","message":{"version":1,"length":8,"data":"Im5vZGF0YSI="}}`)
			jsonContent3, err3 := ParseJSONContent(oldc3)
			if err3 != nil {
				t.Fatal(err3)
				return
			}
			t.Logf("json3:%+v", string(jsonContent3.JSON()))
			if jsonContent3.Message != nil {
				t.Logf("data:%+v", string(jsonContent3.Message.Data))
			}

			oldc4 := []byte(`{"version":"0.0.1","type":6,"message":{"last":1655802406522860007,"version":1,"length":16798,"data":"W3siaWQiOiI0YzhjNzE0OS1iZjZkLTRmMjItODE4Ni0xN2E2MTdkNmY4NGMiLCJjcmVhdGVkX3VuaXgiOjE2MjUwMjU4ODcsInVwZGF0ZWRfdW5peCI6MTYyNTAyNjY0NSwib3MiOiJhbmRyb2lkIiwiYXJjaCI6ImFtZDY0IiwidmVyc2lvbiI6IjAuMC4xIiwicmlkIjoiYmFmeWJlaWRnbzdxaDRybHNlYmc1eGtvZXBoYmwzNHBtaWM3bXVzdmV1dmFlZWFpZmVqNmdqaHl5anEiLCJjcmMzMiI6ImNjYTY3MTQ5IiwiYXR0ciI6ImFwcCIsInRydW5jYXRlIjp0cnVlLCJ0aXRsZSI6IuaJi+acuuWuieijheWMhSIsImRldGFpbCI6IuaJi+acuuWuieijheWMhSIsInNpZ24iOiI2MzcxZjk4NjhlNTljYzljNWI0ZmI3YmU5ZmQ5OTFkYjYwY2Y2MmZhZDE1NzdlY2E3N2QxMDczZTY5NWNhYjBhIiwiZWRnZXMiOnt9fSx7ImlkIjoiYmZiNWJkMjEtN2VkMS00ODQ1LTliNDYtZDNlZTI1NmNjZTI5IiwiY3JlYXRlZF91bml4IjoxNjI1MTExNjU2LCJ1cGRhdGVkX3VuaXgiOjE2MjUxMTQyMzgsIm9zIjoiYW5kcm9pZCIsImFyY2giOiJhbWQ2NCIsInZlcnNpb24iOiIwLjEuMSIsInJpZCI6ImJhZnliZWliZ2pxY2tuMzIyZTRzY2piYjZ1Nm81N3BjYTJoaXlsb2Nxc2trajV4d3N5c3Vlczd1NDNpIiwiY3JjMzIiOiI0MjQ1NDkwNCIsImF0dHIiOiJhcHAiLCJmb3JjaWJseSI6dHJ1ZSwidGl0bGUiOiIwLjEuMSBWZXJzaW9uIHVwZGF0ZSIsImRldGFpbCI6IkltcHJvdmUgdmlkZW8gcGxheWJhY2sgc3BlZWQiLCJzaWduIjoiM2FjNDIyNzE0NmE3Mzc2ZWVjNzgwNDQ3YTMxNDMxMDE5N2RhMjM4MDE4YTMzZTYwYjYxMzI5MjA1YzY2NmVhNyIsImVkZ2VzIjp7fX0seyJpZCI6Ijc1MGY1YTgwLTM5OTAtNGY4ZC04MGMwLTc0NjBiOGQ3MjM3MSIsImNyZWF0ZWRfdW5peCI6MTYyNTExNDMzNCwidXBkYXRlZF91bml4IjoxNjI1MTE0MzU3LCJvcyI6ImFuZHJvaWQiLCJhcmNoIjoiYW1kNjQiLCJ2ZXJzaW9uIjoiMC4xLjMiLCJyaWQiOiJiYWZ5YmVpZDRuZ3R2bXhkcGlia3YzbGp6amdzbXl5eWZoc2RrYWZ5ZDdpdHV3ZzVkZW43dWFvemFqNCIsImNyYzMyIjoiMDBjOWExYjQiLCJhdHRyIjoiYXBwIiwiZm9yY2libHkiOnRydWUsInRpdGxlIjoiMC4xLjMgVmVyc2lvbiB1cGRhdGUiLCJkZXRhaWwiOiJWZXJzaW9uIHVwZGF0ZSIsInNpZ24iOiJkNDFlNzZjNGU5YmU5MTAzMjI2MTMzOGIwMzNlZjE4Y2FkMmIwMTFjNGM5MGQzNDc0NmUzMmMwNjQ0YWYxMzE2IiwiZWRnZXMiOnt9fSx7ImlkIjoiZWU0YWUyYTMtMzU5ZC00MWQzLTlhNTMtYWI5NmQyNDBmOWQxIiwiY3JlYXRlZF91bml4IjoxNjI1MTIwMDEwLCJ1cGRhdGVkX3VuaXgiOjE2MjUxMjAzNTksIm9zIjoiYW5kcm9pZCIsImFyY2giOiJhbWQ2NCIsInZlcnNpb24iOiIwLjEuMyIsInJpZCI6ImJhZnliZWloYTVieTZyN3JtaDViZ3piaDJ0Y2xkbnM1NnJjY3RrZWpicmI2N2xueWZhaGJlN2l6YmJxIiwiY3JjMzIiOiI4NWVmOWM1YSIsImF0dHIiOiJib3giLCJmb3JjaWJseSI6dHJ1ZSwidGl0bGUiOiIwLjEuMyBUViB2ZXJzaW9uIiwiZGV0YWlsIjoiIFRWIHZlcnNpb24iLCJzaWduIjoiMTc1Y2U5NzgwYjc0M2FhOWU2NzM3YTQyYmM0NWFhYjg3YWVhMzg5ODI1MDA3YzcyZjdhYTdkMTMyMjIzNjZhNSIsImVkZ2VzIjp7fX0seyJpZCI6Ijk5MDQ5YzllLWEyNmYtNDI3OC1iY2YzLWI5ZmYwMGEzY2NlMyIsImNyZWF0ZWRfdW5peCI6MTYyNTIxNjk5MywidXBkYXRlZF91bml4IjoxNjI1MjE3MjYxLCJvcyI6ImFuZHJvaWQiLCJhcmNoIjoiYW1kNjQiLCJ2ZXJzaW9uIjoiMC4xLjYiLCJyaWQiOiJiYWZ5YmVpaGQ2aG10aXV2ZWUzNG9zaWdtcGthNjR1cXIybXA3ZDU1Z2IyeXd0aXluYXB1aW5rbmNpaSIsImNyYzMyIjoiMzFkOGRiZDUiLCJhdHRyIjoiYXBwIiwiZm9yY2libHkiOnRydWUsInRpdGxlIjoiMC4xLjYgdmVyc2lvbiB1cGRhdGUiLCJkZXRhaWwiOiJ2ZXJzaW9uIHVwZGF0ZSIsInNpZ24iOiI0YWM1YjRlM2Y1ZTlhMTZjZGU5OWU0ZWFlMDk3NDJkNGMxMmMwYmJhOThiOWNhOTdiYmY4MDZmZmRlMGUyODA4IiwiZWRnZXMiOnt9fSx7ImlkIjoiNDhjNzgyZmEtMmEwYy00ZjNiLThhMGQtY2I1ZGIxMTZmYjdkIiwiY3JlYXRlZF91bml4IjoxNjI1MjE3MTM5LCJ1cGRhdGVkX3VuaXgiOjE2MjUyMTcyNjksIm9zIjoiYW5kcm9pZCIsImFyY2giOiJhbWQ2NCIsInZlcnNpb24iOiIwLjEuNiIsInJpZCI6ImJhZnliZWlkYTJlNm9ua3hhZWFnNGZieW5oaHltZG1uN3l2ajNjbzVnY3h0Yno0aTJvdjRla291dXBxIiwiY3JjMzIiOiI5NDA5OWRhYyIsImF0dHIiOiJib3giLCJmb3JjaWJseSI6dHJ1ZSwidGl0bGUiOiIwLjEuNiB2ZXJzaW9uIHVwZGF0ZSIsImRldGFpbCI6InZlcnNpb24gdXBkYXRlIiwic2lnbiI6IjA4ZmM5NzZkZjE0MDBlNjk1MDdhYTkwNjgxNWMwNDY2YzUyZDY2NTQ3YTBjMDNmOTY3ODMxZGM4MThjZTRiMGUiLCJlZGdlcyI6e319LHsiaWQiOiIwMTBkODNkMC1mMGU0LTQwZTktOGUzZS1kNDA5N2YxYmNjNWMiLCJjcmVhdGVkX3VuaXgiOjE2MjYwODQwMzcsInVwZGF0ZWRfdW5peCI6MTYyNjA4NDExNCwib3MiOiJhbmRyb2lkIiwiYXJjaCI6ImFtZDY0IiwidmVyc2lvbiI6IjAuMS44LjE4IiwicmlkIjoiYmFmeWJlaWJldjN0a2JlNXF4cXR2amh5dzdxbDZiMmd2eWpsMzd0Y2dldjU3cnNzZHFybXVlc3lneHUiLCJjcmMzMiI6IjM5ODMxNTNiIiwiYXR0ciI6ImJveCIsImZvcmNpYmx5Ijp0cnVlLCJ0aXRsZSI6IjAuMS44LjE4IiwiZGV0YWlsIjoiMC4xLjguMTgiLCJzaWduIjoiMzFmMzc2ZGMzMGYwYTA4YTZiYjEzNDAyNGQ0YjcyOWJhNzJjYTRiY2U5NzJlNTZmZWU2YTQyNTE4NWE0NWZhOSIsImVkZ2VzIjp7fX0seyJpZCI6ImNkODljOGVlLTQzZDEtNGQwMi04MGUwLTRmZGRjMmU5ZDYyMyIsImNyZWF0ZWRfdW5peCI6MTYyNjE1MDMyNSwidXBkYXRlZF91bml4IjoxNjI2MTUwOTg0LCJvcyI6ImFuZHJvaWQiLCJhcmNoIjoiYW1kNjQiLCJ2ZXJzaW9uIjoiMC4xLjkuMSIsInJpZCI6ImJhZnliZWlkZW9pbjRpc3hlNG1kNHBxemRsNnRtd2t4bnY1bGpka2Z1Z2tpdnpmNnFjNjJtdTVldGNxIiwiY3JjMzIiOiIzYzc0MDllNSIsImF0dHIiOiJhcHAiLCJmb3JjaWJseSI6dHJ1ZSwidGl0bGUiOiIwLjEuOS4xIiwiZGV0YWlsIjoiMC4xLjkuMSIsInNpZ24iOiIyNGRmOWFkZDQ5Mjg2ODY2NGNlNTdjMmNhNDMzMjNmMmRjYzQ0ZmFjYjg1NmViZGI0ODVhYjFmMmVjNWMxYWI2IiwiZWRnZXMiOnt9fSx7ImlkIjoiNzliZmRjZmQtMGMwNy00OGI4LWIyYzgtZmQ1ZGMzNzFiOTJkIiwiY3JlYXRlZF91bml4IjoxNjI2MTUzNTQ3LCJ1cGRhdGVkX3VuaXgiOjE2MjYxNTM4NDcsIm9zIjoiYW5kcm9pZCIsImFyY2giOiJhbWQ2NCIsInZlcnNpb24iOiIwLjEuOSIsInJpZCI6ImJhZnliZWlnZmJ5Z2Fqc3ZuNmx5N3R4N2Q3cWRheGprcGpndmVndmJrdnV3bWhlNzJxcjdjY2NjNGE0IiwiY3JjMzIiOiJlYmFhOWUyNiIsImF0dHIiOiJhcHAiLCJmb3JjaWJseSI6dHJ1ZSwidGl0bGUiOiIwLjEuOS4yIiwiZGV0YWlsIjoiMC4xLjkuMiIsInNpZ24iOiJmMmM0MWMwYzhiNTJiNTExMTBhMmIzZTlmNDhiYzRhYTU3ZmY1ZmE3NWIyYWE3MGEyZDc3NWY1ZjIxN2JlYjFkIiwiZWRnZXMiOnt9fSx7ImlkIjoiM2RlOTdhNjMtYWM5OS00Y2ZiLTk1MjctMTE4MjdjZDZiODg1IiwiY3JlYXRlZF91bml4IjoxNjI2MTUzODc5LCJ1cGRhdGVkX3VuaXgiOjE2MjYxNTYxNTcsIm9zIjoiYW5kcm9pZCIsImFyY2giOiJhbWQ2NCIsInZlcnNpb24iOiIwLjEuOSIsInJpZCI6ImJhZnliZWloYWFscTR2ZGlubWE0dzN5ajNhYnNicWJkdGl5aG9oeHlra3F5YjZ3eWhienByamR2czRxIiwiY3JjMzIiOiJjNzI2NDdlMCIsImF0dHIiOiJib3giLCJmb3JjaWJseSI6dHJ1ZSwidGl0bGUiOiIwLjEuOS4xIiwiZGV0YWlsIjoiMC4xLjkuMSIsInNpZ24iOiI4ZjQ2MWUyYzJmMjliMzJmMDBmYWM4N2Q5ZTIzOGIyMzhkNTA3NGEyZDQ0OWM3NTVmNmI3OTFkYTlhZTVmNmU3IiwiZWRnZXMiOnt9fSx7ImlkIjoiYzI0ZmFmYTMtZjQwYy00NGM1LWE1MzYtNmQ0NGM3MTQwNmExIiwiY3JlYXRlZF91bml4IjoxNjI2MTU4NzAyLCJ1cGRhdGVkX3VuaXgiOjE2MjYxNTk2ODcsIm9zIjoiYW5kcm9pZCIsImFyY2giOiJhbWQ2NCIsInZlcnNpb24iOiIwLjEuOSIsInJpZCI6ImJhZnliZWllbGNtZ3Z0ejdqcGo1aWJsYWRpNGF5eTczZWh4djdjMjRreWI3ZTVyeHV0YWozNnRjNWxhIiwiY3JjMzIiOiI3Yjc1NDE2NyIsImF0dHIiOiJib3giLCJmb3JjaWJseSI6dHJ1ZSwidGl0bGUiOiIwLjEuOS4yMSIsImRldGFpbCI6IjAuMS45LjIxIiwic2lnbiI6IjczOTgyZWZjY2FlZjdkMTk4NGNlNTI2NGNmZjI5YmFmZTU4ZDJjMWM1MjUwMWZiOGFkZThjZGE5YzJkNTQwODMiLCJlZGdlcyI6e319LHsiaWQiOiIxYjBiMjRkYy0zMTc0LTQ4ZDUtODAxMy1lNDFlYzRiNjgyZWIiLCJjcmVhdGVkX3VuaXgiOjE2MjYxNTk3NTMsInVwZGF0ZWRfdW5peCI6MTYyNjE1OTkzOSwib3MiOiJhbmRyb2lkIiwiYXJjaCI6ImFtZDY0IiwidmVyc2lvbiI6IjAuMS45IiwicmlkIjoiYmFmeWJlaWRsMnZrb2llZTNwcDRkeXF6ZmJ6ZmZ3Mmt1MnR6c2djcnFzamNsb2I2cmFlNHRkdmY1aHUiLCJjcmMzMiI6ImE2ZGM3YWFkIiwiYXR0ciI6ImFwcCIsImZvcmNpYmx5Ijp0cnVlLCJ0aXRsZSI6ImFwcC1yZWxlYXNlLjAuMS45LjciLCJkZXRhaWwiOiJhcHAtcmVsZWFzZS4wLjEuOS43Iiwic2lnbiI6ImQ4Nzk3ZDUwNzA2Zjc2M2RiNDRlZDJlMjc2ZmUwNTE4ZjYzYmY5OTk1MmVmYTNhM2VhMmUwZWE5NDNiZWE5ZjIiLCJlZGdlcyI6e319LHsiaWQiOiJkYjliNTliMi1jNzA5LTQ2OGUtODNlYi1iM2NmZDlhNWNlN2YiLCJjcmVhdGVkX3VuaXgiOjE2MjYxNjI2ODksInVwZGF0ZWRfdW5peCI6MTYyNjE2MzIwMywib3MiOiJhbmRyb2lkIiwiYXJjaCI6ImFtZDY0IiwidmVyc2lvbiI6IjAuMS45IiwicmlkIjoiYmFmeWJlaWhyZ2lzaXB3b2NxM2Fqa2JxdHBwM3FrbGtpcmZvbGlzamxrZm43cnhzdXF0cjdxYXo1c3UiLCJjcmMzMiI6ImQ2MTU3MjYyIiwiYXR0ciI6ImJveCIsImZvcmNpYmx5Ijp0cnVlLCJ0aXRsZSI6InR2LXJlbGVhc2UuMC4xLjkuMjIiLCJkZXRhaWwiOiJ0di1yZWxlYXNlLjAuMS45LjIyIiwic2lnbiI6ImRlOTQ2NjJjYTAzMjY4NDA1NWU2YjBlYWE2MjlmY2JmNTE0ZjU0NzdmNGE2MTZhOTRiOWRiOTdlNTgyMmFlOWUiLCJlZGdlcyI6e319LHsiaWQiOiJiOTM0ZTcyMS00ZDdmLTQ3NTgtOGM0NS1hMTU0Y2U1NDhkODUiLCJjcmVhdGVkX3VuaXgiOjE2MjYyMzAxNTEsInVwZGF0ZWRfdW5peCI6MTYyNjIzMDM1MCwib3MiOiJhbmRyb2lkIiwiYXJjaCI6ImFtZDY0IiwidmVyc2lvbiI6IjAuMS4xMCIsInJpZCI6ImJhZnliZWlmcnR0cnE3a3hxYWNuenozMnFqeWdyZ2pkZmh5YzNsbGV5emU0djUyc24zeXVweW81dm9pIiwiY3JjMzIiOiIzNGYxNDI4YSIsImF0dHIiOiJib3giLCJmb3JjaWJseSI6dHJ1ZSwidGl0bGUiOiJ0di1yZWxlYXNlLjAuMS4xMC4yMyIsImRldGFpbCI6InR2LXJlbGVhc2UuMC4xLjEwLjIzIiwic2lnbiI6IjZiNWE5ZWI4Nzg5ZmZhMTdlY2MyZGQ4YTY0Y2I3Zjk2N2Y5ODNmM2UwODNiNzQwZTIwMGUwYzcyMzVlYjc1N2EiLCJlZGdlcyI6e319LHsiaWQiOiI2NDBiYjQ0YS03YzAxLTQ5NTctYWRjYi1jYTUyOTg0MTgyYTMiLCJjcmVhdGVkX3VuaXgiOjE2MjYyMzA0MTYsInVwZGF0ZWRfdW5peCI6MTYyNjIzMDk1MSwib3MiOiJhbmRyb2lkIiwiYXJjaCI6ImFtZDY0IiwidmVyc2lvbiI6IjAuMS4xMCIsInJpZCI6ImJhZnliZWllY3RnbnA1NWRxMm5mbmhtbjZweXNxNmI0cjJyeDNvcG9xbjNoc2hnM2tlcDdpcGZ0eHFxIiwiY3JjMzIiOiJjNTY4NzIyNCIsImF0dHIiOiJhcHAiLCJmb3JjaWJseSI6dHJ1ZSwidGl0bGUiOiJhcHAtcmVsZWFzZS4wLjEuMTAuMTAiLCJkZXRhaWwiOiJhcHAtcmVsZWFzZS4wLjEuMTAuMTAiLCJzaWduIjoiMDVhMWJiN2UzNzExNGZkZTI2MWYwNDIzOTExMWQ5NjBlODlkNGU5ZjA1NDYwMzcyMzQ4NDJiYzY0NjZmNmQ0ZCIsImVkZ2VzIjp7fX0seyJpZCI6ImE4MzJlZjlkLTFjZjAtNGJmYy05N2U4LTEyMzMyMDMzODgwMCIsImNyZWF0ZWRfdW5peCI6MTYyNjIzOTM1NCwidXBkYXRlZF91bml4IjoxNjI2MjM5NTAwLCJvcyI6ImFuZHJvaWQiLCJhcmNoIjoiYW1kNjQiLCJ2ZXJzaW9uIjoiMC4xLjExIiwicmlkIjoiYmFmeWJlaWZmN3VoeWVoM3IydzJyeHA3amxsbmU3emtocHZqY3R5aXg1a2JsbjZxd3B0a2U3dmt3NnkiLCJjcmMzMiI6IjMxZmY4NmZjIiwiYXR0ciI6ImFwcCIsImZvcmNpYmx5Ijp0cnVlLCJ0aXRsZSI6ImFwcC1yZWxlYXNlLjAuMS4xMS4xMSIsImRldGFpbCI6ImFwcC1yZWxlYXNlLjAuMS4xMS4xMSIsInNpZ24iOiI5MGJhNDFhMjY3YTY4YTJlMGMzMjBmYzVkZTU1YmQ5ODEyNGE2NDlkNmFiMDFjZDFlOTQzOGVlYWNhZDk2NWM0IiwiZWRnZXMiOnt9fSx7ImlkIjoiZDYxOGE2YzEtNWY5YS00ZTczLTllYWEtNTMyNTg1OWNmZWMzIiwiY3JlYXRlZF91bml4IjoxNjI2ODYyNjI3LCJ1cGRhdGVkX3VuaXgiOjE2MjY4NjI3MzgsIm9zIjoiYW5kcm9pZCIsImFyY2giOiJhbWQ2NCIsInZlcnNpb24iOiIwLjEuMTIiLCJyaWQiOiJiYWZ5YmVpY3lvd2J1aDNma3dvdGF4M29pa3c1YWZrM3ZlYnE1b3ljc21wdm91cXd4aWljcGNlZzRqbSIsImNyYzMyIjoiNzA0ZmI1NmUiLCJhdHRyIjoiYXBwIiwiZm9yY2libHkiOnRydWUsInRpdGxlIjoiYXBwLXJlbGVhc2UwLjEuMTIuMTYiLCJkZXRhaWwiOiJhcHAtcmVsZWFzZTAuMS4xMi4xNiIsInNpZ24iOiI4ZDU4ZGM2Mjc1ZGFiZDI1MjgzMmYyZTFjODE5OWQ2NzAzYmM0OTg3NGJjNDdiM2UwNGJiMDY3NzQ4YWM3OGU1IiwiZWRnZXMiOnt9fSx7ImlkIjoiMTFmMzRkZGYtZTEyNS00OGNiLWI5YzEtNjg5NzgzM2RkYjRkIiwiY3JlYXRlZF91bml4IjoxNjI3MDA5OTg2LCJ1cGRhdGVkX3VuaXgiOjE2MjcwMTE1MjEsIm9zIjoiYW5kcm9pZCIsImFyY2giOiJhbWQ2NCIsInZlcnNpb24iOiIwLjEuMTIiLCJyaWQiOiJiYWZ5YmVpZHZ1bHE3NHd2M2dhNmk2b2I2cnRyeG1lcWVpeDc1ZHBwZm01cm1naWxzam9uZGI3c3c3NCIsImNyYzMyIjoiNDExZDVjNDIiLCJhdHRyIjoiYm94IiwiZm9yY2libHkiOnRydWUsInRpdGxlIjoidHYtcmVsZWFzZS4wLjEuMTIuMjgiLCJkZXRhaWwiOiJ0di1yZWxlYXNlLjAuMS4xMi4yOCIsInNpZ24iOiI2MTZiYWM1NWUwZTQyMWNkZjdmN2Q1OTRjYTA0ZTA2ZDY1M2JjZGZlZGE3MGM3MWMzMTI0ODg2M2FiODM4MjAyIiwiZWRnZXMiOnt9fSx7ImlkIjoiZGEzZTljOWQtOWEwMC00YjU2LTgxN2ItMTFmNTg0NTNmNmRiIiwiY3JlYXRlZF91bml4IjoxNjI3NDUwMTc0LCJ1cGRhdGVkX3VuaXgiOjE2Mjc0NTEwMTMsIm9zIjoiYW5kcm9pZCIsImFyY2giOiJhbWQ2NCIsInZlcnNpb24iOiIwLjEuMTUiLCJyaWQiOiJiYWZ5YmVpaGt1M2RwZnBtd3Vqd3U3a3dsaTJxYnhwMmU3dmRuajR1eG80M2l0NzQ2Z2NjbmhlYzZvaSIsImNyYzMyIjoiZjRkOTJiZTMiLCJhdHRyIjoiYXBwIiwiZm9yY2libHkiOnRydWUsInRpdGxlIjoiYXBwLXJlbGVhc2UuMC4xLjE1LjE3IiwiZGV0YWlsIjoiYXBwLXJlbGVhc2UuMC4xLjE1LjE3Iiwic2lnbiI6IjNlNTg3MmQyNTgzMGUzN2JkZDZkZTRjZWZiNWU2YzVjZTdkNTcxMmQ3MTNmNmIxM2UxZjZmNDFmOTdkOGQ2MmIiLCJlZGdlcyI6e319LHsiaWQiOiI3OWFjMjQ5Ni1jM2JkLTQ4NzQtOTM4ZC05NWZiZTUxMDFhOGYiLCJjcmVhdGVkX3VuaXgiOjE2Mjc0NTEwMzEsInVwZGF0ZWRfdW5peCI6MTYyNzQ1MTE5NCwib3MiOiJhbmRyb2lkIiwiYXJjaCI6ImFtZDY0IiwidmVyc2lvbiI6IjAuMS4xNSIsInJpZCI6ImJhZnliZWlmdjJqazMybnV6NTR1dDZoZGljYWFjaHNnNzN2dGp1ZzJyYzVna215N3M0Zm1tcTNhZW5tIiwiY3JjMzIiOiIwOTBkYTczOSIsImF0dHIiOiJib3giLCJmb3JjaWJseSI6dHJ1ZSwidGl0bGUiOiJ0di1yZWxlYXNlLjAuMS4xNS4zMCIsImRldGFpbCI6InR2LXJlbGVhc2UuMC4xLjE1LjMwIiwic2lnbiI6ImZiN2QzYmEyMGE3ZGE0YzljYWFjMWZhZWYyODc3YzU0ZjIyMDBjNjk4NmM3NTY4Y2M0ZDZkMjkyZGY3ODI2NzYiLCJlZGdlcyI6e319LHsiaWQiOiJlOWU0YjlkMC00MzM2LTQ2NmItYjA0Zi0yY2UzMTI3Mzc1OGEiLCJjcmVhdGVkX3VuaXgiOjE2MjgwNzU3NzQsInVwZGF0ZWRfdW5peCI6MTYyODA3NTk2Niwib3MiOiJhbmRyb2lkIiwiYXJjaCI6ImFtZDY0IiwidmVyc2lvbiI6IjAuMS4xNiIsInJpZCI6ImJhZnliZWlmbmh5emhzYm92Zm5wM2lnMzYzeWtjaWpuamQ1am54a2Jzemc1ZG50N2RmY2M0Y2Nsd2FhIiwiY3JjMzIiOiJmNjY0MjA3MyIsImF0dHIiOiJhcHAiLCJmb3JjaWJseSI6dHJ1ZSwidGl0bGUiOiJhcHAtcmVsZWFzZS4wLjEuMTYuMTkiLCJkZXRhaWwiOiJhcHAtcmVsZWFzZS4wLjEuMTYuMTkiLCJzaWduIjoiMTI5NzkyMzNjZWM5NDFmNGYxZGIxZjliYTA3MTQ5ZmQzMWZkMjVlYWM0ZWM0MmVlZGM3YjVlMWM2ZjU0NjU5ZiIsImVkZ2VzIjp7fX0seyJpZCI6ImVjMmM5OTVkLTI3ZTMtNDk0NS1iYTJhLTEzMGJiNjE0MWViMiIsImNyZWF0ZWRfdW5peCI6MTYyODA3NTk4NiwidXBkYXRlZF91bml4IjoxNjI4MDc2MDUyLCJvcyI6ImFuZHJvaWQiLCJhcmNoIjoiYW1kNjQiLCJ2ZXJzaW9uIjoiMC4xLjE2IiwicmlkIjoiYmFmeWJlaWJ5Znl3emxmdDJnM2hvdWFncXJxNmJzczNnNHF4Y3Z1c3VidGlnaTRwNnM0dWdocWZlaXUiLCJjcmMzMiI6IjNlNWE2ZDM0IiwiYXR0ciI6ImJveCIsImZvcmNpYmx5Ijp0cnVlLCJ0aXRsZSI6InR2LXJlbGVhc2UuMC4xLjE2LjMyIiwiZGV0YWlsIjoidHYtcmVsZWFzZS4wLjEuMTYuMzIiLCJzaWduIjoiOWVkODI3MTJhOWNlMzc5ZGE0OGI1ZTU0ZGVhZjk1NTMzNDQ3NzhlYWM4Mzk2ZWUzZmIzMWE1NTIxNmU1MGQ1MyIsImVkZ2VzIjp7fX0seyJpZCI6IjdlYjY1YzUwLWEzMWMtNDc1MS1iYWM4LTc1YzQ5MjFlMjMzOCIsImNyZWF0ZWRfdW5peCI6MTYyODY3OTEzMCwidXBkYXRlZF91bml4IjoxNjI4NzMzNDI0NDM5OTc1MDI2LCJvcyI6ImFuZHJvaWQiLCJhcmNoIjoiYW1kNjQiLCJ2ZXJzaW9uIjoiMC4xLjE3IiwicmlkIjoiYmFmeWJlaWVjZ3B4a2N6YXJzMmhiZmlkenJnNnJ6cXFlMm90ajRocjc0ZmhxaWZlc2xnZ2tnb3BsNXEiLCJjcmMzMiI6IjE3MWJlNDA2IiwiYXR0ciI6ImJveCIsImZvcmNpYmx5Ijp0cnVlLCJ0aXRsZSI6InR2LXJlbGVhc2UuMC4xLjE3LjM1IiwiZGV0YWlsIjoidHYtcmVsZWFzZS4wLjEuMTcuMzUiLCJzaWduIjoiMDY1NjAzMzdmMThlN2UzYjQ3ZTViOGJiZWFmYzBjOWQ5MDhiZTI3MmMwZTJhNTQ1NTY5ODNmODA3Zjc2NmI5YSIsImVkZ2VzIjp7fX0seyJpZCI6Ijk3YTUyZmM4LTFiMzQtNDQyMy05ZTIyLTc4ZDhkYzZmMWUwNCIsImNyZWF0ZWRfdW5peCI6MTYyODc2NDYzNTA1ODA4NzcwNCwidXBkYXRlZF91bml4IjoxNjI4NzY0NzI1NDgyMTU4NTQ5LCJvcyI6ImFuZHJvaWQiLCJhcmNoIjoiYW1kNjQiLCJ2ZXJzaW9uIjoiMC4xLjE4IiwicmlkIjoiYmFmeWJlaWZwMjVpM25oYWx1eWRkMzZ6cnliYXliZjNha3llMmtidmcyZWNpb2JyMmZhb3Zndmt6NnkiLCJjcmMzMiI6IjYyNTI4NDRmIiwiYXR0ciI6ImJveCIsImZvcmNpYmx5Ijp0cnVlLCJ0aXRsZSI6InR2LXJlbGVhc2UuMC4xLjE4LjM2IiwiZGV0YWlsIjoidHYtcmVsZWFzZS4wLjEuMTguMzYiLCJzaWduIjoiODRmZDgwMDkxZmQ4Yjg2N2Q0M2Y3ZjEyNTMxOTdiN2ZiYjFhZjE2ZTllYTI4MDIyMjc5ZjZiZjY3ZWE4ZDM1NSIsImVkZ2VzIjp7fX0seyJpZCI6IjZkMWFhNmNhLTgwNWEtNGU4OC04ZmQ4LTMzMzk2MTQ3ZjY3MyIsImNyZWF0ZWRfdW5peCI6MTYyODg0OTc1NzgwMzg5NDczNSwidXBkYXRlZF91bml4IjoxNjI4ODQ5ODkzMDg0MzY4NjEyLCJvcyI6ImFuZHJvaWQiLCJhcmNoIjoiYW1kNjQiLCJ2ZXJzaW9uIjoiMC4xLjE5IiwicmlkIjoiYmFmeWJlaWVsdnpmcG43NHkybDR6aW5ncXcyaHJzM2Ftbnd3eHJ3NGVpeXN4N3Rtb3A1b2N1a2JzN3UiLCJjcmMzMiI6IjI2N2U3M2I5IiwiYXR0ciI6ImJveCIsImZvcmNpYmx5Ijp0cnVlLCJ0aXRsZSI6InR2LXJlbGVhc2UuMC4xLjE5LjM3IiwiZGV0YWlsIjoidHYtcmVsZWFzZS4wLjEuMTkuMzciLCJzaWduIjoiMjg2M2JiM2ZkYjk1NjY2NjExMmM4MjI4YTJmMzdjYmYxNDhmMzY5MWYxN2M0MGYwODg2MDRlYjk2OWJmYTM0YSIsImVkZ2VzIjp7fX0seyJpZCI6IjU3YmEwYTdkLTY5ZjQtNGY1NS05ZjI4LTFjY2I3MjUxYThlNyIsImNyZWF0ZWRfdW5peCI6MTYyOTA4MzY2MjczMDEzNzY0OCwidXBkYXRlZF91bml4IjoxNjI5MDgzODQwNTI5MTgyMTk3LCJvcyI6ImFuZHJvaWQiLCJhcmNoIjoiYW1kNjQiLCJ2ZXJzaW9uIjoiMC4xLjE5IiwicmlkIjoiYmFmeWJlaWFtMml3aDNtaXg1cmo1bmRsM3BmdGNoajRhdHd5NXN2dGxtazdnY2l5cHJodnFveDV6c20iLCJjcmMzMiI6IjY2NjgxNjdlIiwiYXR0ciI6ImFwcCIsImZvcmNpYmx5Ijp0cnVlLCJ0aXRsZSI6ImFwcC1yZWxlYXNlLjAuMS4xOSIsImRldGFpbCI6ImFwcC1yZWxlYXNlLjAuMS4xOSIsInNpZ24iOiI2OTBjMGNjNWVmNDU5ZjAzMjA1OWFmZTFlMTUzZGY0YmFkZjJlYjZkMjE1ZDE1OTFjNTViNWY1MTMwNDk4ZTJmIiwiZWRnZXMiOnt9fSx7ImlkIjoiN2YzNDk1ODQtOTgyZS00OTgzLTllMTgtYzQ5MTdiYTU3NjUyIiwiY3JlYXRlZF91bml4IjoxNjI5MzU0MTQ3OTA4MDQwMDM1LCJ1cGRhdGVkX3VuaXgiOjE2MjkzNTQyNjI1MDEwOTA2NzksIm9zIjoiYW5kcm9pZCIsImFyY2giOiJhbWQ2NCIsInZlcnNpb24iOiIwLjEuMjEiLCJyaWQiOiJiYWZ5YmVpZWltY3IzZm53enE2dzdwMjJtd3FncG9kdnlwaXVwNzVmZnA1cG95emQ3cTZ4cHQyZGZtcSIsImNyYzMyIjoiNTYwMjU5ODQiLCJhdHRyIjoiYm94IiwiZm9yY2libHkiOnRydWUsInRpdGxlIjoidHYtcmVsZWFzZS4wLjEuMjEuMzkiLCJkZXRhaWwiOiJ0di1yZWxlYXNlLjAuMS4yMS4zOSIsInNpZ24iOiIwNGYyNjRmNDlkYmJlOTc0OWEwMjI2YmIzOTk2MDIxNDk4NzBjNWIzYTdjNDc4NWEwYjliNTE0M2M1NDJkMWZkIiwiZWRnZXMiOnt9fSx7ImlkIjoiYmIzOWNhYWMtOGFhYy00MDE3LTlmNTAtYmYxY2YyM2I5YTBlIiwiY3JlYXRlZF91bml4IjoxNjI5MzU0Mjg0OTIwNDI3NDIzLCJ1cGRhdGVkX3VuaXgiOjE2MjkzNTQzOTM4NzYwODAyMDAsIm9zIjoiYW5kcm9pZCIsImFyY2giOiJhbWQ2NCIsInZlcnNpb24iOiIwLjEuMjEiLCJyaWQiOiJiYWZ5YmVpZHU1bG5wajVjNm5qbzJ4Y3h3bG9teHpoamZpM2kydWNtbmVhaWI0bXh4cXh6cTRveWl2YSIsImNyYzMyIjoiYWQ0Mjk4NzMiLCJhdHRyIjoiYXBwIiwiZm9yY2libHkiOnRydWUsInRpdGxlIjoiYXBwLXJlbGVhc2UuMC4xLjIxLjI1IiwiZGV0YWlsIjoiYXBwLXJlbGVhc2UuMC4xLjIxLjI1Iiwic2lnbiI6IjBhMzQ5OTViNDczOTNkMzE2NjUzMjE5MGM0MDBjZGM0OTBmOTRhNDZlZDU5Zjc3MDRkZmNmM2VkMDFhMWYxNTUiLCJlZGdlcyI6e319LHsiaWQiOiI5M2ZkYzA0MC0yZjZkLTRkODAtOGNiYy01N2Y4MzM4YzgwY2MiLCJjcmVhdGVkX3VuaXgiOjE2MzM2NzY0MTQ3NzYyOTUwODcsInVwZGF0ZWRfdW5peCI6MTYzMzY3NjYzMTkzNTUwOTY1Mywib3MiOiJhbmRyb2lkIiwiYXJjaCI6ImFtZDY0IiwidmVyc2lvbiI6IjAuMS4yMyIsInJpZCI6ImJhZnliZWlleWZwNmI3Z2J0Y2E0dTZtNTNndnJ5cHZ1MmZ4cW53M3JyMmd3ZmJqaGVtbjZia3h6aWw0IiwiY3JjMzIiOiJjYzEwZTI0NCIsImF0dHIiOiJib3giLCJmb3JjaWJseSI6dHJ1ZSwidGl0bGUiOiJ0di1yZWxlYXNlLjAuMS4yMy40MSIsImRldGFpbCI6InR2LXJlbGVhc2UuMC4xLjIzLjQxIiwic2lnbiI6ImViYTQ1N2ZiYTYzZmRiODFiNjY4NTcyNGE1YjkxYzg0MTk3Zjk1MGYyNmNmOTc0NzkxMGRhMzg1ZjJkMjczNmMiLCJlZGdlcyI6e319LHsiaWQiOiI1ZDA1YjZlNS02NWRhLTRlOGEtOThiMy05ZjE3ODMwNjZkMzgiLCJjcmVhdGVkX3VuaXgiOjE2MzM2Nzc3MTMzMTgwNTIzMzksInVwZGF0ZWRfdW5peCI6MTYzNDEyMzk4MzU4MjcxNzk3Nywib3MiOiJhbmRyb2lkIiwiYXJjaCI6ImFtZDY0IiwidmVyc2lvbiI6IjAuMS4yMyIsInJpZCI6ImJhZnliZWlmZnprM3k2Z2d5dHdydGd2Zm82bmhkZmQzcHl2cm13M255eGFzdnFkcGpjZW54eHE3cTd1IiwiY3JjMzIiOiI1YjA0NDczNiIsImF0dHIiOiJhcHAiLCJmb3JjaWJseSI6dHJ1ZSwidGl0bGUiOiJhcHAtcmVsZWFzZS4wLjEuMjMuMjYiLCJkZXRhaWwiOiJhcHAtcmVsZWFzZS4wLjEuMjMuMjYiLCJzaWduIjoiNzgwODI3MjBhYzcyYWQ3YzkyNmQ3OGUwNzY5OTE0ODNjMTc3ZjZiOGE4NjYwNzFhNTI3YTliYTYxNjNkYTQyMyIsImVkZ2VzIjp7fX0seyJpZCI6IjIzNWY5ZmJiLTEzODctNDUwZC1iNzAxLWQ2MjBmN2NmZWM2MCIsImNyZWF0ZWRfdW5peCI6MTYzNzMxOTQ5MzU0NTgzOTIwOSwidXBkYXRlZF91bml4IjoxNjM3MzE5NTc5MTgyODA5OTA0LCJvcyI6ImFuZHJvaWQiLCJhcmNoIjoiYW1kNjQiLCJ2ZXJzaW9uIjoiMC4xLjI1IiwicmlkIjoiYmFmeWJlaWNqYmNuYmFqZjM3emc3aGZ6ZHVtYnh5NW9oZTV5ZXp1YTdlamJrZ2t2b2VwNjJlY2hsNnkiLCJjcmMzMiI6ImM4NDE0OWRlIiwiYXR0ciI6ImJveCIsImZvcmNpYmx5Ijp0cnVlLCJ0aXRsZSI6InR2LXJlbGVhc2UuMC4xLjI1LjQzIiwiZGV0YWlsIjoidHYtcmVsZWFzZS4wLjEuMjUuNDMiLCJzaWduIjoiOGIxODZmZmZjYTAyNGM0NzJjYWJkNGMxNWY2MTRkMjFjOGQ3NDQ1M2VjMjlmZDY3ODliYjhmODE4OWZhNWI2NyIsImVkZ2VzIjp7fX0seyJpZCI6IjgxMTdkYjk4LTQyMTQtNDAwNC05NDhhLWU0ZmUzZTM3M2MxYiIsImNyZWF0ZWRfdW5peCI6MTYzNzMxOTYwNzQwNzI4MDk5MCwidXBkYXRlZF91bml4IjoxNjM3MzIxNzQ3MjY5NzE2ODQ5LCJvcyI6ImFuZHJvaWQiLCJhcmNoIjoiYW1kNjQiLCJ2ZXJzaW9uIjoiMC4xLjI1IiwicmlkIjoiYmFmeWJlaWZtdnJ4dzVxa2t1MmhheDYzN2JmNXJxb3RyaTNlemJibGVnd3RiaWpqbGdnc3htYXJxd2kiLCJjcmMzMiI6IjgxOGJiNzI0IiwiYXR0ciI6ImFwcCIsImZvcmNpYmx5Ijp0cnVlLCJ0aXRsZSI6ImFwcC1yZWxlYXNlLjAuMS4yNS4yNyIsImRldGFpbCI6ImFwcC1yZWxlYXNlLjAuMS4yNS4yNyIsInNpZ24iOiJiNDUxYmVlNGQwMTMyMTRhNzhmZjhiNGUyMjU1ZWExMjkwYTlkMDY4YzZhYWJmNmYyMDA5ZDRkZTE1YzBkYjZhIiwiZWRnZXMiOnt9fSx7ImlkIjoiZGViZmYxNGYtMDQ3Zi00YTc1LWI4ZmQtZWQyMzI5NmUzYzI4IiwiY3JlYXRlZF91bml4IjoxNjQ1OTM3OTkwNDc2Mjg2MjAyLCJ1cGRhdGVkX3VuaXgiOjE2NDU5MzgxNzIzNTI1OTUxNDMsIm9zIjoiYW5kcm9pZCIsImFyY2giOiJhbWQ2NCIsInZlcnNpb24iOiJ2MS4wLjAiLCJyaWQiOiJiYWZ5YmVpZXVwazJvZXVyZWFsYnJhanJrdmF2dWp6emhsYmlnb2tidTZrMjJuamw1bnpvdWY1M3R4cSIsImF0dHIiOiJjb3JlIiwidGl0bGUiOiJjb3JlIHVwZGF0ZSBwYWNrYWdlIiwiZGV0YWlsIjoidXBkYXRlIiwic2lnbiI6IjAzOTg4ZWQxYjMwY2QyZThjZGUzYzc3ODg4OTUxZmY3NjI0M2ZlMTlhNDQwNzdiYmM4Y2ZiNWUxYjg2ODM2MGIiLCJlZGdlcyI6e319LHsiaWQiOiIzNjRhZTNiNi1jZDczLTRjMjItYjRhYy1hYjMxODRiMDlkMTMiLCJjcmVhdGVkX3VuaXgiOjE2NDY2NDY0NjY3NjUxMDE4NzUsInVwZGF0ZWRfdW5peCI6MTY0NjY0Njk4NjQ5NjAxNjMzMSwib3MiOiJhbmRyb2lkIiwiYXJjaCI6ImFybTY0IiwidmVyc2lvbiI6IjEuMC4xLjIiLCJyaWQiOiJiYWZ5YmVpaG9ld3dmbmZreTM2bWZ5Z2FzM2V1cmd3dm93N28zZ2JhanZ1Z21jbXBrYWo3dnB1dmJheSIsImNyYzMyIjoiNTZhZjZiZDQiLCJhdHRyIjoiY29yZSIsImFwcF9uYW1lIjoiTGluayBzZXJ2aWNlIiwiZm9yY2libHkiOnRydWUsInRydW5jYXRlIjp0cnVlLCJ0aXRsZSI6IjEuMC4xLjLmm7TmlrAiLCJkZXRhaWwiOiIxLjAuMS4yXG4xLiA15YiG6ZKf5Yi35pawIiwic2lnbiI6IjdhNzY3MWZjOTgyNjA5ZTU0MDVjM2UwNTU5ZTI1MDdiZGNmMGZkY2VlOWM5Y2ZkNTM5Y2FmNmViMGY4MjE0OWEiLCJlZGdlcyI6e319LHsiaWQiOiJkNTNjYmE2ZS05NzIxLTQxZWYtOWVmNS1lODY3ZDkzZWJiNTIiLCJjcmVhdGVkX3VuaXgiOjE2NDY3OTI4NDMyMTcyMzUwMjIsInVwZGF0ZWRfdW5peCI6MTY0Njc5MzU5MzUzNjIxNDAyNSwib3MiOiJhbmRyb2lkIiwiYXJjaCI6ImFybTY0IiwidmVyc2lvbiI6IjEuMC4yIiwicmlkIjoiYmFmeWJlaWZwZm4ydWF6d292c3gzZ3ZkMm02azVsendxd215eWxwY2NpM3BxbHRjNzZyZWJzaWozdGkiLCJjcmMzMiI6ImNhNzE5MjNjIiwiYXR0ciI6ImNvcmUiLCJhcHBfbmFtZSI6IlR2Qm94TGlua1NlcnZpY2UxLjAuMiIsImZvcmNpYmx5Ijp0cnVlLCJ0cnVuY2F0ZSI6dHJ1ZSwidGl0bGUiOiJUdkJveExpbmtTZXJ2aWNlMS4wLjIiLCJkZXRhaWwiOiLmm7TmlrA15YiG6ZKf5Yi35pawIiwic2lnbiI6IjRkN2ExMzIzYTlhYTFmZGNkOGQ1MTdmNjJkNTIxYjQyNDc4ZGNiYjk1ODBkMWMxYzdjZDNlMmNmMWFhMWE2MWQiLCJlZGdlcyI6e319LHsiaWQiOiIwZWIwODNhOS04Y2EzLTQ5YzQtYWIxOS1kM2RkNDM4YmY3ZDIiLCJjcmVhdGVkX3VuaXgiOjE2NDY4MDc1MjYxMzk0MDcxNDQsInVwZGF0ZWRfdW5peCI6MTY0NjgwNzgxODY1MjI2MTM5Miwib3MiOiJhbmRyb2lkIiwiYXJjaCI6ImFtZDY0IiwidmVyc2lvbiI6IjEuMC4zIiwicmlkIjoiYmFmeWJlaWRsbmtrY2JtamtnNGtieXZlaGFjamRyanNiNTJkbjNncHBmcGNyeHdyN291azN4aGRxc3EiLCJjcmMzMiI6ImZhODM1YWIwIiwiYXR0ciI6ImNvcmUiLCJhcHBfbmFtZSI6IlRWQm94TGlua1NlcnZpY2UxLjAuMyIsImZvcmNpYmx5Ijp0cnVlLCJ0aXRsZSI6IlRWQm94TGlua1NlcnZpY2UxLjAuMyIsImRldGFpbCI6IueJiOacrFRWQm94TGlua1NlcnZpY2UxLjAuMyIsInNpZ24iOiJhMWU5Mjk5NGNlYjQ4M2NjMzk1MjY5ZTc3NmFkYzlmY2I3ZmE5MWU1OGEwM2M2ZDJjNjNkNzJhYTk3NDBlZjI5IiwiZWRnZXMiOnt9fSx7ImlkIjoiOGUxOGY4NWQtYTExOC00ZjQzLTljMjctYWE1MTM0MGJkODgwIiwiY3JlYXRlZF91bml4IjoxNjQ3MzE0Mzk4NjQ0MDQ4MzU1LCJ1cGRhdGVkX3VuaXgiOjE2NDczMTQ0MzY5NzU5NTY3NDMsIm9zIjoiYW5kcm9pZCIsImFyY2giOiJhcm02NCIsInZlcnNpb24iOiIxLjAuNCIsInJpZCI6ImJhZnliZWlhZnRmNmNrYWtpa2tpZDZjaTZoamo3d2p1Y2N4NHppbHptaTd0aWk2YmpldWNscjU3cWt5IiwiY3JjMzIiOiJmOGZlMDVjNSIsImF0dHIiOiJjb3JlIiwiYXBwX25hbWUiOiJUdkJveExpbmtTZXJ2aWNlMS4wLjQiLCJmb3JjaWJseSI6dHJ1ZSwidGl0bGUiOiJUdkJveExpbmtTZXJ2aWNlMS4wLjQiLCJkZXRhaWwiOiLmm7TmlrA15YiG6ZKf5Yi35pawIiwic2lnbiI6ImM3ZGEwOTQxOWYyZjdlNGE2MTZiMGYzMThkNmEzMmZhZGFkZjM3ZDFlODAxNmRjNmE4ZGI3NDllZDQ0ZDZlMzkiLCJlZGdlcyI6e319LHsiaWQiOiIzNTNhYTYwMS02OTZjLTQ2MGItOTc4NS0zNDBhNjQ1NGJkOTYiLCJjcmVhdGVkX3VuaXgiOjE2NDgyMDQ4OTYwMzAyNDEzMzUsInVwZGF0ZWRfdW5peCI6MTY0ODIwNTMzNzMwNzE4MTcxNCwib3MiOiJhbmRyb2lkIiwiYXJjaCI6ImFybTY0IiwidmVyc2lvbiI6IjAuMi4xIiwicmlkIjoiYmFmeWJlaWN3Ym5rM3IyNjNneDdpcG9ocHp5bzRud3dmejV0anVoM3BwN3NweGhybWE1Z3Rxd2xkbXUiLCJjcmMzMiI6ImExMmM5Y2YwIiwiYXR0ciI6ImJveCIsImFwcF9uYW1lIjoibW90b21vdG8gdHZib3ggYXBwIiwiZm9yY2libHkiOnRydWUsInRpdGxlIjoibW90b21vdG8gdHZib3ggYXBwIiwiZGV0YWlsIjoi5rWL6K+V5pu05paw5Yqf6IO9Iiwic2lnbiI6IjY2MzMxMWZlZmVkZmQyYWNiMWY1Mzc1ZDc4NmJmMTJmMmI2YjRmZjUyYmM4OWRhOTlmOGI2NTAzZTYxODRhOTIiLCJlZGdlcyI6e319LHsiaWQiOiI0OWJiYjllYi0yOGZkLTQ2NzUtYmE5My0zMzAwYTUzYjUyMjIiLCJjcmVhdGVkX3VuaXgiOjE2NDg0NTYzNjQ2Mjc5NTkzNTcsInVwZGF0ZWRfdW5peCI6MTY0ODQ1NzYzNDUwMTk5MDQxNSwib3MiOiJhbmRyb2lkIiwiYXJjaCI6ImFybTY0IiwidmVyc2lvbiI6IjEuMi4wIiwicmlkIjoiYmFmeWJlaWZ3bGE1eDJ2MzZsNW9xa2x3enp2MzVlMjJ4N2V0YXJhYnZpZnpibWlraWt5dWM2Njcya3EiLCJjcmMzMiI6IjQ1NWFlYTgyIiwiYXR0ciI6ImJveCIsImFwcF9uYW1lIjoibW90b21vdG8gdHZib3ggYXBwIiwiZm9yY2libHkiOnRydWUsInRpdGxlIjoibW90b21vdG8gdHZib3ggYXBwIiwiZGV0YWlsIjoi5rWL6K+V5pu05paw5Yqf6IO9Iiwic2lnbiI6ImM0YmE5ZmFlYzg4MjUxMWQzMDlhNmFiZTU5Y2Y0MTA2ZmQ4MWQ4YzMzMjM4MTk3NDEzZWJhNmZkNDg1MzY3ODkiLCJlZGdlcyI6e319LHsiaWQiOiI4MTcxNjI2Ny02NTYwLTRiOGUtYjRlYy1jZGJkN2U0YzBiZWYiLCJjcmVhdGVkX3VuaXgiOjE2NTU4MDIwODI0NTA2MDAzMTgsInVwZGF0ZWRfdW5peCI6MTY1NTgwMjQwNjUyMjg2MDAwNywib3MiOiJhbmRyb2lkIiwiYXJjaCI6ImFybTY0IiwidmVyc2lvbiI6IjEuMy4zIiwicmlkIjoiYmFmeWJlaWR3eW9mZWhpZGNuaWViemtkNHJ5YmZhdXV4M3dteDJvcXAzemVzM3Jhb3dkZTZvMjY2ZmkiLCJjcmMzMiI6IjE2NGRmOGY0IiwiYXR0ciI6ImJveCIsImFwcF9uYW1lIjoibW90b21vdG8gdHZib3ggYXBwIiwiZm9yY2libHkiOnRydWUsInRpdGxlIjoibW90b21vdG8gdHZib3ggYXBwIiwiZGV0YWlsIjoi5rWL6K+V5pu05paw5Yqf6IO9Iiwic2lnbiI6IjA5N2Q3YWJmZDljNmJiZmU1MTJkOThjYjU0OThmMjRkOTQzMDY1OTFmZTIxZmNiOTgxNTYzNTM2ZjBiNzJlMDMiLCJlZGdlcyI6e319XQ=="},"node":{}}`)
			jsonContent4, err4 := ParseJSONContent(oldc4)
			if err4 != nil {
				t.Fatal(err4)
				return
			}
			if jsonContent4.Message != nil {
				t.Logf("data size:%+v", len(jsonContent4.Message.Data))
				if len(jsonContent4.Message.Data) > 256 {
					t.Logf("data:%+v", string(jsonContent4.Message.Data[:256]))
				} else {
					t.Logf("data:%+v", string(jsonContent4.Message.Data))
				}
			}
			//t.Logf("json4:%+v", string(jsonContent4.JSON()))
			v2oldc4 := jsonContent4.JSONV2()
			jsonContent5, err5 := ParseJSONContent(v2oldc4)
			if err5 != nil {
				t.Fatal(err5)
				return
			}
			//t.Logf("json5v1:%+v", string(jsonContent5.JSON()))
			//t.Logf("json5v2:%+v", string(jsonContent5.JSONV2()))
			if jsonContent5.Message != nil {
				//t.Logf("data size:%+v", len(jsonContent5.Message.Data))
				if len(jsonContent5.Message.Data) > 256 {
					//t.Logf("data:%+v", string(jsonContent5.Message.Data[:256]))
				} else {
					//t.Logf("data:%+v", string(jsonContent5.Message.Data))
				}
			}

			bytes := jsonContent5.Bytes()
			c1, err := ParseContent(bytes)
			if err != nil {
				t.Fatal(err)
			}
			//t.Logf("content5 :%+v", string(c1.JSON()))
			//t.Logf("content5:%+v", string(c1.JSONV2()))
			if c1.Message != nil {
				//t.Logf("data size:%+v", len(c1.Message.Data))
				if len(c1.Message.Data) > 256 {
					//t.Logf("data:%+v", string(c1.Message.Data[:256]))
				} else {
					//t.Logf("data:%+v", string(c1.Message.Data))
				}
			}

			oldcb := []byte(`{"version":"0.0.1","message":{"version":1},"type":6}`)
			c2, err := ParseJSONContent(oldcb)
			if err != nil {
				t.Fatal(err)
			}
			t.Logf("data:%+v", c2.JSON())
			bc2 := c2.Bytes()
			cb2, err := ParseContent(bc2)
			if err != nil {
				t.Fatal(err)
			}
			t.Logf("data:%+v", cb2.JSON())
		})
	}
}
