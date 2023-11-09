package rajaongkir_go

type API struct {
	config Config
}

func NewAPI(config Config) API {
	return API{
		config: config,
	}
}

func (a API) GetProvinces(opts *ProvinceParams) (ResponseProvinces, error) {
	var respProv ResponseProvinces

	respByte, err := get(a.config.BaseUrl, "province", a.config.Header, urlValues(opts))
	if err != nil {
		return respProv, err
	}

	if err := respProv.Unmarshal(respByte); err != nil {
		return respProv, err
	}
	return respProv, nil
}

func (a API) GetCities(opts *CityParams) (ResponseCities, error) {
	var respCities ResponseCities

	respByte, err := get(a.config.BaseUrl, "city", a.config.Header, urlValues(opts))
	if err != nil {
		return respCities, err
	}

	if err := respCities.Unmarshal(respByte); err != nil {
		return respCities, err
	}
	return respCities, nil
}

func (a API) GetCost(opts *CostParams) (ResponseCost, error) {
	var respCost ResponseCost

	respByte, err := post(a.config.BaseUrl, "cost", a.config.Header, urlValues(opts))
	if err != nil {
		return respCost, err
	}

	if err := respCost.Unmarshal(respByte); err != nil {
		return respCost, err
	}

	return respCost, nil
}
