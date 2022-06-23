package main

import (
	"fmt"

	"github.com/isaqueveras/brasilapi-go"
)

func main() {
	zp, err := brasilapi.GetZipCode("63610000")
	if err != nil {
		fmt.Println(err)
	}

	if zp != nil {
		fmt.Println(*zp.City)
	}
}
