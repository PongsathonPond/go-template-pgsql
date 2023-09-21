package test

import (
	"net/http"

	"github.com/stretchr/testify/mock"
)

func (suite *PackageTestSuite) TestRevokeTokenSuccess() {
	givenInput := suite.makeTestRevokeTokenInput()
	req, resp := suite.makeRevokeTokenRequest(givenInput)

	suite.service.On("RevokeToken", mock.Anything, givenInput).Once().Return(nil)
	suite.router.ServeHTTP(resp, req)

	suite.Equal(http.StatusOK, resp.Code)
}

func (suite *PackageTestSuite) TestRevokeTokenError() {
	givenInput := suite.makeTestRevokeTokenInput()
	req, resp := suite.makeRevokeTokenRequest(givenInput)

	suite.service.On("RevokeToken", mock.Anything, givenInput).Once().Return(givenError)
	suite.router.ServeHTTP(resp, req)

	suite.Equal(http.StatusInternalServerError, resp.Code)
}
