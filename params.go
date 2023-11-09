package rajaongkir_go

import (
	"strings"
)

type ProvinceParams struct {
	ID string `query:"id" json:"id"`
}

type CityParams struct {
	ID         string `query:"id" json:"id"`
	ProvinceID string `query:"province" json:"province"`
}

type CostParams struct {
	Origin      string `query:"origin" json:"origin"`
	Destination string `query:"destination" json:"destination"`
	Weight      int    `query:"weight" json:"weight"`
	Courier     string `query:"courier" json:"courier"`
}

func CourierParams(couriers ...string) string {
	return strings.Join(couriers, ":")
}
