package inout

import (
	"idev-cms-service/domain"
	"idev-cms-service/service/util"
)

type MenuView struct {
	ID        string                `json:"id" example:"123456789012345678"`
	KeySlug   string                `json:"key_slug" example:"key_slug"`
	Name      string                `json:"name" example:"John Smith"`
	Icon      string                `json:"icon" example:"some_icon"`
	SubMenus  []*domain.SubMenuView `json:"sub_menus"`
	Sort      int                   `json:"sort" example:"1"`
	Status    string                `json:"status" example:"active"`
	CreatedAt string                `json:"created_at" example:"2016-01-02 15:04:05"`
	UpdatedAt string                `json:"updated_at" example:"2016-01-02 15:04:05"`
}

type MenuWithRolePermView struct {
	KeySlug  string                    `json:"key_slug" example:"key_slug"`
	Name     string                    `json:"name" example:"some_menu_name"`
	Icon     string                    `json:"icon" example:"some_icon"`
	SubMenus []*domain.RoleSubMenuView `json:"sub_menus"`
}

func MenuToView(menu *domain.Menu, datetime util.DateTime, perm util.Permission) (view *MenuView) {
	return &MenuView{
		ID:        menu.ID,
		KeySlug:   menu.KeySlug,
		Name:      menu.Name,
		Icon:      menu.Icon,
		SubMenus:  perm.ConvertSubMenuToSubMenuView(menu.SubMenus),
		Sort:      menu.Sort,
		Status:    menu.Status,
		CreatedAt: datetime.ConvertUnixToDateTime(menu.CreatedAt),
		UpdatedAt: datetime.ConvertUnixToDateTime(menu.UpdatedAt),
	}
}
