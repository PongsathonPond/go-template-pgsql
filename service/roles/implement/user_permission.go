package implement

import (
	"context"

	"idev-cms-service/domain"
	"idev-cms-service/service/roles/inout"
	"idev-cms-service/service/util"
)

func (impl implementation) GetUserPermission(ctx context.Context, token string) (views []*inout.MenuPermissionView, err error) {
	var keySubMenus []string

	user, err := impl.UserService.ReadByToken(ctx, &token)
	if err != nil {
		return nil, util.RepoReadErr(err)
	}

	role := &domain.Roles{}
	filters := []string{
		impl.FilterString.MakeID(user.Role),
		impl.FilterString.MakeDeletedAtIsNull(),
	}

	if err = impl.Repo.Read(ctx, filters, role); err != nil {
		return nil, util.RepoReadErr(err)
	}

	for _, v := range role.Permissions {
		keySubMenus = append(keySubMenus, v.KeySlug)
	}

	filters = []string{
		impl.FilterString.MakeSubMenuIn(keySubMenus),
		impl.FilterString.MakeStatus("active"),
		impl.FilterString.MakeDeletedAtIsNull(),
	}
	sorts := []string{"sort:asc"}

	opt := &domain.QueryOption{
		Filters: filters,
		Search:  nil,
		Sorts:   sorts,
	}
	_, records, err := impl.MenusRepo.Find(ctx, opt, &domain.Menu{})
	if err != nil {
		return nil, util.RepoFindErr(err)
	}

	views = inout.UserPermissionToView(records, role, impl.Perm)

	return views, nil
}
