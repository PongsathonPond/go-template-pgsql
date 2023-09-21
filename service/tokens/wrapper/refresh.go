package wrapper

import (
	"context"

	"idev-cms-service/service/tokens/inout"

	"github.com/opentracing/opentracing-go"
)

func (wrp *Wrapper) RefreshToken(ctx context.Context, input *inout.RefreshTokenInput) (view *inout.RefreshTokenView, err error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "service.authen.Refresh")
	defer sp.Finish()

	sp.LogKV("Refresh", input.RefreshToken)

	view, err = wrp.Service.RefreshToken(ctx, input)

	if view != nil {
		sp.LogKV("AccessToken", view.AccessToken)
		sp.LogKV("AccessTokenTTL", view.AccessTokenTTL)
	}
	sp.LogKV("Error", err)

	return view, err
}
