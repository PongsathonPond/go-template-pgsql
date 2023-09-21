package implement

import (
	"context"

	"idev-cms-service/domain"
	"idev-cms-service/service/users/inout"
	"idev-cms-service/service/util"

	"github.com/imdario/mergo"
)

func (impl *implementation) UpdateMe(ctx context.Context, input *inout.ProfileUpdateInput) (err error) {
	if err = impl.Validator.Validate(input); err != nil {
		return util.ValidationUpdateErr(err)
	}

	user := &domain.Users{}

	userID := impl.ServiceTokens.GetUserIDByToken(ctx, &input.AccessToken)
	filters := []string{
		impl.FilterString.MakeID(userID),
		impl.FilterString.MakeDeletedAtIsNull(),
	}

	if err = impl.RepoUsers.Read(ctx, filters, user); err != nil {
		return util.RepoReadErr(err)
	}

	update := input.ToDomain(impl.DateTime)

	if err = mergo.Merge(user, update, mergo.WithOverride); err != nil {
		return util.UnknownErr(err)
	}

	if err = impl.RepoUsers.Update(ctx, filters, user); err != nil {
		return util.RepoUpdateErr(err)
	}

	return nil
}
