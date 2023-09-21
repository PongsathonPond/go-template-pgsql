package wrapper

import (
	"context"

	"idev-cms-service/service/tokens/inout"

	"github.com/opentracing/opentracing-go"
)

func (wrp *Wrapper) RevokeToken(ctx context.Context, input *inout.RevokeTokenInput) (err error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "service.authen.Revoke")
	defer sp.Finish()

	sp.LogKV("AccessToken", input.AccessToken)

	err = wrp.Service.RevokeToken(ctx, input)

	sp.LogKV("Error", err)

	return err
}
