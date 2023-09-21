package implement

import (
	"context"

	"idev-cms-service/service/users/inout"
	"idev-cms-service/service/util"
)

func (impl *implementation) CountByRole(ctx context.Context, input *inout.CountByRoleInput) (total int, err error) {
	filters := []string{
		impl.FilterString.MakeRoleID(input.RoleId),
		impl.FilterString.MakeDeletedAtIsNull(),
	}

	total, err = impl.RepoUsers.Count(ctx, filters)
	if err != nil {
		return 0, util.RepoCountErr(err)
	}

	return total, nil
}
