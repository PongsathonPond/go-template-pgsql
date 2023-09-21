package wrapper

import (
	"idev-cms-service/service/authen"
)

type Wrapper struct {
	Service authen.Service
}

func _(service authen.Service) authen.Service {
	return &Wrapper{service}
}
