package test

//func (suite *PackageTestSuite) TestVerifyTokensSuccess() {
//	var givenInput string
//	givenUserIDFilter = "test"
//	givenKeyRedis := fmt.Sprintf("%s-%s", PREFIX_AUTH_TOKEN, givenInput)
//
//	suite.repoRedis.On("Get", mock.Anything, givenKeyRedis, &givenUserID).Once().Return(nil).Run(func(args mock.Arguments) {
//		arg := args[2].(*string)
//		*arg = "test"
//	})
//	suite.filterString.On("MakeID", mock.Anything).Once().Return(givenUserIDFilter)
//	suite.repoUsers.On("Read", mock.Anything, []string{givenUserIDFilter}, givenUsers).Once().Return(nil)
//
//	resp := suite.service.VerifyTokens(suite.ctx, &givenInput)
//
//	suite.NotEqual(&inout.VerifyTokenResponse{}, resp)
//}
//
//func (suite *PackageTestSuite) TestVerifyTokensUserIDEmptyError() {
//	var givenInput string
//	givenKeyRedis := fmt.Sprintf("%s-%s", PREFIX_AUTH_TOKEN, givenInput)
//
//	suite.repoRedis.On("Get", mock.Anything, givenKeyRedis, &givenUserID).Once().Return(nil)
//
//	resp := suite.service.VerifyTokens(suite.ctx, &givenInput)
//
//	suite.Equal(&inout.VerifyTokenResponse{}, resp)
//}
//
//func (suite *PackageTestSuite) TestVerifyTokensUserGetRedisError() {
//	var givenInput string
//	givenKeyRedis := fmt.Sprintf("%s-%s", PREFIX_AUTH_TOKEN, givenInput)
//	givenError = redis.Nil
//
//	suite.repoRedis.On("Get", mock.Anything, givenKeyRedis, &givenUserID).Once().Return(givenError)
//	suite.filterString.On("MakeAccessToken", mock.Anything).Once().Return(givenAccessTokenFilter)
//	suite.filterString.On("MakeAccessTokenExpiredAtGreaterThan", mock.Anything).Once().Return(givenAccessTokenExpAtGTFilter)
//	suite.repoTokens.On("Read", mock.Anything, []string{givenAccessTokenFilter, givenAccessTokenExpAtGTFilter}, givenTokens).Once().Return(nil)
//	suite.repoRedis.On("SetWithTTL", mock.Anything, givenKeyRedis, givenTokens.UserID, mock.Anything).Once().Return(nil)
//
//	resp := suite.service.VerifyTokens(suite.ctx, &givenInput)
//
//	suite.Equal(&inout.VerifyTokenResponse{}, resp)
//}
//
//func (suite *PackageTestSuite) TestVerifyTokensUserGetRedisAndReadTokenError() {
//	var givenInput string
//	givenKeyRedis := fmt.Sprintf("%s-%s", PREFIX_AUTH_TOKEN, givenInput)
//	givenError = redis.Nil
//
//	suite.repoRedis.On("Get", mock.Anything, givenKeyRedis, &givenUserID).Once().Return(givenError)
//	suite.filterString.On("MakeAccessToken", mock.Anything).Once().Return(givenAccessTokenFilter)
//	suite.filterString.On("MakeAccessTokenExpiredAtGreaterThan", mock.Anything).Once().Return(givenAccessTokenExpAtGTFilter)
//	suite.repoTokens.On("Read", mock.Anything, []string{givenAccessTokenFilter, givenAccessTokenExpAtGTFilter}, givenTokens).Once().Return(givenError)
//
//	resp := suite.service.VerifyTokens(suite.ctx, &givenInput)
//
//	suite.Equal(&inout.VerifyTokenResponse{}, resp)
//}
//
//func (suite *PackageTestSuite) TestVerifyTokensReadUserError() {
//	var givenInput string
//	givenUserIDFilter = "test"
//	givenKeyRedis := fmt.Sprintf("%s-%s", PREFIX_AUTH_TOKEN, givenInput)
//
//	suite.repoRedis.On("Get", mock.Anything, givenKeyRedis, &givenUserID).Once().Return(nil).Run(func(args mock.Arguments) {
//		arg := args[2].(*string)
//		*arg = "test"
//	})
//	suite.filterString.On("MakeID", mock.Anything).Once().Return(givenUserIDFilter)
//	suite.repoUsers.On("Read", mock.Anything, []string{givenUserIDFilter}, givenUsers).Once().Return(givenError)
//
//	resp := suite.service.VerifyTokens(suite.ctx, &givenInput)
//
//	suite.Equal(&inout.VerifyTokenResponse{Verify: true, UserID: "test"}, resp)
//}
