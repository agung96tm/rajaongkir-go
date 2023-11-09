package rajaongkir_go

import (
	"encoding/json"
	"errors"
)

type ResponseStatus struct {
	Code        uint   `json:"code"`
	Description string `json:"description"`
}

type ResponseBase[T interface{}, P interface{}] struct {
	Query   P              `json:"query"`
	Status  ResponseStatus `json:"status"`
	Results []T            `json:"results"`
}

func (b *ResponseBase[T, P]) Unmarshal(v []byte) error {
	var rawResp map[string]json.RawMessage
	_ = json.Unmarshal(v, &rawResp)

	if err := json.Unmarshal(rawResp["rajaongkir"], &b); err != nil {
		var rawNested map[string]map[string]json.RawMessage
		_ = json.Unmarshal(v, &rawNested)

		// Problem because results obj instead of list obj, then transform it
		var result T
		if err := json.Unmarshal(rawNested["rajaongkir"]["results"], &result); err == nil {
			b.Results = []T{result}
		} else {
			return errors.New("fail to parse json")
		}
	}

	return nil
}

func (b *ResponseBase[T, P]) StatusValid() error {
	if b.Status.Code >= 200 && b.Status.Code < 300 {
		return nil
	}
	return errors.New(b.Status.Description)
}

type Province struct {
	ProvinceID string `json:"province_id"`
	Province   string `json:"province"`
}

type ResponseProvinces struct {
	ResponseBase[Province, ProvinceParams]
}

type City struct {
	CityID     string `json:"city_id"`
	City       string `json:"city_name"`
	ProvinceID string `json:"province_id"`
	Province   string `json:"province"`
	PostalCode string `json:"postal_code"`
}

type ResponseCities struct {
	ResponseBase[City, CityParams]
}

type Cost struct {
	Value string `json:"value"`
	Etd   string `json:"etd"`
	Note  string `json:"note"`
}

type CourierCost struct {
	Service     string `json:"service"`
	Description string `json:"description"`
	Cost        []Cost `json:"cost"`
}

type Courier struct {
	Code  string        `json:"code"`
	Name  string        `json:"name"`
	Costs []CourierCost `json:"costs"`
}

type ResponseCost struct {
	ResponseBase[Courier, CostParams]
}
