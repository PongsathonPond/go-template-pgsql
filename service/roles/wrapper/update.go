package wrapper

import (
	"context"

	"idev-cms-service/service/roles/inout"

	"github.com/opentracing/opentracing-go"
)

func (wrp Wrapper) Update(ctx context.Context, input *inout.RoleUpdateInput) (err error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "service.roles.Update")
	defer sp.Finish()

	sp.LogKV("ID", input.ID)
	sp.LogKV("Name", input.Name)
	sp.LogKV("Description", input.Description)
	sp.LogKV("Permissions", input.Permissions)
	sp.LogKV("Status", input.Status)

	err = wrp.Service.Update(ctx, input)

	sp.LogKV("Error", err)

	return err
}
