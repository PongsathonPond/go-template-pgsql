package wrapper

import (
	"idev-cms-service/service/users"
)

type Wrapper struct {
	Service users.Service
}

func _(service users.Service) users.Service {
	return &Wrapper{service}
}
