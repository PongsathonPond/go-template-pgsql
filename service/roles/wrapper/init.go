package wrapper

import (
	"idev-cms-service/service/roles"
)

type Wrapper struct {
	Service roles.Service
}

func _(service roles.Service) roles.Service {
	return &Wrapper{service}
}
