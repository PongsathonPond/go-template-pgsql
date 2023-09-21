package implement

import (
	"idev-cms-service/config"
	menus "idev-cms-service/service/menus"
	"idev-cms-service/service/roles"
	wrp "idev-cms-service/service/roles/wrapper"
	"idev-cms-service/service/users"
	"idev-cms-service/service/util"
	"idev-cms-service/service/util/logs"
	"idev-cms-service/service/validator"
)

const (
	PREFIX_ROLE_PERM = "role-perm"
)

type implementation struct {
	*RolesServiceConfig
}

type RolesServiceConfig struct {
	Validator    validator.Validator
	Repo         util.Repository
	MenusRepo    util.Repository
	UserService  users.Service
	ServiceMenus menus.Service
	UUID         util.UUID
	DateTime     util.DateTime
	Perm         util.Permission
	FilterString util.FilterString
	Log          logs.Log
	Config       *config.Config
}

func New(config *RolesServiceConfig) (service roles.Service) {
	return &wrp.Wrapper{
		Service: &implementation{config},
	}
}

var mapMenu = map[string]map[string]string{
	"authentication": {
		"dashboard": "ga",
		"users":     "users",
		"roles":     "roles-and-permissions",
		"content":   "content",
		"category":  "category",
	},
}
