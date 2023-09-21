package inout

import (
	"idev-cms-service/domain"
	"idev-cms-service/service/util"

	"github.com/modern-go/reflect2"
)

type RoleCreateInput struct {
	ID          string            `json:"id" swaggerignore:"true"`
	Name        string            `json:"name" validate:"required" example:"Admin"`
	Description string            `json:"description" example:"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Ut interdum risus ante, nec dignissim orci mollis vitae"`
	Permissions []PermissionInput `json:"permissions" validate:"min=1,required"`
	Status      string            `json:"status" validate:"required,oneof=active inactive" example:"active"`
	CreatedBy   string            `json:"created_by" swaggerignore:"true"`
	CreatedAt   int64             `json:"created_at" swaggerignore:"true"`
	UpdatedAt   int64             `json:"updated_at" swaggerignore:"true"`
	AccessToken string            `json:"-" swaggerignore:"true"`
} // @Name RoleCreateInput

func (input *RoleCreateInput) ToDomain(datetime util.DateTime, perms []domain.MenuPermissions) (role *domain.Roles) {
	if reflect2.IsNil(input) {
		return &domain.Roles{}
	}
	return &domain.Roles{
		ID:          input.ID,
		Name:        input.Name,
		Description: input.Description,
		Permissions: perms,
		Status:      input.Status,
		CreatedBy:   input.CreatedBy,
		CreatedAt:   datetime.GetUnixNow(),
		UpdatedAt:   datetime.GetUnixNow(),
	}
}
