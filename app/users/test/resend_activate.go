package test

import (
	"net/http"

	"github.com/stretchr/testify/mock"
)

func (suite *PackageTestSuite) TestResendActivate() {
	givenInput := suite.makeTestResendActivateInput()
	req, resp, err := suite.makeResendActivateRequest(givenInput)
	suite.NoError(err)

	suite.service.On("ResendActivate", mock.Anything, givenInput).Once().Return(nil)
	suite.router.ServeHTTP(resp, req)

	suite.Equal(http.StatusOK, resp.Code)
}

func (suite *PackageTestSuite) TestResendActivateInvalidJSON() {
	req, resp, err := suite.makeResendActivateRequestInvalidJSON()
	suite.NoError(err)

	suite.service.On("ResendActivate", mock.Anything, nil).Once()
	suite.router.ServeHTTP(resp, req)

	suite.Equal(http.StatusInternalServerError, resp.Code)
}

func (suite *PackageTestSuite) TestResendActivateServiceError() {
	givenInput := suite.makeTestResendActivateInput()
	req, resp, err := suite.makeResendActivateRequest(givenInput)
	suite.NoError(err)

	suite.service.On("ResendActivate", mock.Anything, givenInput).Once().Return(givenError)
	suite.router.ServeHTTP(resp, req)

	suite.Equal(http.StatusInternalServerError, resp.Code)
	suite.Equal("", resp.Header().Get("Content-Location"))
}
