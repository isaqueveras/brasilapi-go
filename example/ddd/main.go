package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/isaqueveras/brasilapi-go"
	"github.com/isaqueveras/brasilapi-go/ddd"
	"github.com/isaqueveras/juazeiro"
)

func main() {
	conn, err := juazeiro.NewClient(brasilapi.ServerURL)
	if err != nil {
		panic(err)
	}

	client := ddd.NewDddClient(conn)

	var zipcode *ddd.Response
	if zipcode, err = client.Obtain(context.Background(), &ddd.Identifier{
		Ddd: pointer(int64(88)),
	}); err != nil {
		panic(err)
	}

	str, _ := json.Marshal(zipcode)
	fmt.Println(string(str))
}

func pointer[t any](value t) *t {
	return &value
}
