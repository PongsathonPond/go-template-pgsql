package wrapper

import (
	"context"
	"github.com/opentracing/opentracing-go"

	"idev-cms-service/domain"
	"idev-cms-service/service/category/inout"
)

func (wrp *Wrapper) List(ctx context.Context, opt *domain.PageOption) (total int, list []*inout.CategoryView, err error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "service.category.List")
	defer sp.Finish()

	sp.LogKV("Page", opt.Page)
	sp.LogKV("PerPage", opt.PerPage)
	sp.LogKV("Filters", opt.Filters)

	total, list, err = wrp.Service.List(ctx, opt)

	sp.LogKV("Total", total)
	sp.LogKV("List", list)
	sp.LogKV("Error", err)

	return total, list, err
}
