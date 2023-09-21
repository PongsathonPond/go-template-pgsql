package wrapper

import (
	"context"

	"idev-cms-service/service/users/inout"

	"github.com/opentracing/opentracing-go"
)

func (wrp *Wrapper) GetMe(ctx context.Context, input *inout.GetMeInput) (view *inout.MeView, err error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "service.users.GetMe")
	defer sp.Finish()

	view, err = wrp.Service.GetMe(ctx, input)

	sp.LogKV("Error", err)

	return view, err
}
