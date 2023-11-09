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
	FooBar string `query:"foo_bar"`
	Bar    string `query:"bar"`
}

type HelperTestSuite struct {
	suite.Suite
}

func (suite *HelperTestSuite) TestScanSuccess() {
	testableScan := []scanTest{
		{
			i: MockHelperScan{
				FooBar: "foo bar",
				Bar:    "bar",
			},
			predefined: url.Values{},
			expected: url.Values{
				"foo_bar": {"foo bar"},
				"bar":     {"bar"},
			},
		},
	}

	for i, tt := range testableScan {
		result := scan(tt.i, tt.predefined)
		if !reflect.DeepEqual(tt.expected, result) {
			suite.FailNow("test #%d: result differs from expected value\n", i)
		}
	}
}

func TestHelperTestSuite(t *testing.T) {
	suite.Run(t, new(HelperTestSuite))
}
