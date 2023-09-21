package test

import (
	"context"

	"idev-cms-service/config"
	"idev-cms-service/service/util"
	"idev-cms-service/service/util/logs"

	"github.com/stretchr/testify/suite"
)

type PackageTestSuite struct {
	suite.Suite
	ctx      context.Context
	log      logs.Log
	datetime util.DateTime
}

func (suite *PackageTestSuite) SetupSuite() {
	suite.ctx = context.Background()
	suite.log = logs.New(&config.Config{AppTimeZone: "Asia/Bangkok"})
	suite.datetime = util.NewDateTime(&config.Config{AppTimeZone: "Asia/Bangkok"})
}
