package test

import (
	"idev-cms-service/service/tokens/inout"
)

func (suite *PackageTestSuite) makeTestRefreshTokenInput() (input *inout.RefreshTokenInput) {
	return &inout.RefreshTokenInput{
		RefreshToken: "test",
	}
}

func (suite *PackageTestSuite) makeTestRevokeTokenInput() (input *inout.RevokeTokenInput) {
	return &inout.RevokeTokenInput{
		AccessToken: "test",
	}
}
