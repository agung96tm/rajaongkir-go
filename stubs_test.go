package rajaongkir_go

const TestAPIKey = ""
const TestAccountType = AccountStarter
const TestBaseAPI = StarterAPIUrl

var MockConfig = NewConfig(TestAPIKey, TestAccountType)

const TestResponseProvinceSuccess = `
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
`

const TestResponseFail = `
{
    "rajaongkir": {
        "status": {
            "code": 400,
            "description": "Invalid key. API key tidak ditemukan di database RajaOngkir."
        }
    }
}
`

const TestResponseCitySuccess = `
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
`

const TestResponseCostSuccess = `
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
						}
					]
				}
			]
		}
	}
`
