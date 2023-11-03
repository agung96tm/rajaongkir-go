package rajaongkir_go

type ProvinceParams struct {
	ID string `query:"id" json:"id"`
}

type CityParams struct {
	ID         string `query:"id" json:"id"`
	ProvinceID string `query:"province" json:"province"`
}
