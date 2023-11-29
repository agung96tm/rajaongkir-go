package rajaongkir_go

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ConfigTestSuite struct {
	suite.Suite
	testable *Config
}

func (suite *ConfigTestSuite) SetupTest() {
	config := NewConfig(TestAPIKey, TestAccountType)
	suite.testable = &config
}

func (suite *ConfigTestSuite) TestConfigValue() {
	assert.Equal(suite.T(), suite.testable.BaseUrl, TestBaseAPI)
	assert.Equal(suite.T(), suite.testable.AccountKey, TestAPIKey)
	assert.Equal(suite.T(), suite.testable.AccountType, TestAccountType)
}

func (suite *ConfigTestSuite) TestGetAPISuccess() {
	typeTest := []struct {
		accountType string
		url         string
	}{
		{AccountStarter, StarterAPIUrl},
		{AccountPro, ProAPIUrl},
		{AccountBasic, BasicAPIUrl},
		{"wrong", StarterAPIUrl},
	}

	for _, tt := range typeTest {
		result := GetApiURL(tt.accountType)
		assert.Equal(suite.T(), result, tt.url)
	}
}

func TestConfigTestSuite(t *testing.T) {
	suite.Run(t, new(ConfigTestSuite))
}
