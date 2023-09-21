package implement

import (
	"context"

	"idev-cms-service/service/users/inout"
	"idev-cms-service/service/util"
)

func (impl *implementation) Create(ctx context.Context, input *inout.UserCreateInput) (id string, err error) {
	if err = impl.Validator.Validate(input); err != nil {
		return "", util.ValidationCreateErr(err)
	}

	input.ID = impl.UUID.Generate()
	//input.CreatedBy = impl.ServiceTokens.GetUserIDByToken(ctx, &input.AccessToken)

	user := input.ToDomain(impl.DateTime)

	_, err = impl.RepoUsers.Create(ctx, user)
	if err != nil {
		return "", util.RepoCreateErr(err)
	}

	return input.ID, nil
}
