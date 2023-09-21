package test

import (
	"net/http"

	"github.com/stretchr/testify/mock"
)

func (suite *PackageTestSuite) TestLogout() {
	givenInput := suite.makeTestLogoutInput()
	req, resp := suite.makeLogoutRequest(givenInput)

	suite.service.On("Logout", mock.Anything, givenInput).Once().Return(nil)
	suite.router.ServeHTTP(resp, req)

	suite.Equal(http.StatusOK, resp.Code)
}

func (suite *PackageTestSuite) TestLogoutError() {
	givenInput := suite.makeTestLogoutInput()
	req, resp := suite.makeLogoutRequest(givenInput)

	suite.service.On("Logout", mock.Anything, givenInput).Once().Return(givenError)
	suite.router.ServeHTTP(resp, req)

	suite.Equal(http.StatusInternalServerError, resp.Code)
}
