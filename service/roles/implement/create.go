package implement

import (
	"context"

	"idev-cms-service/domain"
	"idev-cms-service/service/roles/inout"
	"idev-cms-service/service/util"
)

func (impl *implementation) Create(ctx context.Context, input *inout.RoleCreateInput) (id string, err error) {
	if err = impl.Validator.Validate(input); err != nil {
		return "", util.ValidationCreateErr(err)
	}

	var perms []domain.MenuPermissions

	input.ID = impl.UUID.Generate()

	if len(input.Permissions) > 0 {
		for _, v := range input.Permissions {
			perm := &domain.MenuPermissions{
				KeySlug:    v.KeySubMenu,
				Permission: impl.Perm.ConvertCRUDToLevel(v.Permission.CanCreate, v.Permission.CanRead, v.Permission.CanUpdate, v.Permission.CanDelete),
			}
			perms = append(perms, *perm)
		}
	}

	role := input.ToDomain(impl.DateTime, perms)

	_, err = impl.Repo.Create(ctx, role)
	if err != nil {
		return "", util.RepoCreateErr(err)
	}

	return input.ID, nil
}
