package middleware

import (
	"idev-cms-service/service/authen"
	"idev-cms-service/service/roles"
	"idev-cms-service/service/users"
)

type Service struct {
	*MiddlewareServiceConfig
}

type MiddlewareServiceConfig struct {
	AuthenService authen.Service
	UserService   users.Service
	RoleService   roles.Service
}

func New(config *MiddlewareServiceConfig) Service {
	return Service{config}
}
