package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/isaqueveras/brasilapi-go"
	"github.com/isaqueveras/brasilapi-go/ibge"
	"github.com/isaqueveras/juazeiro"
)

func main() {
	conn, err := juazeiro.NewClient(brasilapi.ServerURL)
	if err != nil {
		panic(err)
	}

	client := ibge.NewIbgeClient(conn)

	input := &ibge.Identifier{UF: pointer(ibge.UfCe)}
	input.NewParams()
	input.WithParamProviders(pointer(ibge.ProviderWikipedia))

	var couties *[]ibge.Response
	if couties, err = client.ObtainCounties(context.Background(), input); err != nil {
		panic(err)
	}

	str, _ := json.Marshal(couties)
	fmt.Println(string(str))
}

func pointer[t any](value t) *t {
	return &value
}
