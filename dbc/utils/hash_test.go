package utils

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/google/uuid"
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
			got := HashString(tt.args.hashes...)
			ints := BytesToInt64(got)
			fmt.Println("hashed", ints)
			if !reflect.DeepEqual(ints, tt.want) {
				t.Errorf("HashString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashUUID(t *testing.T) {
	idsList := []uuid.UUID{
		uuid.Must(uuid.Parse(`76bbaacd-478f-11ed-b9dd-d493900cc0d7`)),
		uuid.Must(uuid.Parse(`76bd0924-478f-11ed-b9dd-d493900cc0d7`)),
		uuid.Must(uuid.Parse(`76bd0924-478f-11ed-b9de-d493900cc0d7`)),
		uuid.Must(uuid.Parse(`76bd0924-478f-11ed-b9df-d493900cc0d7`)),
		uuid.Must(uuid.Parse(`76bd0924-478f-11ed-b9e0-d493900cc0d7`)),
		uuid.Must(uuid.Parse(`76bd0924-478f-11ed-b9e1-d493900cc0d7`)),
		uuid.Must(uuid.Parse(`76bd0924-478f-11ed-b9e2-d493900cc0d7`)),
		uuid.Must(uuid.Parse(`76bd0924-478f-11ed-b9e3-d493900cc0d7`)),
		uuid.Must(uuid.Parse(`76bd0924-478f-11ed-b9e4-d493900cc0d7`)),
		uuid.Must(uuid.Parse(`76bd0924-478f-11ed-b9e5-d493900cc0d7`)),
	}
	type args struct {
		ids []uuid.UUID
	}
	tests := []struct {
		name string
		args args
		want []uint64
	}{
		{
			name: "",
			args: args{
				ids: idsList,
			},
			want: []uint64{
				1251957131256883690, 1698423046219215847, 4183029526008001809, 6187102436395144405, 6962987168164101455, 11916952097348514975, 13713536874971927900, 15279878632808010638, 16034781008331307894, 18200916881414210053,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Hash(tt.args.ids...)
			ints := BytesToInt64(got)
			fmt.Println("hashed", ints)
			if !reflect.DeepEqual(ints, tt.want) {
				t.Errorf("HashString() = %v, want %v", got, tt.want)
			}
		})
	}
}
