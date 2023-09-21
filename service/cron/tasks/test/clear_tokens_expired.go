package test

import (
	"github.com/stretchr/testify/mock"
)

func (suite *PackageTestSuite) TestClearTokenExpiredSuccess() {
	suite.filterString.On("MakeRefreshTokenExpiredAtLessThan", mock.Anything).Once().Return(givenRFTokenExpLTFilter)
	suite.repoTokens.On("Count", mock.Anything, []string{givenRFTokenExpLTFilter}).Once().Return(givenTotal, nil)
	suite.repoTokens.On("DeleteMany", mock.Anything, []string{givenRFTokenExpLTFilter}).Once().Return(nil)

	suite.service.ClearTokensExpired()
}

func (suite *PackageTestSuite) TestClearTokenExpiredCountTokenError() {
	suite.filterString.On("MakeRefreshTokenExpiredAtLessThan", mock.Anything).Once().Return(givenRFTokenExpLTFilter)
	suite.repoTokens.On("Count", mock.Anything, []string{givenRFTokenExpLTFilter}).Once().Return(0, givenError)
	suite.log.On("Error", givenError).Once()

	suite.service.ClearTokensExpired()
}

func (suite *PackageTestSuite) TestNothingClearTokenExpired() {
	givenTotal = 0
	suite.filterString.On("MakeRefreshTokenExpiredAtLessThan", mock.Anything).Once().Return(givenRFTokenExpLTFilter)
	suite.repoTokens.On("Count", mock.Anything, []string{givenRFTokenExpLTFilter}).Once().Return(givenTotal, nil)

	suite.service.ClearTokensExpired()
}

func (suite *PackageTestSuite) TestClearTokenExpiredDeleteTokenError() {
	suite.filterString.On("MakeRefreshTokenExpiredAtLessThan", mock.Anything).Once().Return(givenRFTokenExpLTFilter)
	suite.repoTokens.On("Count", mock.Anything, []string{givenRFTokenExpLTFilter}).Once().Return(givenTotal, nil)
	suite.repoTokens.On("DeleteMany", mock.Anything, []string{givenRFTokenExpLTFilter}).Once().Return(givenError)
	suite.log.On("Error", givenError).Once()

	suite.service.ClearTokensExpired()
}
