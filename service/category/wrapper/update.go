package wrapper

import (
	"context"
	"github.com/opentracing/opentracing-go"

	"idev-cms-service/service/category/inout"
)

func (wrp *Wrapper) Update(ctx context.Context, input *inout.CategoryUpdateInput) (err error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "service.category.Update")
	defer sp.Finish()

	sp.LogKV("ID", input.ID)
	sp.LogKV("Name", input.Name)

	err = wrp.Service.Update(ctx, input)

	sp.LogKV("Error", err)

	return err
}
