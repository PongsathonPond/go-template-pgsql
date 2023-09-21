package implement

import (
	"context"

	"idev-cms-service/domain"
	"idev-cms-service/service/users/inout"
	"idev-cms-service/service/util"
)

func (impl *implementation) ReadByToken(ctx context.Context, token *string) (view *inout.UserViewByToken, err error) {
	user := &domain.Users{}

	userID := impl.ServiceTokens.GetUserIDByToken(ctx, token)
	filters := []string{
		impl.FilterString.MakeID(userID),
		impl.FilterString.MakeDeletedAtIsNull(),
	}
	err = impl.RepoUsers.Read(ctx, filters, user)
	if err != nil {
		return nil, util.RepoReadErr(err)
	}

	return inout.UserByTokenToView(user), nil
}
