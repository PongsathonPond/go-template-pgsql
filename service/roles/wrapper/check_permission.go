package wrapper

import (
	"context"

	"idev-cms-service/service/roles/inout"

	"github.com/opentracing/opentracing-go"
)

func (wrp Wrapper) CheckPermission(ctx context.Context, input *inout.CheckPermissionInput) (isPermission bool) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "service.roles.CheckPermission")
	defer sp.Finish()

	sp.LogKV("RoleID", input.RoleID)
	sp.LogKV("Method", input.Method)
	sp.LogKV("KeySubMenus", input.EndPoint)
	sp.LogKV("Service", input.Service)

	isPermission = wrp.Service.CheckPermission(ctx, input)

	sp.LogKV("IsPermission", isPermission)

	return isPermission
}
