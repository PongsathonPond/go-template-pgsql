package test

import (
	"idev-cms-service/config"
	"idev-cms-service/service/util/mocks"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	cb "github.com/uniplaces/carbon"
)

type PackageTestSuite struct {
	suite.Suite
	datetime *mocks.DateTime
	config   *config.Config
}

func (suite *PackageTestSuite) SetupSuite() {
	suite.datetime = &mocks.DateTime{}
	suite.config = config.New()
}

func (suite *PackageTestSuite) SetupTest() {
	suite.datetime.On("GetNow").Return(&cb.Carbon{})
	suite.datetime.On("GetUnixNow").Return(int64(0))
	suite.datetime.On("ConvertUnixToDateTime", mock.Anything).Return("")
}
