package inout

import (
	"idev-cms-service/domain"
	"idev-cms-service/service/util"
)

type UserUpdateInput struct {
	ID           string `json:"-" swaggerignore:"true"`
	FirstName    string `json:"first_name" validate:"required,not-empty" example:"John"`
	LastName     string `json:"last_name" validate:"required,not-empty" example:"Smith"`
	Email        string `json:"email" example:"email@example.com"`
	Mobile       string `json:"mobile" example:"0900000000"`
	UserName     string `json:"-" swaggerignore:"true"`
	Role         string `json:"role" validate:"required" example:"1400000000000000000"`
	Status       string `json:"status" validate:"required,oneof=active inactive" example:"active"`
	ImageProfile string `json:"image_profile" example:"https://media.idev.com/images/1624440506.jpeg"`
} // @Name UserUpdateInput

type UserUpdatePasswordInput struct {
	Email    string `json:"email" example:"email@example.com"`
	Password string `json:"password" example:"123456789"`
} // @Name UserUpdatePasswordInput

type UserChangePasswordInput struct {
	UserID      string `json:"-" swaggerignore:"true"`
	OldPassword string `json:"old_password" validate:"required,not-empty" example:"123456789"`
	NewPassword string `json:"new_password" validate:"required,not-empty" example:"123456789"`
} // @Name UserChangePasswordInput

func (input *UserUpdateInput) ToDomain(datetime util.DateTime) (user *domain.Users) {
	return &domain.Users{
		ID:           input.ID,
		FirstName:    input.FirstName,
		LastName:     input.LastName,
		Email:        input.Email,
		Mobile:       input.Mobile,
		Username:     input.UserName,
		Role:         input.Role,
		Status:       input.Status,
		ImageProfile: input.ImageProfile,
		UpdatedAt:    datetime.GetUnixNow(),
	}
}
