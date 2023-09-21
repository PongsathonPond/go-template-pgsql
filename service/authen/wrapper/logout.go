package wrapper

import (
	"context"

	"idev-cms-service/service/authen/inout"

	"github.com/opentracing/opentracing-go"
)

func (wrp *Wrapper) Logout(ctx context.Context, input *inout.LogoutInput) (err error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "service.authen.Logout")
	defer sp.Finish()

	sp.LogKV("AccessToken", input.AccessToken)

	err = wrp.Service.Logout(ctx, input)

	sp.LogKV("Error", err)

	return err
}
