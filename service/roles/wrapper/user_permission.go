package wrapper

import (
	"context"

	"idev-cms-service/service/roles/inout"

	"github.com/opentracing/opentracing-go"
)

func (wrp Wrapper) GetUserPermission(ctx context.Context, token string) (views []*inout.MenuPermissionView, err error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "service.roles.GetUserPermission")
	defer sp.Finish()

	views, err = wrp.Service.GetUserPermission(ctx, token)

	sp.LogKV("Data", views)
	sp.LogKV("Err", err)

	return views, err
}
