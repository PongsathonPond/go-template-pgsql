package wrapper

import (
	"context"

	"idev-cms-service/service/users/inout"

	"github.com/opentracing/opentracing-go"
)

func (wrp *Wrapper) ResendActivate(ctx context.Context, input *inout.ResendActivateInput) (err error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "service.users.ResendActivate")
	defer sp.Finish()

	err = wrp.Service.ResendActivate(ctx, input)

	sp.LogKV("Error", err)

	return err
}
