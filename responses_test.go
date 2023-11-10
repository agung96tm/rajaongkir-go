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
	_ = respProvince.Unmarshal([]byte(TestResponseProvinceSuccess))
	assert.Equal(suite.T(), "Bali", respProvince.Results[0].Province)
	assert.Equal(suite.T(), "1", respProvince.Results[0].ProvinceID)
	assert.Nil(suite.T(), respProvince.StatusValid())
}

func (suite *ResponsesTestSuite) TestCitySuccess() {
	respCity := ResponseCities{}
	_ = respCity.Unmarshal([]byte(TestResponseCitySuccess))
	assert.Equal(suite.T(), "Nanggroe Aceh Darussalam (NAD)", respCity.Results[0].Province)
	assert.Equal(suite.T(), "1", respCity.Results[0].CityID)
	assert.Nil(suite.T(), respCity.StatusValid())
}

func (suite *ResponsesTestSuite) TestCostSuccess() {
	respCost := ResponseCost{}

	_ = respCost.Unmarshal([]byte(TestResponseCostSuccess))
	assert.Equal(suite.T(), "jne", respCost.Results[0].Code)
	assert.Equal(suite.T(), 1, len(respCost.Results[0].Costs))
	assert.Nil(suite.T(), respCost.StatusValid())
}

func TestResponsesTestSuite(t *testing.T) {
	suite.Run(t, new(ResponsesTestSuite))
}
