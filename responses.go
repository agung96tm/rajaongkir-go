package rajaongkir_go

type ResponseStatus struct {
	Code        uint   `json:"code"`
	Description string `json:"description"`
}

type ResponseBase struct {
	query   interface{}
	status  ResponseStatus
	results any
}

func (r ResponseBase) GetResults() any {
	return r.results
}

type Province struct {
	ProvinceID int `json:"province_id"`
	Province   int `json:"province"`
}

type ResponseProvinces struct {
	ResponseBase
	results []Province
}

func (p ResponseProvinces) GetResults() any {
	return p.results
}

type City struct {
	CityID     int    `json:"city_id"`
	City       string `json:"city_name"`
	ProvinceID int    `json:"province_id"`
	Province   string `json:"province"`
	PostalCode string `json:"postal_code"`
}

type ResponseCities struct {
	ResponseBase
	results []City
}

func (p ResponseCities) GetResults() any {
	return p.results
}
