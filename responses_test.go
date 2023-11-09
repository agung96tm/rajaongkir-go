package rajaongkir_go

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ResponsesTestSuite struct {
	suite.Suite
}

func (suite *ResponsesTestSuite) TestProvinceSuccess() {
	respProvince := ResponseProvinces{}
	_ = respProvince.Unmarshal([]byte(`
		{
			"rajaongkir": {
				"query": {
					"id": "1"
				},
				"status": {
					"code": 200,
					"description": "OK"
				},
				"results": {
					"province_id": "1",
					"province": "Bali"
				}
			}
		}
	`))
	assert.Equal(suite.T(), respProvince.Results[0].Province, "Bali")
	assert.Equal(suite.T(), respProvince.Results[0].ProvinceID, "1")
	assert.Equal(suite.T(), respProvince.StatusValid(), nil)
}

func (suite *ResponsesTestSuite) TestCitySuccess() {
	respCity := ResponseCities{}
	_ = respCity.Unmarshal([]byte(`
		{
			"rajaongkir": {
				"query": {
					"id": "1"
				},
				"status": {
					"code": 200,
					"description": "OK"
				},
				"results": {
					"city_id": "1",
					"province_id": "21",
					"province": "Nanggroe Aceh Darussalam (NAD)",
					"type": "Kabupaten",
					"city_name": "Aceh Barat",
					"postal_code": "23681"
				}
			}
		}
	`))
	assert.Equal(suite.T(), respCity.Results[0].Province, "Nanggroe Aceh Darussalam (NAD)")
	assert.Equal(suite.T(), respCity.Results[0].CityID, "1")
	assert.Equal(suite.T(), respCity.StatusValid(), nil)
}

func (suite *ResponsesTestSuite) TestCostSuccess() {
	respCost := ResponseCost{}

	_ = respCost.Unmarshal([]byte(`
		{
			"rajaongkir": {
				"query": {
					"origin": 501,
					"destination": "114",
					"weight": 1,
					"courier": "jne"
				},
				"status": {
					"code": 200,
					"description": "OK"
				},
				"origin_details": {
					"city_id": "501",
					"province_id": "5",
					"province": "DI Yogyakarta",
					"type": "Kota",
					"city_name": "Yogyakarta",
					"postal_code": "55111"
				},
				"destination_details": {
					"city_id": "114",
					"province_id": "1",
					"province": "Bali",
					"type": "Kota",
					"city_name": "Denpasar",
					"postal_code": "80227"
				},
				"results": [
					{
						"code": "jne",
						"name": "Jalur Nugraha Ekakurir (JNE)",
						"costs": [
							{
								"service": "OKE",
								"description": "Ongkos Kirim Ekonomis",
								"cost": [
									{
										"value": 27000,
										"etd": "4-5",
										"note": ""
									}
								]
							},
							{
								"service": "REG",
								"description": "Layanan Reguler",
								"cost": [
									{
										"value": 31000,
										"etd": "2-3",
										"note": ""
									}
								]
							}
						]
					}
				]
			}
		}
	`))
	assert.Equal(suite.T(), respCost.Results[0].Code, "jne")
	assert.Equal(suite.T(), len(respCost.Results[0].Costs), 2)
	assert.Equal(suite.T(), respCost.StatusValid(), nil)
}

func TestResponsesTestSuite(t *testing.T) {
	suite.Run(t, new(ResponsesTestSuite))
}
