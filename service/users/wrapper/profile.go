package wrapper

import (
	"context"

	"idev-cms-service/service/users/inout"

	"github.com/opentracing/opentracing-go"
)

func (wrp *Wrapper) UpdateMe(ctx context.Context, input *inout.ProfileUpdateInput) (err error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "service.users.UpdateMe")
	defer sp.Finish()

	sp.LogKV("ID", input.ID)
	sp.LogKV("Firstname", input.Firstname)
	sp.LogKV("Lastname", input.Lastname)
	sp.LogKV("Mobile", input.Mobile)
	sp.LogKV("ImageProfile", input.ImageProfile)

	err = wrp.Service.UpdateMe(ctx, input)

	sp.LogKV("Error", err)

	return err
}
