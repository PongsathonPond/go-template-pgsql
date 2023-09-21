package test

import (
	"errors"
	"net/http/httptest"

	"idev-cms-service/app/users"
	"idev-cms-service/service/users/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

var (
	givenError error
	givenTest  string
)

type PackageTestSuite struct {
	suite.Suite
	router  *gin.Engine
	ctx     *gin.Context
	ctrl    *users.Controller
	service *mocks.Service
}

func (suite *PackageTestSuite) SetupSuite() {
	suite.service = &mocks.Service{}
	suite.ctrl = users.New(suite.service)
	suite.ctx, suite.router = gin.CreateTestContext(httptest.NewRecorder())

	suite.router.GET("/users", suite.ctrl.List)
	suite.router.POST("/users", suite.ctrl.Create)
	suite.router.GET("/users/:id", suite.ctrl.Read)
	suite.router.PUT("/users/:id", suite.ctrl.Update)
	suite.router.DELETE("/users/:id", suite.ctrl.Delete)

	suite.router.GET("/me", suite.ctrl.GetMe)
	suite.router.PUT("/me", suite.ctrl.UpdateMe)

	suite.router.POST("/me/activate", suite.ctrl.Activate)
	suite.router.POST("/activate/resend", suite.ctrl.ResendActivate)
}

func (suite *PackageTestSuite) SetupTest() {
	givenError = errors.New("some error message")
	givenTest = "test"
}
