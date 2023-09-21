package implement

import (
	"idev-cms-service/config"
	"idev-cms-service/service/tokens"
	wrp "idev-cms-service/service/tokens/wrapper"
	"idev-cms-service/service/util"
	"idev-cms-service/service/util/logs"
	"idev-cms-service/service/validator"
)

const (
	PREFIX_AUTH_TOKEN = "auth-token"
)

type implementation struct {
	*TokensServiceConfig
}

type TokensServiceConfig struct {
	Validator    validator.Validator
	RepoUsers    util.Repository
	RepoTokens   util.Repository
	RepoRedis    util.RepositoryRedis
	Config       *config.Config
	DateTime     util.DateTime
	FilterString util.FilterString
	Encrypt      util.Encrypt
	Log          logs.Log
}

func New(config *TokensServiceConfig) (service tokens.Service) {
	return &wrp.Wrapper{
		Service: &implementation{config},
	}
}
