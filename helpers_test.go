package rajaongkir_go

import (
	"github.com/stretchr/testify/suite"
	"net/url"
	"reflect"
	"testing"
)

type scanTest struct {
	i          any
	predefined url.Values
	expected   url.Values
}

type MockHelperScan struct {
	MockInt   int      `query:"mock_int"`
	MockStr   string   `query:"mock_str"`
	MockFloat float64  `query:"mock_float"`
	MockBool  bool     `query:"mock_bool"`
	MockArr   []string `query:"mock_arr"`
}

type HelperTestSuite struct {
	suite.Suite
}

func (suite *HelperTestSuite) TestScanSuccess() {
	testableScan := []scanTest{
		{
			i: MockHelperScan{
				MockInt:   666,
				MockStr:   "bar",
				MockFloat: 10.2,
				MockBool:  true,
				MockArr:   []string{"one", "two", "three"},
			},
			predefined: url.Values{},
			expected: url.Values{
				"mock_int":   {"666"},
				"mock_str":   {"bar"},
				"mock_float": {"10.2"},
				"mock_bool":  {"true"},
				"mock_arr":   {`["one","two","three"]`},
			},
		},
	}

	for i, tt := range testableScan {
		result := scan(tt.i, tt.predefined)
		if !reflect.DeepEqual(tt.expected, result) {
			suite.T().Fatalf("test #%d: result differs from expected value\n", i)
		}
	}
}

func TestHelperTestSuite(t *testing.T) {
	suite.Run(t, new(HelperTestSuite))
}
