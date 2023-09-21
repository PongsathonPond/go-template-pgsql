package test

func (suite *PackageTestSuite) TestTokenInputToDomain() {
	givenInput := suite.makeTestTokenInput()

	actual := givenInput.ToDomain(suite.datetime, suite.config)

	suite.Equal(givenInput.ID, actual.ID)
	suite.Equal(givenInput.UserID, actual.UserID)
	suite.Equal(givenInput.UserAgent, actual.UserAgent)
	suite.Equal(givenInput.AccessToken, actual.AccessToken)
	suite.Equal(givenInput.RefreshToken, actual.RefreshToken)
}

func (suite *PackageTestSuite) TestTokenUpdateToDomain() {
	givenInput := suite.makeTestTokenUpdate()

	actual := givenInput.ToDomain(suite.datetime, suite.config)

	suite.Equal(givenInput.AccessToken, actual.AccessToken)
	suite.Equal(givenInput.RefreshToken, actual.RefreshToken)
}
