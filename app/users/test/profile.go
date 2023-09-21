package test

import (
	"net/http"

	"github.com/stretchr/testify/mock"
)

func (suite *PackageTestSuite) TestUpdateMe() {
	givenInput := suite.makeTestUpdateMeInput()
	req, resp, err := suite.makeUpdateMeRequest(givenInput)
	suite.NoError(err)

	suite.service.On("UpdateMe", mock.Anything, givenInput).Once().Return(nil)
	suite.router.ServeHTTP(resp, req)

	suite.Equal(http.StatusOK, resp.Code)
}

func (suite *PackageTestSuite) TestUpdateMeInvalidJSON() {
	req, resp, err := suite.makeUpdateMeRequestInvalidJSON()
	suite.NoError(err)

	suite.service.On("UpdateMe", mock.Anything, nil).Once()
	suite.router.ServeHTTP(resp, req)

	suite.Equal(http.StatusInternalServerError, resp.Code)
}

func (suite *PackageTestSuite) TestUpdateMeServiceError() {
	givenInput := suite.makeTestUpdateMeInput()
	req, resp, err := suite.makeUpdateMeRequest(givenInput)
	suite.NoError(err)

	suite.service.On("UpdateMe", mock.Anything, givenInput).Once().Return(givenError)
	suite.router.ServeHTTP(resp, req)

	suite.Equal(http.StatusInternalServerError, resp.Code)
}
