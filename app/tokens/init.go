package tokens

import (
	"idev-cms-service/service/tokens"
)

type Controller struct {
	service tokens.Service
}

func New(service tokens.Service) (tokens *Controller) {
	return &Controller{service}
}
