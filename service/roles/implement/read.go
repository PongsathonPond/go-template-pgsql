package implement

import (
	"context"

	"idev-cms-service/domain"
	"idev-cms-service/service/roles/inout"
	"idev-cms-service/service/util"
)

func (impl *implementation) Read(ctx context.Context, input *inout.RoleReadInput) (view *inout.RoleReadView, err error) {
	role := &domain.Roles{}
	filters := []string{
		impl.FilterString.MakeID(input.ID),
		impl.FilterString.MakeDeletedAtIsNull(),
	}

	if err = impl.Repo.Read(ctx, filters, role); err != nil {
		return nil, util.RepoReadErr(err)
	}

	_, records, err := impl.ServiceMenus.All(ctx)
	rolesMenus := inout.MapRolesWithSubMenus(records, role, impl.Perm)

	return inout.RoleToReadView(role, impl.DateTime, rolesMenus), nil

}
