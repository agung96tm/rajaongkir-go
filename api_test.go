package rajaongkir_go

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type APITestSuite struct {
	suite.Suite
}

func TestAPITestSuite(t *testing.T) {
	suite.Run(t, new(APITestSuite))
}
