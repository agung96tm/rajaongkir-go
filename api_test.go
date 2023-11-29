package rajaongkir_go

import (
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

type APITestSuite struct {
	suite.Suite
	testable *API
}

func (suite *APITestSuite) SetupTest() {
	api := NewAPI(MockConfig)
	suite.testable = &api

	httpmock.Activate()
}

func (suite *APITestSuite) TestGetProvinceSuccess() {
	httpmock.RegisterResponder(
		http.MethodGet,
		MockConfig.BaseUrl+"/province",
		httpmock.NewBytesResponder(http.StatusOK, []byte(TestResponseProvinceSuccess)),
	)

	resp, err := suite.testable.GetProvinces(nil)

	assert.Equal(suite.T(), 200, resp.Status.Code)
	assert.NotNil(suite.T(), resp.Results)
	assert.Nil(suite.T(), err)
}

func (suite *APITestSuite) TestGetProvinceFail() {
	httpmock.RegisterResponder(
		http.MethodGet,
		MockConfig.BaseUrl+"/province",
		httpmock.NewBytesResponder(http.StatusBadRequest, []byte(TestResponseFail)),
	)

	resp, err := suite.testable.GetProvinces(nil)

	assert.Equal(suite.T(), 400, resp.Status.Code)
	assert.Nil(suite.T(), resp.Results)
	assert.NotNil(suite.T(), err)
}

func (suite *APITestSuite) TestGetCitySuccess() {
	httpmock.RegisterResponder(
		http.MethodGet,
		MockConfig.BaseUrl+"/city",
		httpmock.NewBytesResponder(http.StatusOK, []byte(TestResponseProvinceSuccess)),
	)

	resp, err := suite.testable.GetCities(nil)

	assert.Equal(suite.T(), 200, resp.Status.Code)
	assert.NotNil(suite.T(), resp.Results)
	assert.Nil(suite.T(), err)
}

func (suite *APITestSuite) TestGetCityFail() {
	httpmock.RegisterResponder(
		http.MethodGet,
		MockConfig.BaseUrl+"/city",
		httpmock.NewBytesResponder(http.StatusBadRequest, []byte(TestResponseFail)),
	)

	resp, err := suite.testable.GetCities(nil)

	assert.Equal(suite.T(), 400, resp.Status.Code)
	assert.Nil(suite.T(), resp.Results)
	assert.NotNil(suite.T(), err)
}

func (suite *APITestSuite) TestGetCostSuccess() {
	httpmock.RegisterResponder(
		http.MethodPost,
		MockConfig.BaseUrl+"/cost",
		httpmock.NewBytesResponder(http.StatusOK, []byte(TestResponseProvinceSuccess)),
	)

	resp, err := suite.testable.GetCost(&CostParams{
		Origin:      "1",
		Destination: "1",
		Weight:      1,
		Courier:     JNE,
	})

	assert.Equal(suite.T(), 200, resp.Status.Code)
	assert.NotNil(suite.T(), resp.Results)
	assert.Nil(suite.T(), err)
}

func (suite *APITestSuite) TestGetCostFail() {
	httpmock.RegisterResponder(
		http.MethodPost,
		MockConfig.BaseUrl+"/cost",
		httpmock.NewBytesResponder(http.StatusOK, []byte(TestResponseFail)),
	)

	resp, err := suite.testable.GetCost(&CostParams{
		Origin:      "1",
		Destination: "1",
		Weight:      1,
		Courier:     JNE,
	})

	assert.Equal(suite.T(), 400, resp.Status.Code)
	assert.Nil(suite.T(), resp.Results)
	assert.NotNil(suite.T(), err)
}

func (suite *APITestSuite) TearDownTest() {
	httpmock.DeactivateAndReset()
}

func TestAPITestSuite(t *testing.T) {
	suite.Run(t, new(APITestSuite))
}
