package wrapper

import (
	"context"

	"idev-cms-service/service/users/inout"

	"github.com/opentracing/opentracing-go"
)

func (wrp *Wrapper) Update(ctx context.Context, input *inout.UserUpdateInput) (err error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "service.users.Update")
	defer sp.Finish()

	sp.LogKV("ID", input.ID)
	sp.LogKV("FirstName", input.FirstName)
	sp.LogKV("LastName", input.LastName)
	sp.LogKV("Mobile", input.Mobile)
	sp.LogKV("Role", input.Role)
	sp.LogKV("Status", input.Status)
	sp.LogKV("ImageProfile", input.ImageProfile)

	err = wrp.Service.Update(ctx, input)

	sp.LogKV("Error", err)

	return err
}
