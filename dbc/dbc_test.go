package dbc

import (
	"context"
	"fmt"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestDBC(t *testing.T) {
	dbc, err := Open("test")
	if err != nil {
		panic(err)
	}
	x := dbc.Media().C.InformationV1.Query().AllX(context.Background())
	fmt.Println(x)
}
