package domain

type Permissions struct {
	KeySubMenu string `bson:"key_sub_menu"`
	Level      uint8  `bson:"level"`
}

type PermissionsCRUD struct {
	KeySubMenu  string          `json:"key_sub_menu"`
	Permissions RolesPermission `json:"permissions"`
}

type RolesPermission struct {
	CanCreate bool `json:"can_create" example:"false"`
	CanRead   bool `json:"can_read" example:"true"`
	CanUpdate bool `json:"can_update" example:"true"`
	CanDelete bool `json:"can_delete" example:"false"`
} // @name RolesPermission
