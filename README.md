Rajakongkir Go
=======================================
[![Go Report Card](https://goreportcard.com/badge/github.com/agung96tm/rajaongkir-go)](https://goreportcard.com/report/github.com/agung96tm/rajaongkir-go)
[![Codecov](https://codecov.io/gh/agung96tm/rajaongkir-go/branch/main/graph/badge.svg)](https://codecov.io/gh/agung96tm/rajaongkir-go)
[![Version](https://img.shields.io/github/go-mod/go-version/agung96tm/rajaongkir-go)](https://go.dev/doc/go1.20)
[![Release](https://img.shields.io/github/v/release/agung96tm/rajaongkir-go.svg)](https://github.com/agung96tm/rajaongkir-go/releases)
[![License](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/agung96tm/rajaongkir-go/blob/master/LICENSE)
[![GoDoc](https://pkg.go.dev/badge/github.com/agung96tm/rajaongkir-go)](https://pkg.go.dev/github.com/agung96tm/rajaongkir-go)

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
	//provResp, err := api.GetProvinces(&ongkir.ProvinceParams{ID: "1"})

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
	//citiesResp, err := api.GetCities(&ongkir.CityParams{ID: "1", ProvinceID: "1"})

	if err != nil {
		panic(err)
	}
	
	fmt.Println(citiesResp.Results)
	// citiesResp.Status { Code, Description }
	// citiesResp.Results []City

	// --------------------------------------
	// Cost
	// -------------------------------------- 
	costResp, err := app.GetCost(&ongkir.CostParams{
		Courier:     JNE, // CouriersParams(JNE),
		Origin:      "501",
		Destination: "114",
		Weight:      1,
	})
	
	fmt.Println(costResp.Results)
	// costResp.Results []Courier
	// costResp.Results[0].Courier {Code, Name, Costs}
}
```

## Contributors
* Agung Yuliyanto: [Github](https://github.com/agung96tm), [LinkedIn](https://www.linkedin.com/in/agung96tm/)

Hey you! Interest become contributor? Then make issue or PR :)
