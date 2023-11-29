package rajaongkir_go

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/url"
	"testing"
)

type ParamsTestSuite struct {
	suite.Suite
}

func (suite *ParamsTestSuite) TestProvinceParamsSuccess() {
	urlValues := urlValues(ProvinceParams{ID: "1"})
	assert.Equal(suite.T(), url.Values{"id": {"1"}}, urlValues)
}

func (suite *ParamsTestSuite) TestCityParamsSuccess() {
	urlValues := urlValues(CityParams{ID: "1", ProvinceID: "1"})
	assert.Equal(suite.T(), url.Values{"id": {"1"}, "province": {"1"}}, urlValues)
}

func (suite *ParamsTestSuite) TestCostParamsSuccess() {
	urlValues := urlValues(CostParams{
		Origin:      "1",
		Destination: "1",
		Weight:      1,
		Courier:     JNE,
	})

	assert.Equal(
		suite.T(),
		url.Values{
			"origin":      {"1"},
			"destination": {"1"},
			"weight":      {"1"},
			"courier":     {"jne"},
		},
		urlValues,
	)
}

func (suite *ParamsTestSuite) TestCourierParamsSuccess() {
	result := CourierParams("JNE", "TIKI")
	assert.Equal(suite.T(), "JNE:TIKI", result)
}

func TestParamsTestSuite(t *testing.T) {
	suite.Run(t, new(ParamsTestSuite))
}
