package test

import (
	"time"

	"idev-cms-service/service/authen/inout"
)

func (suite *PackageTestSuite) TestUserWithTokenToView() {
	users := suite.makeTestUsers()
	roles := suite.makeTestRoles()
	accessToken := "test"
	refreshToken := "test"

	actual := inout.UserWithTokenToView(users, roles, suite.datetime, suite.config, &accessToken, &refreshToken)

	suite.Equal(users.ID, actual.ID)
	suite.Equal(users.FirstName, actual.FirstName)
	suite.Equal(users.LastName, actual.LastName)
	suite.Equal(users.Email, actual.Email)
	suite.Equal(users.Mobile, actual.Mobile)
	suite.Equal(users.Username, actual.Username)
	suite.Equal(users.Role, actual.Role)
	suite.Equal(users.ImageProfile, actual.ImageProfile)
	suite.Equal(accessToken, actual.AccessToken)
	suite.Equal(refreshToken, actual.RefreshToken)
	suite.Equal(int64((time.Duration(suite.config.AccessTokenTTL) * time.Minute).Seconds()), actual.AccessTokenTTL)
	suite.Equal(int64((time.Duration(suite.config.RefreshTokenTTL) * time.Minute).Seconds()), actual.RefreshTokenTTL)
	suite.Equal(users.Status, actual.Status)
	suite.Equal(suite.datetime.ConvertUnixToDateTime(users.CreatedAt), actual.CreatedAt)
	suite.Equal(suite.datetime.ConvertUnixToDateTime(users.UpdatedAt), actual.UpdatedAt)

	suite.Equal(roles.ID, actual.RoleData.ID)
	suite.Equal(roles.Name, actual.RoleData.Name)
}
