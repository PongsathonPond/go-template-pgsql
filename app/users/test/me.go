package test

import (
	"net/http"

	"idev-cms-service/service/users/inout"

	"github.com/stretchr/testify/mock"
)

func (suite *PackageTestSuite) TestGetMe() {
	givenInput := suite.makeTestGetMeInput()
	req, resp, err := suite.makeGetMeRequest()
	suite.NoError(err)

	givenView := &inout.MeView{}

	suite.service.On("GetMe", mock.Anything, givenInput).Once().Return(givenView, nil)
	suite.router.ServeHTTP(resp, req)

	suite.Equal(http.StatusOK, resp.Code)
}

func (suite *PackageTestSuite) TestGetMeServiceError() {
	givenInput := suite.makeTestGetMeInput()
	req, resp, err := suite.makeGetMeRequest()
	suite.NoError(err)

	givenView := &inout.MeView{}

	suite.service.On("GetMe", mock.Anything, givenInput).Once().Return(givenView, givenError)
	suite.router.ServeHTTP(resp, req)

	suite.Equal(http.StatusInternalServerError, resp.Code)
}
