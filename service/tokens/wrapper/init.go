package wrapper

import (
	"idev-cms-service/service/tokens"
)

type Wrapper struct {
	Service tokens.Service
}

func _(service tokens.Service) tokens.Service {
	return &Wrapper{service}
}
