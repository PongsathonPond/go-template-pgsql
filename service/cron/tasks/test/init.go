package test

import (
	"context"
	"errors"

	"idev-cms-service/service/cron/tasks"
	logMocks "idev-cms-service/service/util/logs/mocks"
	"idev-cms-service/service/util/mocks"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var (
	givenRFTokenExpLTFilter string
	givenTotal              int
	givenError              error
)

type PackageTestSuite struct {
	suite.Suite
	ctx          context.Context
	repoTokens   *mocks.Repository
	datetime     *mocks.DateTime
	filterString *mocks.FilterString
	log          *logMocks.Log
	service      *tasks.Tasks
}

func (suite *PackageTestSuite) SetupSuite() {
	suite.ctx = context.Background()
	suite.repoTokens = &mocks.Repository{}

	suite.datetime = &mocks.DateTime{}
	suite.filterString = &mocks.FilterString{}
	suite.log = &logMocks.Log{}

	suite.service = tasks.New(&tasks.CronTasksConfig{
		RepoTokens:   suite.repoTokens,
		FilterString: suite.filterString,
		DateTime:     suite.datetime,
		Log:          suite.log,
	})
}

func (suite *PackageTestSuite) SetupTest() {
	givenTotal = 1
	givenError = errors.New("some error message")

	suite.log.On("Info", mock.Anything)
	suite.datetime.On("GetUnixNow").Return(int64(0))
}
