package authen

import (
	"idev-cms-service/service/authen"
)

type Controller struct {
	service authen.Service
}

func New(service authen.Service) (users *Controller) {
	return &Controller{service}
}
