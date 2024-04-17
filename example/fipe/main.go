package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/isaqueveras/brasilapi-go"
	"github.com/isaqueveras/brasilapi-go/fipe"
	"github.com/isaqueveras/juazeiro"
)

func main() {
	conn, err := juazeiro.NewClient(brasilapi.ServerURL)
	if err != nil {
		panic(err)
	}

	client := fipe.NewFipeClient(conn)

	var brands *[]fipe.Response
	if brands, err = client.ObtainVehicleBrand(context.Background(), &fipe.Identifier{
		Type: pointer(fipe.VehicleTypeMotos),
	}); err != nil {
		panic(err)
	}

	str, _ := json.Marshal(brands)
	fmt.Println(string(str))
}

func pointer[t any](value t) *t {
	return &value
}
