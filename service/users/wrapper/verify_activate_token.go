package wrapper

import (
	"context"

	"idev-cms-service/service/users/inout"

	"github.com/opentracing/opentracing-go"
)

func (wrp *Wrapper) VerifyActivateToken(ctx context.Context, input *inout.VerifyActivateTokenInput) (err error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "service.users.VerifyActivateToken")
	defer sp.Finish()

	sp.LogKV("ActivateToken", input.ActivateToken)

	err = wrp.Service.VerifyActivateToken(ctx, input)

	sp.LogKV("Error", err)

	return err
}
