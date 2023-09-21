package wrapper

import (
	"context"

	"idev-cms-service/service/users/inout"

	"github.com/opentracing/opentracing-go"
)

func (wrp *Wrapper) UserChangePassword(ctx context.Context, input *inout.UserChangePasswordInput) (err error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "service.users.UserChangePassword")
	defer sp.Finish()

	sp.LogKV("OldPassword", input.OldPassword)
	sp.LogKV("NewPassword", input.NewPassword)

	err = wrp.Service.UserChangePassword(ctx, input)

	sp.LogKV("Error", err)

	return err
}
