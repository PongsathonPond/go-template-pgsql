package implement

import (
	"context"

	"idev-cms-service/domain"
	"idev-cms-service/service/users/inout"
	"idev-cms-service/service/util"
)

func (impl *implementation) Delete(ctx context.Context, input *inout.UserDeleteInput) (err error) {
	user := &domain.Users{}
	filters := []string{
		impl.FilterString.MakeID(input.ID),
		impl.FilterString.MakeDeletedAtIsNull(),
	}

	if err = impl.RepoUsers.Read(ctx, filters, user); err != nil {
		return util.RepoReadErr(err)
	}

	if err = impl.RepoUsers.SoftDelete(ctx, filters); err != nil {
		return util.RepoDeleteErr(err)
	}

	//if err = impl.ServiceTokens.RevokeTokenByUserId(ctx, &input.ID); err != nil {
	//	impl.Log.Error(err)
	//}

	return nil
}
