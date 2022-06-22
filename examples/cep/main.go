package main

import (
	"fmt"

	"github.com/isaqueveras/brasilapi-go/cep"
)

func main() {
	var data = cep.CEP{
		Cep:     nil,
		State:   nil,
		City:    nil,
		Service: nil,
	}

	fmt.Println(data)
}
