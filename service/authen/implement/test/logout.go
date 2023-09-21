package test

import (
	"idev-cms-service/service/util"

	"github.com/stretchr/testify/mock"
)

func (suite *PackageTestSuite) TestLogoutSuccess() {
	givenInput := suite.makeTestLogoutInput()

	suite.validator.On("Validate", givenInput).Once().Return(nil)
	suite.filterString.On("MakeAccessToken", mock.Anything).Once().Return(givenAccessTokenFilter)
	suite.repoTokens.On("Read", mock.Anything, []string{givenAccessTokenFilter}, givenTokens).Once().Return(nil)
	suite.repoTokens.On("Delete", mock.Anything, []string{givenAccessTokenFilter}).Once().Return(nil)
	suite.repoRedis.On("Clear", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).Once().Return(nil)

	err := suite.service.Logout(suite.ctx, givenInput)

	suite.NoError(err)
}

func (suite *PackageTestSuite) TestLogoutValidateError() {
	givenInput := suite.makeTestLogoutInput()

	suite.validator.On("Validate", givenInput).Once().Return(givenError)

	err := suite.service.Logout(suite.ctx, givenInput)

	suite.Error(err)
	suite.True(util.TypeOfErr(err).IsType(util.ValidationErr))
}

func (suite *PackageTestSuite) TestLogoutReadTokenError() {
	givenInput := suite.makeTestLogoutInput()

	suite.validator.On("Validate", givenInput).Once().Return(nil)
	suite.filterString.On("MakeAccessToken", mock.Anything).Once().Return(givenAccessTokenFilter)
	suite.repoTokens.On("Read", mock.Anything, []string{givenAccessTokenFilter}, givenTokens).Once().Return(givenError)

	err := suite.service.Logout(suite.ctx, givenInput)

	suite.Error(err)
	suite.True(util.TypeOfErr(err).IsType(util.RepoReadErr))
}

func (suite *PackageTestSuite) TestLogoutDeleteTokenError() {
	givenInput := suite.makeTestLogoutInput()

	suite.validator.On("Validate", givenInput).Once().Return(nil)
	suite.filterString.On("MakeAccessToken", mock.Anything).Once().Return(givenAccessTokenFilter)
	suite.repoTokens.On("Read", mock.Anything, []string{givenAccessTokenFilter}, givenTokens).Once().Return(nil)
	suite.repoTokens.On("Delete", mock.Anything, []string{givenAccessTokenFilter}).Once().Return(givenError)

	err := suite.service.Logout(suite.ctx, givenInput)

	suite.Error(err)
	suite.True(util.TypeOfErr(err).IsType(util.RepoDeleteErr))
}
