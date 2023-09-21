package wrapper

import (
	"context"

	"idev-cms-service/service/users/inout"

	"github.com/opentracing/opentracing-go"
)

func (wrp *Wrapper) UpdatePassword(ctx context.Context, input *inout.UserUpdatePasswordInput) (err error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "service.users.UpdatePassword")
	defer sp.Finish()

	sp.LogKV("Email", input.Email)

	err = wrp.Service.UpdatePassword(ctx, input)

	sp.LogKV("Error", err)

	return err
}
