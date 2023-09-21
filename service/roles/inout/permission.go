package inout

type PermissionInput struct {
	KeySubMenu string          `json:"key_sub_menu" validate:"required" example:"ga"`
	Permission RolesPermission `json:"permission"`
} // @Name PermissionInput

type CheckPermissionInput struct {
	RoleID   string
	Method   string
	EndPoint string
	Service  string
} // @Name CheckPermissionInput

type RolesPermission struct {
	CanCreate bool `json:"can_create" example:"false"`
	CanRead   bool `json:"can_read" example:"true"`
	CanUpdate bool `json:"can_update" example:"false"`
	CanDelete bool `json:"can_delete" example:"true"`
} // @Name RolesPermission
