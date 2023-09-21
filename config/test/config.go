package test

import (
	"errors"

	"idev-cms-service/config"
)

func (suite *PackageTestSuite) TestGetConfig() {
	conf := config.New()
	if conf == nil {
		suite.NoError(errors.New("Cannot get config value"))
	}
}
