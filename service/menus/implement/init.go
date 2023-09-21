package implement

import (
	menu "idev-cms-service/service/menus"
	wrp "idev-cms-service/service/menus/wrapper"
	"idev-cms-service/service/util"
	"idev-cms-service/service/util/logs"
	"idev-cms-service/service/validator"
)

const (
	Depa = "depa"
)

type implementation struct {
	*MenusServiceConfig
}

type MenusServiceConfig struct {
	Validator    validator.Validator
	Repo         util.Repository
	RepoRoles    util.Repository
	UUID         util.UUID
	DateTime     util.DateTime
	Perm         util.Permission
	FilterString util.FilterString
	Log          logs.Log
}

func New(config *MenusServiceConfig) (service menu.Service) {
	return &wrp.Wrapper{
		Service: &implementation{config},
	}
}
