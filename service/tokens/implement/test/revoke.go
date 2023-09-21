package test

//func (suite *PackageTestSuite) TestRevokeTokenSuccessWithNoItems() {
//	givenInput := suite.makeTestRevokeTokenInput()
//
//	suite.validator.On("Validate", givenInput).Once().Return(nil)
//	suite.filterString.On("MakeAccessToken", mock.Anything).Once().Return(givenAccessTokenFilter)
//	suite.repoTokens.On("Read", mock.Anything, []string{givenAccessTokenFilter}, givenTokens).Once().Return(nil)
//	suite.filterString.On("MakeUserID", mock.Anything).Once().Return(givenUserIDFilter)
//	suite.repoTokens.On("Find", mock.Anything, []string{givenUserIDFilter}, mock.Anything, givenTokens).Once().Return(givenTotal, givenNoItems, nil)
//
//	err := suite.service.RevokeToken(suite.ctx, givenInput)
//
//	suite.NoError(err)
//}
//
//func (suite *PackageTestSuite) TestRevokeTokenSuccessWithItems() {
//	givenInput := suite.makeTestRevokeTokenInput()
//	givenItems = []interface{}{
//		&domain.Tokens{},
//	}
//
//	suite.validator.On("Validate", givenInput).Once().Return(nil)
//	suite.filterString.On("MakeAccessToken", mock.Anything).Once().Return(givenAccessTokenFilter)
//	suite.repoTokens.On("Read", mock.Anything, []string{givenAccessTokenFilter}, givenTokens).Once().Return(nil)
//	suite.filterString.On("MakeUserID", mock.Anything).Once().Return(givenUserIDFilter)
//	suite.repoTokens.On("Find", mock.Anything, []string{givenUserIDFilter}, mock.Anything, givenTokens).Once().Return(givenTotal, givenItems, nil)
//	suite.filterString.On("MakeID", mock.Anything).Once().Return(givenIDFilter)
//	suite.repoTokens.On("Delete", mock.Anything, []string{givenIDFilter}).Once().Return(nil)
//	suite.repoRedis.On("Clear", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).Once().Return(nil)
//
//	err := suite.service.RevokeToken(suite.ctx, givenInput)
//
//	suite.NoError(err)
//}
//
//func (suite *PackageTestSuite) TestRevokeTokenSuccessWithItemsDeleteTokenError() {
//	givenInput := suite.makeTestRevokeTokenInput()
//	givenItems = []interface{}{
//		&domain.Tokens{},
//	}
//
//	suite.validator.On("Validate", givenInput).Once().Return(nil)
//	suite.filterString.On("MakeAccessToken", mock.Anything).Once().Return(givenAccessTokenFilter)
//	suite.repoTokens.On("Read", mock.Anything, []string{givenAccessTokenFilter}, givenTokens).Once().Return(nil)
//	suite.filterString.On("MakeUserID", mock.Anything).Once().Return(givenUserIDFilter)
//	suite.repoTokens.On("Find", mock.Anything, []string{givenUserIDFilter}, mock.Anything, givenTokens).Once().Return(givenTotal, givenItems, nil)
//	suite.filterString.On("MakeID", mock.Anything).Once().Return(givenIDFilter)
//	suite.repoTokens.On("Delete", mock.Anything, []string{givenIDFilter}).Once().Return(givenError)
//
//	err := suite.service.RevokeToken(suite.ctx, givenInput)
//
//	suite.Error(err)
//	suite.True(util.TypeOfErr(err).IsType(util.RepoDeleteErr))
//}
//
//func (suite *PackageTestSuite) TestRevokeTokenValidateError() {
//	givenInput := suite.makeTestRevokeTokenInput()
//
//	suite.validator.On("Validate", givenInput).Once().Return(givenError)
//
//	err := suite.service.RevokeToken(suite.ctx, givenInput)
//
//	suite.Error(err)
//	suite.True(util.TypeOfErr(err).IsType(util.ValidationErr))
//}
//
//func (suite *PackageTestSuite) TestRevokeTokenReadTokenError() {
//	givenInput := suite.makeTestRevokeTokenInput()
//
//	suite.validator.On("Validate", givenInput).Once().Return(nil)
//	suite.filterString.On("MakeAccessToken", mock.Anything).Once().Return(givenAccessTokenFilter)
//	suite.repoTokens.On("Read", mock.Anything, []string{givenAccessTokenFilter}, givenTokens).Once().Return(givenError)
//
//	err := suite.service.RevokeToken(suite.ctx, givenInput)
//
//	suite.Error(err)
//	suite.True(util.TypeOfErr(err).IsType(util.RepoReadErr))
//}
//
//func (suite *PackageTestSuite) TestRevokeTokenFindTokenError() {
//	givenInput := suite.makeTestRevokeTokenInput()
//
//	suite.validator.On("Validate", givenInput).Once().Return(nil)
//	suite.filterString.On("MakeAccessToken", mock.Anything).Once().Return(givenAccessTokenFilter)
//	suite.repoTokens.On("Read", mock.Anything, []string{givenAccessTokenFilter}, givenTokens).Once().Return(nil)
//	suite.filterString.On("MakeUserID", mock.Anything).Once().Return(givenUserIDFilter)
//	suite.repoTokens.On("Find", mock.Anything, []string{givenUserIDFilter}, mock.Anything, givenTokens).Once().Return(givenTotal, givenNoItems, givenError)
//
//	err := suite.service.RevokeToken(suite.ctx, givenInput)
//
//	suite.Error(err)
//	suite.True(util.TypeOfErr(err).IsType(util.RepoListErr))
//}
