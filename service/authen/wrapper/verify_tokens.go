package wrapper

import (
	"context"

	"idev-cms-service/service/authen/inout"

	"github.com/opentracing/opentracing-go"
)

func (wrp *Wrapper) VerifyTokens(ctx context.Context, token *string) (resp *inout.VerifyTokenResponse) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "service.authen.VerifyToken")
	defer sp.Finish()

	sp.LogKV("AccessToken", *token)

	resp = wrp.Service.VerifyTokens(ctx, token)

	sp.LogKV("Verify", resp.Verify)
	sp.LogKV("userID", resp.UserID)
	//sp.LogKV("userGroup", resp.UserGroup)
	sp.LogKV("role", resp.Role)

	return resp
}
