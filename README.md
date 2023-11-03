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
	ongkir "github.com/agung96tm/rajaongkir-go"
)

func main() {
	api := ongkir.NewAPI(ongkir.NewConfig(
		"key",
		ongkir.AccountStarter,
	))

	// --------------------------------------
	// Provinces
	// --------------------------------------
	provResp, err := api.GetProvinces(nil)
	//provResp, err := api.GetProvinces(ongkir.ProvinceParams{ID: "1"})

	if err != nil {
		panic(err)
	}
	
	fmt.Println(provResp.Results)
	// provResp.Status { Code, Description }
	// provResp.Results []Province

	// --------------------------------------
	// Cities
	// --------------------------------------
	citiesResp, err := api.GetCities(nil)
	//citiesResp, err := api.GetCities(ongkir.CityParams{ID: "1", ProvinceID: "1"})

	if err != nil {
		panic(err)
	}
	
	fmt.Println(citiesResp.Results)
	// citiesResp.Results []City
	// citiesResp.Status { Code, Description }
}
```

## Contributors
* Agung Yuliyanto: [Github](https://github.com/agung96tm), [LinkedIn](https://www.linkedin.com/in/agung96tm/)

Hey you! Interest become contribtors? Then make issue or PR :)
