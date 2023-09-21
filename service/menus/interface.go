package roles

import (
	"context"

	"idev-cms-service/service/menus/inout"
)

//go:generate mockery --name=Service
type Service interface {
	All(ctx context.Context) (total int, items []*inout.MenuView, err error)
}
