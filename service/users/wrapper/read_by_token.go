package wrapper

import (
	"context"

	"idev-cms-service/service/users/inout"

	"github.com/opentracing/opentracing-go"
)

func (wrp *Wrapper) ReadByToken(ctx context.Context, token *string) (view *inout.UserViewByToken, err error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "service.users.Read")
	defer sp.Finish()

	sp.LogKV("Token", token)

	view, err = wrp.Service.ReadByToken(ctx, token)

	sp.LogKV("Error", err)

	return view, err
}
