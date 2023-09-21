package test

import (
	"net/http"

	"github.com/stretchr/testify/mock"
)

func (suite *PackageTestSuite) TestLogin() {
	givenInput := suite.makeTestLoginInput()
	req, resp := suite.makeLoginRequest(givenInput)

	suite.service.On("Login", mock.Anything, givenInput).Once().Return(givenViewWithToken, nil)
	suite.router.ServeHTTP(resp, req)

	suite.Equal(http.StatusOK, resp.Code)
}

func (suite *PackageTestSuite) TestLoginError() {
	givenInput := suite.makeTestLoginInput()
	req, resp := suite.makeLoginRequest(givenInput)

	suite.service.On("Login", mock.Anything, givenInput).Once().Return(nil, givenError)
	suite.router.ServeHTTP(resp, req)

	suite.Equal(http.StatusInternalServerError, resp.Code)
}

func (suite *PackageTestSuite) TestLoginInvalidJSON() {
	req, resp := suite.makeLoginRequestInvalidJson()

	suite.service.On("Login", mock.Anything, nil).Once().Return(nil, givenError)
	suite.router.ServeHTTP(resp, req)

	suite.Equal(http.StatusInternalServerError, resp.Code)
}
