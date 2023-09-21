package wrapper

import (
	"context"
	"fmt"

	"idev-cms-service/service/authen/inout"

	"github.com/opentracing/opentracing-go"
)

func (wrp *Wrapper) Login(ctx context.Context, input *inout.LoginInput) (view *inout.UserViewWithToken, err error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "service.authen.Login")
	defer sp.Finish()

	sp.LogKV("Username", input.Username)
	sp.LogKV("Password", input.Password)

	view, err = wrp.Service.Login(ctx, input)

	if view != nil {
		sp.LogKV("ID", view.ID)
		sp.LogKV("Name", fmt.Sprintf("%s %s", view.FirstName, view.LastName))
	} else {
		sp.LogKV("ID", nil)
	}
	sp.LogKV("Error", err)

	return view, err
}
