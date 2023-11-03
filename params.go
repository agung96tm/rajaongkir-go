package rajaongkir_go

type ProvinceParams struct {
	ID string `query:"id"`
}

type CityParams struct {
	ID         string `query:"id"`
	ProvinceID string `query:"province"`
}
