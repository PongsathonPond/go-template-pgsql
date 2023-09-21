package wrapper

import (
	"context"

	"idev-cms-service/service/menus/inout"

	"github.com/opentracing/opentracing-go"
)

func (wrp Wrapper) All(ctx context.Context) (total int, items []*inout.MenuView, err error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "service.menus.All")
	defer sp.Finish()

	total, items, err = wrp.Service.All(ctx)

	sp.LogKV("Total", total)
	sp.LogKV("All", items)
	sp.LogKV("Error", err)

	return total, items, err
}
