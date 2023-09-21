package util

import (
	"idev-cms-service/domain"
)

const (
	Create = 1 << iota
	Read
	Update
	Delete
)

//go:generate mockery --name=Permission
type Permission interface {
	ConvertCRUDToLevel(isCreate, isRead, isUpdate, isDelete bool) (level uint8)
	ConvertLevelToCRUD(level uint8) *domain.RolesPermission
	ConvertSliceLevelToSliceCRUD(input []*domain.Permissions) (permData []*domain.PermissionsCRUD)
	ConvertSubMenuToSubMenuView(input []*domain.SubMenu) (view []*domain.SubMenuView)
}

type permission struct{}

type permissionUInt8 uint8

func NewPermission() (perm Permission) {
	return &permission{}
}

func (perm *permission) ConvertCRUDToLevel(isCreate, isRead, isUpdate, isDelete bool) (level uint8) {
	level += perm.checkPerm(isCreate, Create)
	level += perm.checkPerm(isRead, Read)
	level += perm.checkPerm(isUpdate, Update)
	level += perm.checkPerm(isDelete, Delete)
	return level
}

func (perm *permission) checkPerm(input bool, level uint8) (res uint8) {
	if input {
		return level
	}
	return res
}

func (perm *permission) ConvertLevelToCRUD(level uint8) *domain.RolesPermission {
	return &domain.RolesPermission{
		CanCreate: permissionUInt8(level).canCreate(),
		CanRead:   permissionUInt8(level).canRead(),
		CanUpdate: permissionUInt8(level).canUpdate(),
		CanDelete: permissionUInt8(level).canDelete(),
	}
}

func (perm *permission) ConvertSliceLevelToSliceCRUD(input []*domain.Permissions) (permData []*domain.PermissionsCRUD) {
	for _, v := range input {
		permData = append(permData, &domain.PermissionsCRUD{
			KeySubMenu: v.KeySubMenu,
			Permissions: domain.RolesPermission{
				CanCreate: permissionUInt8(v.Level).canCreate(),
				CanRead:   permissionUInt8(v.Level).canRead(),
				CanUpdate: permissionUInt8(v.Level).canUpdate(),
				CanDelete: permissionUInt8(v.Level).canDelete(),
			},
		})
	}
	return permData
}

func (perm *permission) ConvertSubMenuToSubMenuView(input []*domain.SubMenu) (view []*domain.SubMenuView) {
	for _, v := range input {
		view = append(view, &domain.SubMenuView{
			KeySlug: v.KeySlug,
			Name:    v.Name,
			Icon:    v.Icon,
			Path:    v.Path,
			CanSetPermission: domain.RolesPermission{
				CanCreate: permissionUInt8(v.CanSetPermission).canCreate(),
				CanRead:   permissionUInt8(v.CanSetPermission).canRead(),
				CanUpdate: permissionUInt8(v.CanSetPermission).canUpdate(),
				CanDelete: permissionUInt8(v.CanSetPermission).canDelete(),
			},
			Status: v.Status,
		})
	}
	return view
}

func (p permissionUInt8) canCreate() bool {
	return p&Create != 0
}

func (p permissionUInt8) canRead() bool {
	return p&Read != 0
}

func (p permissionUInt8) canUpdate() bool {
	return p&Update != 0
}

func (p permissionUInt8) canDelete() bool {
	return p&Delete != 0
}
