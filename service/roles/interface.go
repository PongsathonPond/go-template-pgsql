package roles

import (
	"context"

	"idev-cms-service/domain"
	"idev-cms-service/service/roles/inout"
)

//go:generate mockery --name=Service
type Service interface {
	All(ctx context.Context, input *inout.RoleAllInput) (items []*inout.RolesAllView, err error)
	List(ctx context.Context, opt *domain.PageOption) (total int, items []*inout.RolesListView, err error)
	Create(ctx context.Context, input *inout.RoleCreateInput) (ID string, err error)
	Read(ctx context.Context, input *inout.RoleReadInput) (view *inout.RoleReadView, err error)
	Update(ctx context.Context, input *inout.RoleUpdateInput) (err error)
	Delete(ctx context.Context, input *inout.RoleDeleteInput) (err error)
	GetUserPermission(ctx context.Context, token string) (views []*inout.MenuPermissionView, err error)
	CheckPermission(ctx context.Context, input *inout.CheckPermissionInput) (isPermission bool)
}
