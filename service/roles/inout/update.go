package inout

import (
	"idev-cms-service/domain"
	"idev-cms-service/service/util"

	"github.com/modern-go/reflect2"
)

type RoleUpdateInput struct {
	ID          string            `json:"id" swaggerignore:"true"`
	Name        string            `json:"name" validate:"required" example:"Admin"`
	Description string            `json:"description" example:"Lorem ipsum dolor"`
	Permissions []PermissionInput `json:"permissions" validate:"min=1,required"`
	Status      string            `json:"status" validate:"required,oneof=active inactive" example:"inactive"`
}

func (input *RoleUpdateInput) ToDomain(datetime util.DateTime, perms []domain.MenuPermissions) (role *domain.Roles) {
	if reflect2.IsNil(input) {
		return &domain.Roles{}
	}

	return &domain.Roles{
		ID:          input.ID,
		Name:        input.Name,
		Description: input.Description,
		Permissions: perms,
		Status:      input.Status,
		UpdatedAt:   datetime.GetUnixNow(),
	}
}
