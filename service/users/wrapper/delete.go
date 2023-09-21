package wrapper

import (
	"context"

	"idev-cms-service/service/users/inout"

	"github.com/opentracing/opentracing-go"
)

func (wrp *Wrapper) Delete(ctx context.Context, input *inout.UserDeleteInput) (err error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "service.users.Delete")
	defer sp.Finish()

	sp.LogKV("ID", input.ID)

	err = wrp.Service.Delete(ctx, input)

	sp.LogKV("Error", err)

	return err
}
