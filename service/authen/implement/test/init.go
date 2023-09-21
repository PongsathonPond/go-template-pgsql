package test

import (
	"context"
	"errors"

	"idev-cms-service/config"
	"idev-cms-service/domain"
	"idev-cms-service/service/authen"
	"idev-cms-service/service/authen/implement"
	logMocks "idev-cms-service/service/util/logs/mocks"
	"idev-cms-service/service/util/mocks"
	validatorMocks "idev-cms-service/service/validator/mocks"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/uniplaces/carbon"
)

const (
	mockPassword      = "$2a$14$aenXLcqrZrJxJh/dzd/1ROYjvWjKdSIuGF65sIlAgv.Mtr7vakGKq"
	PREFIX_AUTH_TOKEN = "auth-token"
)

var (
	givenUsers      *domain.Users
	givenTokens     *domain.Tokens
	givenError      error
	givenTokenCount int
	givenMD5Token   string
	givenUserID     string

	givenUserNameFilter           string
	givenUserIDFilter             string
	givenUserAgentFilter          string
	givenKeySlugFilter            string
	givenAccessTokenFilter        string
	givenAccessTokenExpAtGTFilter string
	givenStatusInFilter           string
)

type PackageTestSuite struct {
	suite.Suite
	ctx            context.Context
	validator      *validatorMocks.Validator
	repoUsers      *mocks.Repository
	repoUserGroups *mocks.Repository
	repoTokens     *mocks.Repository
	repoRedis      *mocks.RepositoryRedis
	config         *config.Config
	uuid           *mocks.UUID
	datetime       *mocks.DateTime
	filterString   *mocks.FilterString
	encrypt        *mocks.Encrypt
	log            *logMocks.Log
	service        authen.Service
}

func (suite *PackageTestSuite) SetupSuite() {
	suite.ctx = context.Background()
	suite.validator = &validatorMocks.Validator{}

	suite.repoUsers = &mocks.Repository{}
	suite.repoUserGroups = &mocks.Repository{}
	suite.repoTokens = &mocks.Repository{}
	suite.repoRedis = &mocks.RepositoryRedis{}

	suite.config = config.New()
	suite.uuid = &mocks.UUID{}
	suite.datetime = &mocks.DateTime{}
	suite.filterString = &mocks.FilterString{}
	suite.encrypt = &mocks.Encrypt{}
	suite.log = &logMocks.Log{}

	suite.service = implement.New(&implement.AuthenServiceConfig{
		Validator:      suite.validator,
		RepoUsers:      suite.repoUsers,
		RepoUserGroups: suite.repoUserGroups,
		RepoTokens:     suite.repoTokens,
		RepoRedis:      suite.repoRedis,
		UUID:           suite.uuid,
		Config:         suite.config,
		DateTime:       suite.datetime,
		FilterString:   suite.filterString,
		Encrypt:        suite.encrypt,
		Log:            suite.log,
	})
}

func (suite *PackageTestSuite) SetupTest() {
	givenUsers = &domain.Users{}
	givenTokens = &domain.Tokens{}
	givenError = errors.New("some error message")
	givenTokenCount = 1

	suite.datetime.On("GetNow").Return(&carbon.Carbon{})
	suite.datetime.On("GetUnixNow").Return(int64(0))
	suite.datetime.On("ConvertUnixToDateTime", mock.Anything).Return("")
	suite.encrypt.On("MD5", mock.Anything).Return(givenMD5Token)
}
