package test

import (
	"errors"
	"net/http/httptest"

	"idev-cms-service/app/tokens"
	"idev-cms-service/service/tokens/inout"
	"idev-cms-service/service/tokens/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

var (
	givenRefreshTokenView *inout.RefreshTokenView
	givenError            error
)

type PackageTestSuite struct {
	suite.Suite
	router  *gin.Engine
	ctx     *gin.Context
	ctrl    *tokens.Controller
	service *mocks.Service
}

func (suite *PackageTestSuite) SetupSuite() {
	suite.service = &mocks.Service{}
	suite.ctrl = tokens.New(suite.service)
	suite.ctx, suite.router = gin.CreateTestContext(httptest.NewRecorder())

	suite.router.POST("/token/refresh", suite.ctrl.Refresh)
	suite.router.POST("/token/revoke", suite.ctrl.Revoke)
}

func (suite *PackageTestSuite) SetupTest() {
	givenRefreshTokenView = &inout.RefreshTokenView{}
	givenError = errors.New("some error message")
}
