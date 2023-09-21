package users

import "idev-cms-service/service/users"

type Controller struct {
	service users.Service
}

func New(userSvc users.Service) (user *Controller) {
	return &Controller{service: userSvc}
}
