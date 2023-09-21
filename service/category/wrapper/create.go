package wrapper

import (
	"context"
	"github.com/opentracing/opentracing-go"

	"idev-cms-service/service/category/inout"
)

func (wrp *Wrapper) Create(ctx context.Context, input *inout.CategoryCreateInput) (id string, err error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "service.category.Create")
	defer sp.Finish()

	sp.LogKV("Name", input.Name)

	id, err = wrp.Service.Create(ctx, input)

	sp.LogKV("ID", id)
	sp.LogKV("Error", err)

	return id, err
}
