package wrapper

import (
	"context"

	"idev-cms-service/service/users/inout"

	"github.com/opentracing/opentracing-go"
)

func (wrp *Wrapper) Read(ctx context.Context, input *inout.ReadInput) (view *inout.UserReadView, err error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "service.users.Read")
	defer sp.Finish()

	sp.LogKV("ID", input.ID)

	view, err = wrp.Service.Read(ctx, input)

	sp.LogKV("Error", err)

	return view, err
}
