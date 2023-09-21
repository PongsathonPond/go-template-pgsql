package menus

import (
	menu "idev-cms-service/service/menus"
)

type Controller struct {
	service menu.Service
}

func New(menuSvc menu.Service) (menu *Controller) {
	return &Controller{service: menuSvc}
}
