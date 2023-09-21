package test

import (
	"net/http"

	"idev-cms-service/service/users/inout"

	"github.com/stretchr/testify/mock"
)

func (suite *PackageTestSuite) TestRead() {
	givenInput := suite.makeTestReadInput()
	req, resp, err := suite.makeReadRequest(givenInput)
	suite.NoError(err)

	givenView := &inout.UserReadView{}

	suite.service.On("Read", mock.Anything, givenInput).Once().Return(givenView, nil)
	suite.router.ServeHTTP(resp, req)

	suite.Equal(http.StatusOK, resp.Code)
}

func (suite *PackageTestSuite) TestReadServiceError() {
	givenInput := suite.makeTestReadInput()
	req, resp, err := suite.makeReadRequest(givenInput)
	suite.NoError(err)

	suite.service.On("Read", mock.Anything, givenInput).Once().Return(nil, givenError)
	suite.router.ServeHTTP(resp, req)

	suite.Equal(http.StatusInternalServerError, resp.Code)
}
