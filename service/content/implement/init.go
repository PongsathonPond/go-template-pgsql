package implement

import (
	content "idev-cms-service/service/content"
	wrp "idev-cms-service/service/content/wrapper"
	"idev-cms-service/service/tokens"
	"idev-cms-service/service/util"
	"idev-cms-service/service/util/logs"
	"idev-cms-service/service/validator"
)

type implementation struct {
	*ContentServiceConfig
}

type ContentServiceConfig struct {
	Validator     validator.Validator
	RepoContent   util.Repository
	RepoUsers     util.Repository
	RepoCategory  util.Repository
	RepoRedis     util.RepositoryRedis
	UUID          util.UUID
	DateTime      util.DateTime
	FilterString  util.FilterString
	Log           logs.Log
	ServiceTokens tokens.Service
}

func New(config *ContentServiceConfig) (service content.Service) {
	return &wrp.Wrapper{
		Service: &implementation{config},
	}
}
