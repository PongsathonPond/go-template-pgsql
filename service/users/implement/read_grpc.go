package implement

import (
	"context"

	"idev-cms-service/domain"
	"idev-cms-service/service/users/inout"
	"idev-cms-service/service/util"
)

func (impl *implementation) ReadGRPC(ctx context.Context, input *inout.ReadInput) (view *inout.UserViewGRPC, err error) {
	user := &domain.Users{}
	filters := []string{
		impl.FilterString.MakeID(input.ID),
		impl.FilterString.MakeDeletedAtIsNull(),
	}

	err = impl.RepoUsers.Read(ctx, filters, user)
	if err != nil {
		return nil, util.RepoReadErr(err)
	}

	return inout.UserToViewGRPC(user), nil
}
