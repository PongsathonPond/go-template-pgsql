package domain

type Menu struct {
	ID        string     `bson:"_id"`
	KeySlug   string     `bson:"key_slug"`
	Name      string     `bson:"name"`
	Icon      string     `bson:"icon"`
	SubMenus  []*SubMenu `bson:"sub_menus"`
	Sort      int        `bson:"sort"`
	Status    string     `bson:"status"`
	CreatedAt int64      `bson:"created_at"`
	UpdatedAt int64      `bson:"updated_at"`
	DeletedAt int64      `bson:"deleted_at,omitempty"`
}

type SubMenu struct {
	KeySlug          string `bson:"key_slug"`
	Name             string `bson:"name"`
	Icon             string `bson:"icon"`
	Path             string `bson:"path"`
	Sort             int    `bson:"sort"`
	CanSetPermission uint8  `bson:"can_set_permission"`
	Status           string `bson:"status"`
}

type SubMenuView struct {
	KeySlug          string          `json:"key_slug" example:"key_slug"`
	Name             string          `json:"name" example:"John Smith"`
	Icon             string          `json:"icon" example:"some_icon"`
	Path             string          `json:"path" example:"some_path"`
	CanSetPermission RolesPermission `json:"can_set_permission"`
	Status           string          `json:"status" example:"active"`
} // @name SubMenuView

type RoleSubMenuView struct {
	KeySlug          string          `json:"key_slug" example:"key_slug"`
	Name             string          `json:"name" example:"some_sub_menu_name"`
	Icon             string          `json:"icon" example:"some_sub_menu_icon"`
	Path             string          `json:"path" example:"some_sub_menu_path"`
	CanSetPermission RolesPermission `json:"can_set_permission"`
	RolesPermission  RolesPermission `json:"roles_permission"`
} // @name RoleSubMenuView
