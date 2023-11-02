package rajaongkir_go

type ResponseStatus struct {
	Code        uint   `json:"code"`
	Description string `json:"description"`
}

type RajaongkirResponse[T any] struct {
	Query   any            `json:"query"`
	Status  ResponseStatus `json:"status"`
	Results T              `json:"results"`
}

type ResponseBase[T any] struct {
	Rajaongkir RajaongkirResponse[T] `json:"rajaongkir"`
}

func (r ResponseBase[T]) GetResults() T {
	return r.Rajaongkir.Results
}

func (r ResponseBase[T]) GetStatus() ResponseStatus {
	return r.Rajaongkir.Status
}

type Province struct {
	ProvinceID int `json:"province_id"`
	Province   int `json:"province"`
}

type ResponseProvinces struct {
	ResponseBase[[]Province]
}

type City struct {
	CityID     int    `json:"city_id"`
	City       string `json:"city_name"`
	ProvinceID int    `json:"province_id"`
	Province   string `json:"province"`
	PostalCode string `json:"postal_code"`
}

type ResponseCities struct {
	ResponseBase[[]City]
}
