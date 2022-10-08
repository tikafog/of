package utils

import (
	"fmt"
	"reflect"
	"testing"
)

func TestHash(t *testing.T) {
	type args struct {
		hashes []string
	}
	tests := []struct {
		name string
		args args
		want []uint64
	}{
		{
			name: "",
			args: args{
				hashes: []string{
					"A", "B", "C", "D", "E", "F", "G",
				},
			},
			want: []uint64{
				1371800463213966980, 1440108869279352788, 7622816676894411104, 7884081726600927225, 10116856886307834793, 17504886469506087110, 17568128090902181845,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Hash(tt.args.hashes...)
			ints := BytesToInt64(got)
			fmt.Println("hashed", ints)
			if !reflect.DeepEqual(ints, tt.want) {
				t.Errorf("Hash() = %v, want %v", got, tt.want)
			}
		})
	}
}
