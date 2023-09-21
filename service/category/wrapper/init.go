package wrapper

import (
	category "idev-cms-service/service/category"
)

type Wrapper struct {
	Service category.Service
}

func _(service category.Service) category.Service {
	return &Wrapper{service}
}
