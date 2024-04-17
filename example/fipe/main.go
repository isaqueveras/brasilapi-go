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
	ctx := context.Background()

	// ============ BRANDS

	var brands *[]fipe.ResponseBrand
	if brands, err = client.ObtainVehicleBrand(ctx, &fipe.IdentifierBrand{
		Type: pointer(fipe.VehicleTypeMotos),
	}); err != nil {
		panic(err)
	}

	str, _ := json.Marshal(brands)
	fmt.Println("brands", string(str))

	// ============ PRICES

	inputPrice := &fipe.IdentifierPrice{FipeCode: pointer("001004-9")}
	inputPrice.NewParams()
	inputPrice.WithParamTabelaReferencia(pointer(int64(1)))

	var prices *[]fipe.ResponsePrice
	if prices, err = client.ObtainVehiclePrice(ctx, inputPrice); err != nil {
		panic(err)
	}

	str, _ = json.Marshal(prices)
	fmt.Println("prices", string(str))
}

func pointer[t any](value t) *t {
	return &value
}
