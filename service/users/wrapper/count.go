package wrapper

import (
	"context"

	"idev-cms-service/service/users/inout"

	"github.com/opentracing/opentracing-go"
)

func (wrp *Wrapper) CountByRole(ctx context.Context, input *inout.CountByRoleInput) (total int, err error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "service.users.CountByRole")
	defer sp.Finish()

	sp.LogKV("RoleId", input.RoleId)

	total, err = wrp.Service.CountByRole(ctx, input)

	sp.LogKV("Error", err)

	return total, err
}
