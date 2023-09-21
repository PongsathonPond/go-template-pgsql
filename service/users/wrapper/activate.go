package wrapper

import (
	"context"

	"idev-cms-service/service/users/inout"

	"github.com/opentracing/opentracing-go"
)

func (wrp *Wrapper) Activate(ctx context.Context, input *inout.UserActivateInput) (err error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "service.users.Activate")
	defer sp.Finish()

	sp.LogKV("ActivateToken", input.ActivateToken)
	sp.LogKV("Password", input.Password)

	err = wrp.Service.Activate(ctx, input)

	sp.LogKV("Error", err)

	return err
}
