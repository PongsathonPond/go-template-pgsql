package category

import (
	category "idev-cms-service/service/category"
)

type Controller struct {
	service category.Service
}

func New(categorySvc category.Service) (category *Controller) {
	return &Controller{service: categorySvc}
}
