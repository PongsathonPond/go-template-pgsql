package wrapper

import (
	roles "idev-cms-service/service/menus"
)

type Wrapper struct {
	Service roles.Service
}

func _(service roles.Service) roles.Service {
	return &Wrapper{service}
}
