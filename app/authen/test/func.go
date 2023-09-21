package test

import (
	"idev-cms-service/service/authen/inout"
)

func (suite *PackageTestSuite) makeTestLoginInput() (input *inout.LoginInput) {
	return &inout.LoginInput{
		Username: "test",
		Password: "test",
	}
}

func (suite *PackageTestSuite) makeTestLogoutInput() (input *inout.LogoutInput) {
	return &inout.LogoutInput{}
}
