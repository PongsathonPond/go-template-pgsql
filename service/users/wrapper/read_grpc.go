package wrapper

import (
	"context"

	"idev-cms-service/service/users/inout"

	"github.com/opentracing/opentracing-go"
)

func (wrp *Wrapper) ReadGRPC(ctx context.Context, input *inout.ReadInput) (view *inout.UserViewGRPC, err error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "service.users.ReadGRPC")
	defer sp.Finish()

	sp.LogKV("ID", input.ID)

	view, err = wrp.Service.ReadGRPC(ctx, input)

	sp.LogKV("Error", err)

	return view, err
}
