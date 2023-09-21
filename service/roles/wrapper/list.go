package wrapper

import (
	"context"

	"idev-cms-service/domain"
	"idev-cms-service/service/roles/inout"

	"github.com/opentracing/opentracing-go"
)

func (wrp Wrapper) List(ctx context.Context, opt *domain.PageOption) (total int, items []*inout.RolesListView, err error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "service.userGroups.All")
	defer sp.Finish()

	sp.LogKV("Page", opt.Page)
	sp.LogKV("PerPage", opt.PerPage)
	sp.LogKV("Filters", opt.Filters)

	total, items, err = wrp.Service.List(ctx, opt)

	sp.LogKV("Total", total)
	sp.LogKV("All", items)
	sp.LogKV("Error", err)

	return total, items, err
}
