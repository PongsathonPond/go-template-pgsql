package wrapper

import (
	"context"

	"idev-cms-service/service/roles/inout"

	"github.com/opentracing/opentracing-go"
)

func (wrp Wrapper) Create(ctx context.Context, input *inout.RoleCreateInput) (id string, err error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "service.roles.Create")
	defer sp.Finish()

	sp.LogKV("Name", input.Name)

	id, err = wrp.Service.Create(ctx, input)

	sp.LogKV("ID", id)
	sp.LogKV("Error", err)

	return id, err
}
