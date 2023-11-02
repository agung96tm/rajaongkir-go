package rajaongkir_go

type API struct {
	config Config
}

func NewAPI(config Config) API {
	return API{
		config: config,
	}
}

func (a API) GetProvinces(opts ProvinceParams) (ResponseProvinces, error) {
	return get[ResponseProvinces](a.config.BaseUrl, "province", a.config.Header, urlValues(opts))
}

func (a API) GetCities(opts CityParams) (ResponseCities, error) {
	return get[ResponseCities](a.config.BaseUrl, "city", a.config.Header, urlValues(opts))
}
