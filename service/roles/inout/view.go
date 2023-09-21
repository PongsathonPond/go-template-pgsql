package inout

import (
	"idev-cms-service/domain"
	"idev-cms-service/service/menus/inout"
	"idev-cms-service/service/util"
)

type RoleReadView struct {
	ID          string                        `json:"id" example:"123456789012345678"`
	Name        string                        `json:"name" example:"Admin"`
	Description string                        `json:"description" example:"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Ut interdum risus ante, nec dignissim orci mollis vitae"`
	RolesMenus  []*inout.MenuWithRolePermView `json:"roles_menus"`
	Status      string                        `json:"status" example:"active"`
	CreatedAt   string                        `json:"created_at" example:"2016-01-02 15:04:05"`
	UpdatedAt   string                        `json:"updated_at" example:"2016-01-02 15:04:05"`
} // name@ RoleReadView

type RolesListView struct {
	ID            string         `json:"id" example:"123456789012345678"`
	Name          string         `json:"name" example:"Admin"`
	Status        string         `json:"status" example:"active"`
	CreatedBy     string         `json:"created_by" example:"2016-01-02 15:04:05"`
	CreatedByData *createdByData `json:"created_by_data"`
	CreatedAt     string         `json:"created_at" example:"2016-01-02 15:04:05"`
	UpdatedAt     string         `json:"updated_at" example:"2016-01-02 15:04:05"`
} // name@ RolesListView

type RolesAllView struct {
	ID        string `json:"id" example:"123456789012345678"`
	Name      string `json:"name" example:"Admin"`
	Status    string `json:"status" example:"active"`
	CreatedAt string `json:"created_at" example:"2016-01-02 15:04:05"`
	UpdatedAt string `json:"updated_at" example:"2016-01-02 15:04:05"`
} // name@ RolesAllView

type MenuPermissionView struct {
	KeySlug  string               `json:"key_slug" example:"key_slug"`
	Name     string               `json:"name" example:"some_menu_name"`
	Icon     string               `json:"icon" example:"some_icon"`
	SubMenus []*subMenuPermission `json:"sub_menus"`
	Sort     int                  `json:"sort" example:"1"`
}

type subMenuPermission struct {
	KeySlug         string          `json:"key_slug" example:"key_slug"`
	Name            string          `json:"name" example:"some_sub_menu_name"`
	Icon            string          `json:"icon" example:"some_sub_menu_icon"`
	Path            string          `json:"path" example:"some_sub_menu_path"`
	RolesPermission *rolePermission `json:"role_permission"`
} // @Name subMenuPermission

type rolePermission struct {
	CanCreate bool `json:"can_create" example:"false"`
	CanRead   bool `json:"can_read" example:"true"`
	CanUpdate bool `json:"can_update" example:"true"`
	CanDelete bool `json:"can_delete" example:"false"`
}

type userGroupData struct {
	KeySlug string `json:"key_slug" example:"key_slug"`
	Name    string `json:"name" example:"John Smith"`
} // name@ userGroupData

type createdByData struct {
	Id        string `json:"id" example:"1410548436742172672"`
	FirstName string `json:"first_name" example:"John"`
	LastName  string `json:"last_name" example:"Smith"`
} // name@ createdByData

func RoleToReadView(role *domain.Roles, datetime util.DateTime, rolesMenus []*inout.MenuWithRolePermView) (view *RoleReadView) {
	view = &RoleReadView{
		ID:          role.ID,
		Name:        role.Name,
		Description: role.Description,
		RolesMenus:  rolesMenus,
		Status:      role.Status,
		CreatedAt:   datetime.ConvertUnixToDateTime(role.CreatedAt),
		UpdatedAt:   datetime.ConvertUnixToDateTime(role.UpdatedAt),
	}

	return view
}

func RoleToListView(role *domain.Roles, users *domain.Users, datetime util.DateTime) (view *RolesListView) {
	view = &RolesListView{
		ID:        role.ID,
		Name:      role.Name,
		Status:    role.Status,
		CreatedBy: role.CreatedBy,
		CreatedAt: datetime.ConvertUnixToDateTime(role.CreatedAt),
		UpdatedAt: datetime.ConvertUnixToDateTime(role.UpdatedAt),
	}

	if users != nil {
		view.CreatedByData = &createdByData{
			Id:        users.ID,
			FirstName: users.FirstName,
			LastName:  users.LastName,
		}
	}

	return view
}

func RoleAllToView(roles *domain.Roles, datetime util.DateTime) (view *RolesAllView) {
	view = &RolesAllView{
		ID:        roles.ID,
		Name:      roles.Name,
		Status:    roles.Status,
		CreatedAt: datetime.ConvertUnixToDateTime(roles.CreatedAt),
		UpdatedAt: datetime.ConvertUnixToDateTime(roles.UpdatedAt),
	}

	return view
}

func MapRolesWithSubMenus(menuList []*inout.MenuView, role *domain.Roles, perm util.Permission) []*inout.MenuWithRolePermView {
	var output []*inout.MenuWithRolePermView
	var subMenuView []*domain.RoleSubMenuView
	crud := &domain.RolesPermission{}

	for _, v := range menuList { // Loop menus
		subMenuView = []*domain.RoleSubMenuView{}
		for _, w := range v.SubMenus { // Loop sub menus
			crud = &domain.RolesPermission{}
			for _, permission := range role.Permissions { // Loop Permission role
				if w.KeySlug == permission.KeySlug { // if key sub menu match with perm key
					crud = perm.ConvertLevelToCRUD(permission.Permission)
					break
				}
			}
			subMenuView = append(subMenuView, &domain.RoleSubMenuView{
				KeySlug:          w.KeySlug,
				Name:             w.Name,
				Icon:             w.Icon,
				Path:             w.Path,
				CanSetPermission: w.CanSetPermission,
				RolesPermission:  *crud,
			})
		}
		output = append(output, &inout.MenuWithRolePermView{
			KeySlug:  v.KeySlug,
			Name:     v.Name,
			Icon:     v.Icon,
			SubMenus: subMenuView,
		})
	}

	return output
}

func UserPermissionToView(menus []interface{}, role *domain.Roles, perm util.Permission) (views []*MenuPermissionView) {
	var subMenuView []*subMenuPermission
	var crud *domain.RolesPermission
	var menu *domain.Menu

	for _, v := range menus {
		menu = v.(*domain.Menu)
		subMenuView = []*subMenuPermission{}
		for _, w := range menu.SubMenus {
			crud = &domain.RolesPermission{}
			for _, permission := range role.Permissions { // Loop Permission role
				if w.KeySlug == permission.KeySlug { // if key sub menu match with perm key
					crud = perm.ConvertLevelToCRUD(permission.Permission)
					subMenuView = append(subMenuView, &subMenuPermission{
						KeySlug: w.KeySlug,
						Name:    w.Name,
						Icon:    w.Icon,
						Path:    w.Path,
						RolesPermission: &rolePermission{
							CanCreate: crud.CanCreate,
							CanRead:   crud.CanRead,
							CanUpdate: crud.CanUpdate,
							CanDelete: crud.CanDelete,
						},
					})
					break
				}
			}
		}
		views = append(views, &MenuPermissionView{
			KeySlug:  menu.KeySlug,
			Name:     menu.Name,
			Icon:     menu.Icon,
			SubMenus: subMenuView,
			Sort:     menu.Sort,
		})
	}

	return views
}
