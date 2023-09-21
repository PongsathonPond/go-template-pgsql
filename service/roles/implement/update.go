package implement

import (
	"context"

	"idev-cms-service/domain"
	"idev-cms-service/service/roles/inout"
	"idev-cms-service/service/util"

	"github.com/imdario/mergo"
)

func (impl *implementation) Update(ctx context.Context, input *inout.RoleUpdateInput) (err error) {
	if err = impl.Validator.Validate(input); err != nil {
		return util.ValidationUpdateErr(err)
	}

	role := &domain.Roles{}
	filters := []string{
		impl.FilterString.MakeID(input.ID),
		impl.FilterString.MakeDeletedAtIsNull(),
	}

	if err = impl.Repo.Read(ctx, filters, role); err != nil {
		return util.RepoReadErr(err)
	}

	var perms []domain.MenuPermissions

	if len(input.Permissions) > 0 {
		for _, v := range input.Permissions {
			perm := &domain.MenuPermissions{
				KeySlug:    v.KeySubMenu,
				Permission: impl.Perm.ConvertCRUDToLevel(v.Permission.CanCreate, v.Permission.CanRead, v.Permission.CanUpdate, v.Permission.CanDelete),
			}

			perms = append(perms, *perm)
		}
	} else {
		perms = role.Permissions
	}

	update := input.ToDomain(impl.DateTime, perms)

	if err = mergo.Merge(role, update, mergo.WithOverride); err != nil {
		return util.UnknownErr(err)
	}

	if err = impl.Repo.Update(ctx, filters, role); err != nil {
		return util.RepoUpdateErr(err)
	}

	return nil
}
