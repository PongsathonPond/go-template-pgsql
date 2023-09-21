package content

import (
	content "idev-cms-service/service/content"
)

type Controller struct {
	service content.Service
}

func New(contentSvc content.Service) (content *Controller) {
	return &Controller{service: contentSvc}
}
