package roles

import (
	"idev-cms-service/service/roles"
)

type Controller struct {
	service roles.Service
}

func New(roleSvc roles.Service) (role *Controller) {
	return &Controller{service: roleSvc}
}
