package wrapper

import (
	"context"

	"github.com/opentracing/opentracing-go"
)

func (wrp *Wrapper) CheckEmailExist(ctx context.Context, email *string) (err error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "service.users.CheckEmailExist")
	defer sp.Finish()

	sp.LogKV("Email", email)

	err = wrp.Service.CheckEmailExist(ctx, email)

	sp.LogKV("Error", err)

	return err
}
