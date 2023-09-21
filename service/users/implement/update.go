package implement

import (
	"context"

	"idev-cms-service/domain"
	"idev-cms-service/service/users/inout"
	"idev-cms-service/service/util"

	"github.com/imdario/mergo"
)

func (impl *implementation) Update(ctx context.Context, input *inout.UserUpdateInput) (err error) {
	if err = impl.Validator.Validate(input); err != nil {
		return util.ValidationUpdateErr(err)
	}

	user := &domain.Users{}
	filters := []string{
		impl.FilterString.MakeID(input.ID),
		impl.FilterString.MakeDeletedAtIsNull(),
	}

	if err = impl.RepoUsers.Read(ctx, filters, user); err != nil {
		return util.RepoReadErr(err)
	}

	update := input.ToDomain(impl.DateTime)

	if err = mergo.Merge(user, update, mergo.WithOverride); err != nil {
		return util.UnknownErr(err)
	}

	user.Mobile = input.Mobile
	user.ImageProfile = input.ImageProfile

	if err = impl.RepoUsers.Update(ctx, filters, user); err != nil {
		return util.RepoUpdateErr(err)
	}

	//if input.Status == "inactive" {
	//	if err = impl.ServiceTokens.RevokeTokenByUserId(ctx, &input.ID); err != nil {
	//		impl.Log.Error(err)
	//	}
	//}

	return nil
}
