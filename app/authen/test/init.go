package test

import (
	"errors"
	"net/http/httptest"

	"idev-cms-service/app/authen"
	"idev-cms-service/service/authen/inout"
	"idev-cms-service/service/authen/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

var (
	givenViewWithToken *inout.UserViewWithToken
	givenError         error
)

type PackageTestSuite struct {
	suite.Suite
	router  *gin.Engine
	ctx     *gin.Context
	ctrl    *authen.Controller
	service *mocks.Service
}

func (suite *PackageTestSuite) SetupSuite() {
	suite.service = &mocks.Service{}
	suite.ctrl = authen.New(suite.service)
	suite.ctx, suite.router = gin.CreateTestContext(httptest.NewRecorder())

	suite.router.POST("/login", suite.ctrl.Login)
	suite.router.POST("/logout", suite.ctrl.Logout)
}

func (suite *PackageTestSuite) SetupTest() {
	givenViewWithToken = &inout.UserViewWithToken{}
	givenError = errors.New("some error message")
}
