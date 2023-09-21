package wrapper

import (
	"context"

	"idev-cms-service/service/users/inout"

	"github.com/opentracing/opentracing-go"
)

func (wrp *Wrapper) SetVerify(ctx context.Context, input *inout.SetVerifyInput) (resp *inout.SetVerifyResponse, err error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "service.users.SetVerify")
	defer sp.Finish()

	sp.LogKV("Token", input.Token)
	sp.LogKV("Email", input.Email)

	resp, err = wrp.Service.SetVerify(ctx, input)

	if resp != nil {
		sp.LogKV("UserID", resp.ID)
		sp.LogKV("Email", resp.Email)
		sp.LogKV("Verified", resp.Verified)
	}
	sp.LogKV("Error", err)

	return resp, err
}
