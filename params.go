package rajaongkir_go

type ProvinceParams struct {
	ID uint `json:"id"`
}

type CityParams struct {
	ID         uint `json:"id"`
	ProvinceID uint `json:"province"`
}
