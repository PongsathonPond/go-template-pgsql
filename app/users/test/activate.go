package test

import (
	"net/http"

	"github.com/stretchr/testify/mock"
)

func (suite *PackageTestSuite) TestActivate() {
	givenInput := suite.makeTestActivateInput()
	req, resp, err := suite.makeActivateRequest(givenInput)
	suite.NoError(err)

	suite.service.On("Activate", mock.Anything, givenInput).Once().Return(nil)
	suite.router.ServeHTTP(resp, req)

	suite.Equal(http.StatusOK, resp.Code)
}

func (suite *PackageTestSuite) TestActivateInvalidJSON() {
	req, resp, err := suite.makeActivateRequestInvalidJSON()
	suite.NoError(err)

	suite.service.On("Activate", mock.Anything, nil).Once()
	suite.router.ServeHTTP(resp, req)

	suite.Equal(http.StatusInternalServerError, resp.Code)
}

func (suite *PackageTestSuite) TestActivateServiceError() {
	givenInput := suite.makeTestActivateInput()
	req, resp, err := suite.makeActivateRequest(givenInput)
	suite.NoError(err)

	suite.service.On("Activate", mock.Anything, givenInput).Once().Return(givenError)
	suite.router.ServeHTTP(resp, req)

	suite.Equal(http.StatusInternalServerError, resp.Code)
	suite.Equal("", resp.Header().Get("Content-Location"))
}
