package test

import (
	"net/http"

	"github.com/stretchr/testify/mock"
)

func (suite *PackageTestSuite) TestCreate() {
	givenInput := suite.makeTestCreateInput()
	req, resp, err := suite.makeCreateRequest(givenInput)
	suite.NoError(err)

	newID := "test"
	suite.service.On("Create", mock.Anything, givenInput).Once().Return(newID, nil)
	suite.router.ServeHTTP(resp, req)

	suite.Equal(http.StatusCreated, resp.Code)
	suite.Equal(newID, resp.Header().Get("Content-Location"))
}

func (suite *PackageTestSuite) TestCreateInvalidJSON() {
	req, resp, err := suite.makeCreateRequestInvalidJSON()
	suite.NoError(err)

	suite.service.On("Create", mock.Anything, nil).Once()
	suite.router.ServeHTTP(resp, req)

	suite.Equal(http.StatusInternalServerError, resp.Code)
}

func (suite *PackageTestSuite) TestCreateServiceError() {
	givenInput := suite.makeTestCreateInput()
	req, resp, err := suite.makeCreateRequest(givenInput)
	suite.NoError(err)

	suite.service.On("Create", mock.Anything, givenInput).Once().Return("", givenError)
	suite.router.ServeHTTP(resp, req)

	suite.Equal(http.StatusInternalServerError, resp.Code)
	suite.Equal("", resp.Header().Get("Content-Location"))
}
