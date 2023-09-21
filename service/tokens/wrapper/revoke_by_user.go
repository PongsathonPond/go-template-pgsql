package wrapper

import (
	"context"

	"github.com/opentracing/opentracing-go"
)

func (wrp *Wrapper) RevokeTokenByUserId(ctx context.Context, userId *string) (err error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "service.authen.RevokeByUserId")
	defer sp.Finish()

	sp.LogKV("UserID", userId)

	err = wrp.Service.RevokeTokenByUserId(ctx, userId)

	sp.LogKV("Error", err)

	return err
}
