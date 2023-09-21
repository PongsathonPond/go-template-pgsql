package implement

import (
	"idev-cms-service/config"
	"idev-cms-service/service/authen"
	wrp "idev-cms-service/service/authen/wrapper"
	"idev-cms-service/service/util"
	"idev-cms-service/service/util/logs"
	"idev-cms-service/service/validator"
)

const (
	PREFIX_AUTH_TOKEN = "auth-token"
)

type implementation struct {
	*AuthenServiceConfig
}

type AuthenServiceConfig struct {
	Validator      validator.Validator
	RepoUsers      util.Repository
	RepoUserGroups util.Repository
	RepoTokens     util.Repository
	RepoRedis      util.RepositoryRedis
	UUID           util.UUID
	Config         *config.Config
	DateTime       util.DateTime
	FilterString   util.FilterString
	Encrypt        util.Encrypt
	Log            logs.Log
}

func New(config *AuthenServiceConfig) (service authen.Service) {
	return &wrp.Wrapper{
		Service: &implementation{config},
	}
}
