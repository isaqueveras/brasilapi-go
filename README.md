# brasilapi-go
Cliente em Golang para se comunicar com BrasilApi

## Cep (zipcode)
Para buscar dados referente a um CEP basta utilizar a função `GetZipCode(zipcode)`

```go
package main

import (
	"fmt"

	"github.com/isaqueveras/brasilapi-go"
)

func main() {
	zp, err := brasilapi.GetZipCode("63900-193")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(*zp.Cep)
	fmt.Println(*zp.City)
	fmt.Println(*zp.Service)
	fmt.Println(*zp.State)

	if zp.Location != nil {
		fmt.Println(*zp.Location.Type)
	}

	if zp.Location != nil && zp.Location.Coordenates != nil {
		fmt.Println(*zp.Location.Coordenates.Longitude)
		fmt.Println(*zp.Location.Coordenates.Latitude)
	}
}
```
