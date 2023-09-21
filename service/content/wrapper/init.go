package wrapper

import (
	content "idev-cms-service/service/content"
)

type Wrapper struct {
	Service content.Service
}

func _(service content.Service) content.Service {
	return &Wrapper{service}
}
