package test

import (
	"context"
	"errors"

	"idev-cms-service/config"
	"idev-cms-service/domain"
	"idev-cms-service/service/tokens"
	"idev-cms-service/service/tokens/implement"
	logMocks "idev-cms-service/service/util/logs/mocks"
	"idev-cms-service/service/util/mocks"
	validatorMocks "idev-cms-service/service/validator/mocks"

	"github.com/stretchr/testify/suite"
	"github.com/uniplaces/carbon"
)

var (
	givenUsers  *domain.Users
	givenTokens *domain.Tokens
	givenError  error

	givenRefreshTokenFilter string
	givenIDFilter           string
	givenAccessTokenFilter  string
	givenUserIDFilter       string
	givenNoItems            interface{}
	givenItems              []interface{}
	givenTotal              int
)

type PackageTestSuite struct {
	suite.Suite
	ctx          context.Context
	validator    *validatorMocks.Validator
	repoUsers    *mocks.Repository
	repoTokens   *mocks.Repository
	repoRedis    *mocks.RepositoryRedis
	config       *config.Config
	datetime     *mocks.DateTime
	filterString *mocks.FilterString
	log          *logMocks.Log
	service      tokens.Service
}

func (suite *PackageTestSuite) SetupSuite() {
	suite.ctx = context.Background()
	suite.validator = &validatorMocks.Validator{}

	suite.repoUsers = &mocks.Repository{}
	suite.repoTokens = &mocks.Repository{}
	suite.repoRedis = &mocks.RepositoryRedis{}

	suite.config = config.New()
	suite.datetime = &mocks.DateTime{}
	suite.filterString = &mocks.FilterString{}
	suite.log = &logMocks.Log{}

	suite.service = implement.New(&implement.TokensServiceConfig{
		Validator:    suite.validator,
		RepoUsers:    suite.repoUsers,
		RepoTokens:   suite.repoTokens,
		RepoRedis:    suite.repoRedis,
		Config:       suite.config,
		DateTime:     suite.datetime,
		FilterString: suite.filterString,
		Log:          suite.log,
	})
}

func (suite *PackageTestSuite) SetupTest() {
	givenUsers = &domain.Users{}
	givenTokens = &domain.Tokens{}
	givenError = errors.New("some error message")
	givenTotal = 0

	suite.datetime.On("GetNow").Return(&carbon.Carbon{})
	suite.datetime.On("GetUnixNow").Return(int64(0))
}
