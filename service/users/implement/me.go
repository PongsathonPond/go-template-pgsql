package implement

import (
	"context"

	"idev-cms-service/domain"
	"idev-cms-service/service/users/inout"
	"idev-cms-service/service/util"
)

func (impl *implementation) GetMe(ctx context.Context, input *inout.GetMeInput) (view *inout.MeView, err error) {
	user := &domain.Users{}

	userID := impl.ServiceTokens.GetUserIDByToken(ctx, &input.AccessToken)
	filters := []string{
		impl.FilterString.MakeID(userID),
		impl.FilterString.MakeDeletedAtIsNull(),
	}

	err = impl.RepoUsers.Read(ctx, filters, user)
	if err != nil {
		return nil, util.RepoReadErr(err)
	}

	filters = []string{
		impl.FilterString.MakeID(user.Role),
		impl.FilterString.MakeDeletedAtIsNull(),
	}
	role := &domain.Roles{}
	if err = impl.RepoRoles.Read(ctx, filters, role); err != nil {
		role = nil
	}

	return inout.MeToView(user, role, impl.DateTime), nil
}
