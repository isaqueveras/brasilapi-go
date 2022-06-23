<h1 align='center'>brasilapi-go</h1>
<p align='center'>Cliente em Golang para se comunicar com BrasilApi</p>

## CEP (zipcode)
```go
package main

import (
	"fmt"

	"github.com/isaqueveras/brasilapi-go"
)

func main() {
	zc, err := brasilapi.GetZipCode("63900-193")
	if err != nil {
		fmt.Println(err)
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
```
