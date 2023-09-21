package test

import (
	"net/http"

	"github.com/stretchr/testify/mock"
)

func (suite *PackageTestSuite) TestRefreshTokenSuccess() {
	givenInput := suite.makeTestRefreshTokenInput()
	req, resp := suite.makeRefreshTokenRequest(givenInput)

	suite.service.On("RefreshToken", mock.Anything, givenInput).Once().Return(givenRefreshTokenView, nil)
	suite.router.ServeHTTP(resp, req)

	suite.Equal(http.StatusOK, resp.Code)
}

func (suite *PackageTestSuite) TestRefreshTokenError() {
	givenInput := suite.makeTestRefreshTokenInput()
	req, resp := suite.makeRefreshTokenRequest(givenInput)

	suite.service.On("RefreshToken", mock.Anything, givenInput).Once().Return(nil, givenError)
	suite.router.ServeHTTP(resp, req)

	suite.Equal(http.StatusInternalServerError, resp.Code)
}

func (suite *PackageTestSuite) TestRefreshTokenInvalidJSON() {
	req, resp := suite.makeRefreshTokenRequestInvalidJson()

	suite.service.On("RefreshToken", mock.Anything, nil).Once().Return(nil, givenError)
	suite.router.ServeHTTP(resp, req)

	suite.Equal(http.StatusInternalServerError, resp.Code)
}
