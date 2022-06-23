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

	fmt.Println(zp)
}
