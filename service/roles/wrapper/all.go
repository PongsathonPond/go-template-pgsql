package wrapper

import (
	"context"

	"idev-cms-service/service/roles/inout"

	"github.com/opentracing/opentracing-go"
)

func (wrp Wrapper) All(ctx context.Context, input *inout.RoleAllInput) (items []*inout.RolesAllView, err error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "service.roles.All")
	defer sp.Finish()

	sp.LogKV("Filters", input.Filters)

	items, err = wrp.Service.All(ctx, input)

	sp.LogKV("Items", items)
	sp.LogKV("Error", err)

	return items, err
}
