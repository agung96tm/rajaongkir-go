package rajaongkir_go

import (
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/url"
	"testing"
)

type RequestTestSuite struct {
	suite.Suite
}

func (suite *RequestTestSuite) SetupTest() {
	httpmock.Activate()
}

func (suite *RequestTestSuite) TestSendGetRequestSuccess() {
	httpmock.RegisterResponder(
		http.MethodGet,
		MockConfig.BaseUrl+"/province",
		httpmock.NewBytesResponder(http.StatusOK, []byte(TestResponseProvinceSuccess)),
	)

	resp, err := sendGetRequest(MockConfig.BaseUrl+"/province", MockConfig.Header)
	assert.JSONEq(suite.T(), TestResponseProvinceSuccess, string(resp))
	assert.Nil(suite.T(), err, nil)
}

func (suite *RequestTestSuite) TestSendPostRequestSuccess() {
	httpmock.RegisterResponder(
		http.MethodPost,
		MockConfig.BaseUrl+"/cost",
		httpmock.NewBytesResponder(http.StatusOK, []byte(TestResponseCostSuccess)),
	)

	resp, err := sendPostRequest(MockConfig.BaseUrl+"/cost", MockConfig.Header, url.Values{})
	assert.JSONEq(suite.T(), TestResponseCostSuccess, string(resp))
	assert.Nil(suite.T(), err, nil)
}

func (suite *RequestTestSuite) TearDownTest() {
	httpmock.DeactivateAndReset()
}

func TestRequestTestSuite(t *testing.T) {
	suite.Run(t, new(RequestTestSuite))
}
