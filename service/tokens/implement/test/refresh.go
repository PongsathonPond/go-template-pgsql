package test

//func (suite *PackageTestSuite) TestRefreshTokenSuccess() {
//	givenInput := suite.makeTestRefreshTokenInput()
//
//	suite.validator.On("Validate", givenInput).Once().Return(nil)
//	suite.filterString.On("MakeRefreshToken", mock.Anything).Once().Return(givenRefreshTokenFilter)
//	suite.repoTokens.On("Read", mock.Anything, []string{givenRefreshTokenFilter}, givenTokens).Once().Return(nil)
//	suite.filterString.On("MakeID", mock.Anything).Once().Return(givenIDFilter)
//	suite.repoUsers.On("Read", mock.Anything, []string{givenIDFilter}, givenUsers).Once().Return(nil).Run(func(args mock.Arguments) {
//		arg := args[2].(*domain.Users)
//		arg.Status = "active"
//	})
//	suite.filterString.On("MakeID", mock.Anything).Once().Return(givenIDFilter)
//	suite.repoTokens.On("Update", mock.Anything, []string{givenIDFilter}, mock.AnythingOfType("*domain.Tokens")).Once().Return(nil)
//
//	_, err := suite.service.RefreshToken(suite.ctx, givenInput)
//
//	suite.NoError(err)
//}
//
//func (suite *PackageTestSuite) TestRefreshTokenValidateError() {
//	givenInput := suite.makeTestRefreshTokenInput()
//
//	suite.validator.On("Validate", givenInput).Once().Return(givenError)
//
//	_, err := suite.service.RefreshToken(suite.ctx, givenInput)
//
//	suite.Error(err)
//	suite.True(util.TypeOfErr(err).IsType(util.ValidationErr))
//}
//
//func (suite *PackageTestSuite) TestRefreshTokenReadTokenError() {
//	givenInput := suite.makeTestRefreshTokenInput()
//
//	suite.validator.On("Validate", givenInput).Once().Return(nil)
//	suite.filterString.On("MakeRefreshToken", mock.Anything).Once().Return(givenRefreshTokenFilter)
//	suite.repoTokens.On("Read", mock.Anything, []string{givenRefreshTokenFilter}, givenTokens).Once().Return(givenError)
//
//	_, err := suite.service.RefreshToken(suite.ctx, givenInput)
//
//	suite.Error(err)
//	suite.True(util.TypeOfErr(err).IsType(util.RepoReadErr))
//}
//
//func (suite *PackageTestSuite) TestRefreshTokenTokenExpiredError() {
//	givenInput := suite.makeTestRefreshTokenInput()
//
//	suite.validator.On("Validate", givenInput).Once().Return(nil)
//	suite.filterString.On("MakeRefreshToken", mock.Anything).Once().Return(givenRefreshTokenFilter)
//	suite.repoTokens.On("Read", mock.Anything, []string{givenRefreshTokenFilter}, givenTokens).Once().Return(nil).Run(func(args mock.Arguments) {
//		arg := args[2].(*domain.Tokens)
//		arg.RefreshTokenExpiredAt = -1
//	})
//
//	_, err := suite.service.RefreshToken(suite.ctx, givenInput)
//
//	suite.Error(err)
//	suite.True(util.TypeOfErr(err).IsType(util.ValidationTokenExpired))
//}
//
//func (suite *PackageTestSuite) TestRefreshTokenReadUserError() {
//	givenInput := suite.makeTestRefreshTokenInput()
//
//	suite.validator.On("Validate", givenInput).Once().Return(nil)
//	suite.filterString.On("MakeRefreshToken", mock.Anything).Once().Return(givenRefreshTokenFilter)
//	suite.repoTokens.On("Read", mock.Anything, []string{givenRefreshTokenFilter}, givenTokens).Once().Return(nil)
//	suite.filterString.On("MakeID", mock.Anything).Once().Return(givenIDFilter)
//	suite.repoUsers.On("Read", mock.Anything, []string{givenIDFilter}, givenUsers).Once().Return(givenError)
//
//	_, err := suite.service.RefreshToken(suite.ctx, givenInput)
//
//	suite.Error(err)
//	suite.True(util.TypeOfErr(err).IsType(util.RepoReadErr))
//}
//
//func (suite *PackageTestSuite) TestRefreshTokenUserStatusNotActiveError() {
//	givenInput := suite.makeTestRefreshTokenInput()
//
//	suite.validator.On("Validate", givenInput).Once().Return(nil)
//	suite.filterString.On("MakeRefreshToken", mock.Anything).Once().Return(givenRefreshTokenFilter)
//	suite.repoTokens.On("Read", mock.Anything, []string{givenRefreshTokenFilter}, givenTokens).Once().Return(nil)
//	suite.filterString.On("MakeID", mock.Anything).Once().Return(givenIDFilter)
//	suite.repoUsers.On("Read", mock.Anything, []string{givenIDFilter}, givenUsers).Once().Return(nil).Run(func(args mock.Arguments) {
//		arg := args[2].(*domain.Users)
//		arg.Status = "inactive"
//	})
//
//	_, err := suite.service.RefreshToken(suite.ctx, givenInput)
//
//	suite.Error(err)
//	suite.True(util.TypeOfErr(err).IsType(util.RepoReadErr))
//}
//
//func (suite *PackageTestSuite) TestRefreshTokenUpdateTokenError() {
//	givenInput := suite.makeTestRefreshTokenInput()
//
//	suite.validator.On("Validate", givenInput).Once().Return(nil)
//	suite.filterString.On("MakeRefreshToken", mock.Anything).Once().Return(givenRefreshTokenFilter)
//	suite.repoTokens.On("Read", mock.Anything, []string{givenRefreshTokenFilter}, givenTokens).Once().Return(nil)
//	suite.filterString.On("MakeID", mock.Anything).Once().Return(givenIDFilter)
//	suite.repoUsers.On("Read", mock.Anything, []string{givenIDFilter}, givenUsers).Once().Return(nil).Run(func(args mock.Arguments) {
//		arg := args[2].(*domain.Users)
//		arg.Status = "active"
//	})
//	suite.filterString.On("MakeID", mock.Anything).Once().Return(givenIDFilter)
//	suite.repoTokens.On("Update", mock.Anything, []string{givenIDFilter}, mock.AnythingOfType("*domain.Tokens")).Once().Return(givenError)
//
//	_, err := suite.service.RefreshToken(suite.ctx, givenInput)
//
//	suite.Error(err)
//	suite.True(util.TypeOfErr(err).IsType(util.RepoUpdateErr))
//}
