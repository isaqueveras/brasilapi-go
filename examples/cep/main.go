package main

import (
	"fmt"

	"github.com/isaqueveras/brasilapi-go"
)

func main() {
	zc, err := brasilapi.GetZipCode("63900-193")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(*zc.Cep)
	fmt.Println(*zc.City)
	fmt.Println(*zc.Service)
	fmt.Println(*zc.State)

	if zc.Location != nil {
		fmt.Println(*zc.Location.Type)
	}

	if zc.Location != nil && zc.Location.Coordenates != nil {
		fmt.Println(*zc.Location.Coordenates.Longitude)
		fmt.Println(*zc.Location.Coordenates.Latitude)
	}
}
