Rajakongkir Go
=======================================

## Installation
```shell
go get -u github.com/agung96tm/rajaongkir-go
```

## Usage

```go
package main

import (
	"fmt"
	ongkir "github.com/kachit/rajaongkir-go"
)

func main() {
	api := ongkir.NewAPI(ongkir.NewConfig(
		"key",
		ongkir.AccountStarter,
	))

	// --------------------------------------
	// Provinces
	// --------------------------------------
	provRes, err := api.GetProvinces(nil)
	if err != nil {
		panic(err)
	}
	// provRes.GetStatus() { Code, Description }
	// provRes.GetResults() []Province

	// --------------------------------------
	// City
	// --------------------------------------
	cityRes, err := api.GetCities(nil)
	if err != nil {
		panic(err)
	}
	// cityRes.GetStatus() { Code, Description }
	// cityRes.GetResults() []Province

}
```