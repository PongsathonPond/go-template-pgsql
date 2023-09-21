package implement

import (
	category "idev-cms-service/service/category"
	wrp "idev-cms-service/service/category/wrapper"
	"idev-cms-service/service/util"
	"idev-cms-service/service/util/logs"
	"idev-cms-service/service/validator"
)

type implementation struct {
	*CategoryServiceConfig
}

type CategoryServiceConfig struct {
	Validator    validator.Validator
	RepoCategory util.Repository
	RepoUsers    util.Repository
	RepoRedis    util.RepositoryRedis
	UUID         util.UUID
	DateTime     util.DateTime
	FilterString util.FilterString
	Log          logs.Log
}

func New(config *CategoryServiceConfig) (service category.Service) {
	return &wrp.Wrapper{
		Service: &implementation{config},
	}
}
