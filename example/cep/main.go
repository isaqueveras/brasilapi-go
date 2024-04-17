package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/isaqueveras/brasilapi-go"
	"github.com/isaqueveras/brasilapi-go/cep"
	"github.com/isaqueveras/juazeiro"
)

func main() {
	conn, err := juazeiro.NewClient(brasilapi.ServerClient)
	if err != nil {
		panic(err)
	}

	client := cep.NewCepClient(conn)

	var zipcode *cep.Response
	if zipcode, err = client.GetZipCode(context.Background(), &cep.Identifier{
		Cep: pointer("63900-193"),
	}); err != nil {
		panic(err)
	}

	str, _ := json.Marshal(zipcode)
	fmt.Println(string(str))
}

func pointer[t any](value t) *t {
	return &value
}
