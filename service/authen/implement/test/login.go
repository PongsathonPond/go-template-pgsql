package test

//func (suite *PackageTestSuite) TestLoginWithUpdateOldTokenSuccess() {
//	givenInput := suite.makeTestLoginInput()
//
//	suite.validator.On("Validate", givenInput).Once().Return(nil)
//	suite.filterString.On("MakeUserName", mock.Anything).Once().Return(givenUserNameFilter)
//	suite.filterString.On("MakeStatusIn", mock.Anything, mock.Anything).Once().Return(givenStatusInFilter)
//	suite.repoUsers.On("Read", mock.Anything, []string{givenUserNameFilter, givenStatusInFilter}, givenUsers).Once().Return(nil).Run(func(args mock.Arguments) {
//		arg := args[2].(*domain.Users)
//		arg.Password = mockPassword
//		arg.Status = "active"
//	})
//	suite.encrypt.On("CheckPasswordHash", givenInput.Password, mockPassword).Once().Return(true)
//	suite.filterString.On("MakeUserID", mock.Anything).Once().Return(givenUserIDFilter)
//	suite.filterString.On("MakeUserAgent", mock.Anything).Once().Return(givenUserAgentFilter)
//	suite.repoTokens.On("Count", mock.Anything, []string{givenUserIDFilter, givenUserAgentFilter}).Once().Return(givenTokenCount, nil)
//	suite.repoTokens.On("Update", mock.Anything, mock.AnythingOfType("[]string"), mock.AnythingOfType("*domain.Tokens")).Once().Return(nil)
//	suite.repoRedis.On("SetWithTTL", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string"), 10*time.Minute).Once().Return(nil)
//	suite.filterString.On("MakeKeySlug", mock.Anything).Once().Return(givenKeySlugFilter)
//
//	_, err := suite.service.Login(suite.ctx, givenInput)
//
//	suite.NoError(err)
//}
//
//func (suite *PackageTestSuite) TestLoginWithAddNewTokenSuccess() {
//	givenInput := suite.makeTestLoginInput()
//	givenTokenCount = 0
//
//	suite.validator.On("Validate", givenInput).Once().Return(nil)
//	suite.filterString.On("MakeUserName", mock.Anything).Once().Return(givenUserNameFilter)
//	suite.filterString.On("MakeStatusIn", mock.Anything, mock.Anything).Once().Return(givenStatusInFilter)
//	suite.repoUsers.On("Read", mock.Anything, []string{givenUserNameFilter, givenStatusInFilter}, givenUsers).Once().Return(nil).Run(func(args mock.Arguments) {
//		arg := args[2].(*domain.Users)
//		arg.Password = mockPassword
//		arg.Status = "active"
//	})
//	suite.encrypt.On("CheckPasswordHash", givenInput.Password, mockPassword).Once().Return(true)
//	suite.filterString.On("MakeUserID", mock.Anything).Once().Return(givenUserIDFilter)
//	suite.filterString.On("MakeUserAgent", mock.Anything).Once().Return(givenUserAgentFilter)
//	suite.repoTokens.On("Count", mock.Anything, []string{givenUserIDFilter, givenUserAgentFilter}).Once().Return(givenTokenCount, nil)
//	suite.uuid.On("Generate").Once().Return("")
//	suite.repoTokens.On("Create", mock.Anything, mock.AnythingOfType("*domain.Tokens")).Once().Return("", nil)
//	suite.repoRedis.On("SetWithTTL", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string"), 10*time.Minute).Once().Return(nil)
//	suite.filterString.On("MakeKeySlug", mock.Anything).Once().Return(givenKeySlugFilter)
//
//	_, err := suite.service.Login(suite.ctx, givenInput)
//
//	suite.NoError(err)
//}
//
//func (suite *PackageTestSuite) TestLoginWithValidationError() {
//	givenInput := suite.makeTestLoginInput()
//
//	suite.validator.On("Validate", givenInput).Once().Return(givenError)
//
//	_, err := suite.service.Login(suite.ctx, givenInput)
//
//	suite.Error(err)
//	suite.True(util.TypeOfErr(err).IsType(util.ValidationErr))
//}
//
//func (suite *PackageTestSuite) TestLoginReadUserError() {
//	givenInput := suite.makeTestLoginInput()
//
//	suite.validator.On("Validate", givenInput).Once().Return(nil)
//	suite.filterString.On("MakeUserName", mock.Anything).Once().Return(givenUserNameFilter)
//	suite.filterString.On("MakeStatusIn", mock.Anything, mock.Anything).Once().Return(givenStatusInFilter)
//	suite.repoUsers.On("Read", mock.Anything, []string{givenUserNameFilter, givenStatusInFilter}, givenUsers).Once().Return(givenError).Run(func(args mock.Arguments) {
//		arg := args[2].(*domain.Users)
//		arg.Password = mockPassword
//		arg.Status = "active"
//	})
//
//	_, err := suite.service.Login(suite.ctx, givenInput)
//
//	suite.Error(err)
//	suite.True(util.TypeOfErr(err).IsType(util.RepoReadErr))
//}
//
//func (suite *PackageTestSuite) TestLoginPasswordNotMatchError() {
//	givenInput := suite.makeTestLoginInput()
//
//	suite.validator.On("Validate", givenInput).Once().Return(nil)
//	suite.filterString.On("MakeUserName", mock.Anything).Once().Return(givenUserNameFilter)
//	suite.filterString.On("MakeStatusIn", mock.Anything, mock.Anything).Once().Return(givenStatusInFilter)
//	suite.repoUsers.On("Read", mock.Anything, []string{givenUserNameFilter, givenStatusInFilter}, givenUsers).Once().Return(nil).Run(func(args mock.Arguments) {
//		arg := args[2].(*domain.Users)
//		arg.Password = "unknown"
//		arg.Status = "active"
//	})
//	suite.encrypt.On("CheckPasswordHash", givenInput.Password, "unknown").Once().Return(false)
//
//	_, err := suite.service.Login(suite.ctx, givenInput)
//
//	suite.Error(err)
//	suite.True(util.TypeOfErr(err).IsType(util.RepoReadErr))
//}
//
//func (suite *PackageTestSuite) TestLoginUserStatusNotEqualActiveError() {
//	givenInput := suite.makeTestLoginInput()
//
//	suite.validator.On("Validate", givenInput).Once().Return(nil)
//	suite.filterString.On("MakeUserName", mock.Anything).Once().Return(givenUserNameFilter)
//	suite.filterString.On("MakeStatusIn", mock.Anything, mock.Anything).Once().Return(givenStatusInFilter)
//	suite.repoUsers.On("Read", mock.Anything, []string{givenUserNameFilter, givenStatusInFilter}, givenUsers).Once().Return(nil).Run(func(args mock.Arguments) {
//		arg := args[2].(*domain.Users)
//		arg.Password = mockPassword
//		arg.Status = "inactive"
//	})
//	suite.encrypt.On("CheckPasswordHash", givenInput.Password, mockPassword).Once().Return(true)
//
//	_, err := suite.service.Login(suite.ctx, givenInput)
//
//	suite.Error(err)
//	suite.True(util.TypeOfErr(err).IsType(util.RepoReadErr))
//}
//
//func (suite *PackageTestSuite) TestLoginGenerateTokensError() {
//	givenInput := suite.makeTestLoginInput()
//	givenTokenCount = 0
//
//	suite.validator.On("Validate", givenInput).Once().Return(nil)
//	suite.filterString.On("MakeUserName", mock.Anything).Once().Return(givenUserNameFilter)
//	suite.filterString.On("MakeStatusIn", mock.Anything, mock.Anything).Once().Return(givenStatusInFilter)
//	suite.repoUsers.On("Read", mock.Anything, []string{givenUserNameFilter, givenStatusInFilter}, givenUsers).Once().Return(nil).Run(func(args mock.Arguments) {
//		arg := args[2].(*domain.Users)
//		arg.Password = mockPassword
//		arg.Status = "active"
//	})
//	suite.encrypt.On("CheckPasswordHash", givenInput.Password, mockPassword).Once().Return(true)
//	suite.filterString.On("MakeUserID", mock.Anything).Once().Return(givenUserIDFilter)
//	suite.filterString.On("MakeUserAgent", mock.Anything).Once().Return(givenUserAgentFilter)
//	suite.repoTokens.On("Count", mock.Anything, []string{givenUserIDFilter, givenUserAgentFilter}).Once().Return(givenTokenCount, givenError)
//
//	_, err := suite.service.Login(suite.ctx, givenInput)
//
//	suite.Error(err)
//	suite.True(util.TypeOfErr(err).IsType(util.RepoReadErr))
//}
//
//func (suite *PackageTestSuite) TestLoginUpdateTokenError() {
//	givenInput := suite.makeTestLoginInput()
//
//	suite.validator.On("Validate", givenInput).Once().Return(nil)
//	suite.filterString.On("MakeUserName", mock.Anything).Once().Return(givenUserNameFilter)
//	suite.filterString.On("MakeStatusIn", mock.Anything, mock.Anything).Once().Return(givenStatusInFilter)
//	suite.repoUsers.On("Read", mock.Anything, []string{givenUserNameFilter, givenStatusInFilter}, givenUsers).Once().Return(nil).Run(func(args mock.Arguments) {
//		arg := args[2].(*domain.Users)
//		arg.Password = mockPassword
//		arg.Status = "active"
//	})
//	suite.encrypt.On("CheckPasswordHash", givenInput.Password, mockPassword).Once().Return(true)
//	suite.filterString.On("MakeUserID", mock.Anything).Once().Return(givenUserIDFilter)
//	suite.filterString.On("MakeUserAgent", mock.Anything).Once().Return(givenUserAgentFilter)
//	suite.repoTokens.On("Count", mock.Anything, []string{givenUserIDFilter, givenUserAgentFilter}).Once().Return(givenTokenCount, nil)
//	suite.repoTokens.On("Update", mock.Anything, mock.AnythingOfType("[]string"), mock.AnythingOfType("*domain.Tokens")).Once().Return(givenError)
//
//	_, err := suite.service.Login(suite.ctx, givenInput)
//
//	suite.Error(err)
//	suite.True(util.TypeOfErr(err).IsType(util.RepoUpdateErr))
//}
//
//func (suite *PackageTestSuite) TestLoginCreateTokenError() {
//	givenInput := suite.makeTestLoginInput()
//	givenTokenCount = 0
//
//	suite.validator.On("Validate", givenInput).Once().Return(nil)
//	suite.filterString.On("MakeUserName", mock.Anything).Once().Return(givenUserNameFilter)
//	suite.filterString.On("MakeStatusIn", mock.Anything, mock.Anything).Once().Return(givenStatusInFilter)
//	suite.repoUsers.On("Read", mock.Anything, []string{givenUserNameFilter, givenStatusInFilter}, givenUsers).Once().Return(nil).Run(func(args mock.Arguments) {
//		arg := args[2].(*domain.Users)
//		arg.Password = mockPassword
//		arg.Status = "active"
//	})
//	suite.encrypt.On("CheckPasswordHash", givenInput.Password, mockPassword).Once().Return(true)
//	suite.filterString.On("MakeUserID", mock.Anything).Once().Return(givenUserIDFilter)
//	suite.filterString.On("MakeUserAgent", mock.Anything).Once().Return(givenUserAgentFilter)
//	suite.repoTokens.On("Count", mock.Anything, []string{givenUserIDFilter, givenUserAgentFilter}).Once().Return(givenTokenCount, nil)
//	suite.uuid.On("Generate").Once().Return("")
//	suite.repoTokens.On("Create", mock.Anything, mock.AnythingOfType("*domain.Tokens")).Once().Return("", givenError)
//
//	_, err := suite.service.Login(suite.ctx, givenInput)
//
//	suite.Error(err)
//	suite.True(util.TypeOfErr(err).IsType(util.RepoCreateErr))
//}
//
//func (suite *PackageTestSuite) TestLoginSetRedisError() {
//	givenInput := suite.makeTestLoginInput()
//
//	suite.validator.On("Validate", givenInput).Once().Return(nil)
//	suite.filterString.On("MakeUserName", mock.Anything).Once().Return(givenUserNameFilter)
//	suite.filterString.On("MakeStatusIn", mock.Anything, mock.Anything).Once().Return(givenStatusInFilter)
//	suite.repoUsers.On("Read", mock.Anything, []string{givenUserNameFilter, givenStatusInFilter}, givenUsers).Once().Return(nil).Run(func(args mock.Arguments) {
//		arg := args[2].(*domain.Users)
//		arg.Password = mockPassword
//		arg.Status = "active"
//	})
//	suite.encrypt.On("CheckPasswordHash", givenInput.Password, mockPassword).Once().Return(true)
//	suite.filterString.On("MakeUserID", mock.Anything).Once().Return(givenUserIDFilter)
//	suite.filterString.On("MakeUserAgent", mock.Anything).Once().Return(givenUserAgentFilter)
//	suite.repoTokens.On("Count", mock.Anything, []string{givenUserIDFilter, givenUserAgentFilter}).Once().Return(givenTokenCount, nil)
//	suite.repoTokens.On("Update", mock.Anything, mock.AnythingOfType("[]string"), mock.AnythingOfType("*domain.Tokens")).Once().Return(nil)
//	suite.repoRedis.On("SetWithTTL", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string"), 10*time.Minute).Once().Return(givenError)
//	suite.log.On("Error", givenError)
//	suite.filterString.On("MakeKeySlug", mock.Anything).Once().Return(givenKeySlugFilter)
//
//	_, err := suite.service.Login(suite.ctx, givenInput)
//
//	suite.NoError(err)
//}
//
//func (suite *PackageTestSuite) TestLoginReadUserGroupsError() {
//	givenInput := suite.makeTestLoginInput()
//
//	suite.validator.On("Validate", givenInput).Once().Return(nil)
//	suite.filterString.On("MakeUserName", mock.Anything).Once().Return(givenUserNameFilter)
//	suite.filterString.On("MakeStatusIn", mock.Anything, mock.Anything).Once().Return(givenStatusInFilter)
//	suite.repoUsers.On("Read", mock.Anything, []string{givenUserNameFilter, givenStatusInFilter}, givenUsers).Once().Return(nil).Run(func(args mock.Arguments) {
//		arg := args[2].(*domain.Users)
//		arg.Password = mockPassword
//		arg.Status = "active"
//	})
//	suite.encrypt.On("CheckPasswordHash", givenInput.Password, mockPassword).Once().Return(true)
//	suite.filterString.On("MakeUserID", mock.Anything).Once().Return(givenUserIDFilter)
//	suite.filterString.On("MakeUserAgent", mock.Anything).Once().Return(givenUserAgentFilter)
//	suite.repoTokens.On("Count", mock.Anything, []string{givenUserIDFilter, givenUserAgentFilter}).Once().Return(givenTokenCount, nil)
//	suite.repoTokens.On("Update", mock.Anything, mock.AnythingOfType("[]string"), mock.AnythingOfType("*domain.Tokens")).Once().Return(nil)
//	suite.repoRedis.On("SetWithTTL", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string"), 10*time.Minute).Once().Return(nil)
//	suite.filterString.On("MakeKeySlug", mock.Anything).Once().Return(givenKeySlugFilter)
//
//	_, err := suite.service.Login(suite.ctx, givenInput)
//
//	suite.Error(err)
//	suite.True(util.TypeOfErr(err).IsType(util.RepoReadErr))
//}
//
//func (suite *PackageTestSuite) TestLoginGetRoleGRPCAuthorError() {
//	givenInput := suite.makeTestLoginInput()
//
//	suite.validator.On("Validate", givenInput).Once().Return(nil)
//	suite.filterString.On("MakeUserName", mock.Anything).Once().Return(givenUserNameFilter)
//	suite.filterString.On("MakeStatusIn", mock.Anything, mock.Anything).Once().Return(givenStatusInFilter)
//	suite.repoUsers.On("Read", mock.Anything, []string{givenUserNameFilter, givenStatusInFilter}, givenUsers).Once().Return(nil).Run(func(args mock.Arguments) {
//		arg := args[2].(*domain.Users)
//		arg.Password = mockPassword
//		arg.Status = "active"
//	})
//	suite.encrypt.On("CheckPasswordHash", givenInput.Password, mockPassword).Once().Return(true)
//	suite.filterString.On("MakeUserID", mock.Anything).Once().Return(givenUserIDFilter)
//	suite.filterString.On("MakeUserAgent", mock.Anything).Once().Return(givenUserAgentFilter)
//	suite.repoTokens.On("Count", mock.Anything, []string{givenUserIDFilter, givenUserAgentFilter}).Once().Return(givenTokenCount, nil)
//	suite.repoTokens.On("Update", mock.Anything, mock.AnythingOfType("[]string"), mock.AnythingOfType("*domain.Tokens")).Once().Return(nil)
//	suite.repoRedis.On("SetWithTTL", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string"), 10*time.Minute).Once().Return(nil)
//	suite.filterString.On("MakeKeySlug", mock.Anything).Once().Return(givenKeySlugFilter)
//
//	_, err := suite.service.Login(suite.ctx, givenInput)
//
//	suite.Error(err)
//	suite.True(util.TypeOfErr(err).IsType(util.ServiceGRPCErr))
//}
