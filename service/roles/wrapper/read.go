package wrapper

import (
	"context"

	"idev-cms-service/service/roles/inout"

	"github.com/opentracing/opentracing-go"
)

func (wrp Wrapper) Read(ctx context.Context, input *inout.RoleReadInput) (view *inout.RoleReadView, err error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "service.roles.Read")
	defer sp.Finish()

	sp.LogKV("ID", input.ID)

	view, err = wrp.Service.Read(ctx, input)

	sp.LogKV("Data", view)
	sp.LogKV("Err", err)

	return view, err
}
