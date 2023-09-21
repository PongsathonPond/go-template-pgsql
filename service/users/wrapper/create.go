package wrapper

import (
	"context"

	"idev-cms-service/service/users/inout"

	"github.com/opentracing/opentracing-go"
)

func (wrp *Wrapper) Create(ctx context.Context, input *inout.UserCreateInput) (id string, err error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "service.users.Create")
	defer sp.Finish()

	sp.LogKV("FirstName", input.FirstName)
	sp.LogKV("LastName", input.LastName)
	sp.LogKV("Email", input.Email)
	sp.LogKV("Mobile", input.Mobile)
	sp.LogKV("Role", input.Role)
	sp.LogKV("Status", input.Status)
	sp.LogKV("ProfileImage", input.ImageProfile)
	sp.LogKV("CreatedBy", input.CreatedBy)

	id, err = wrp.Service.Create(ctx, input)

	sp.LogKV("Error", err)

	return id, err
}
