package inout

import (
	"idev-cms-service/domain"
	"idev-cms-service/service/util"
)

type UserCreateInput struct {
	ID           string `json:"-" swaggerignore:"true"`
	FirstName    string `json:"first_name" validate:"required,not-empty" example:"John"`
	LastName     string `json:"last_name" validate:"required,not-empty" example:"Smith"`
	Email        string `json:"email" validate:"required,email" example:"email@example.com"`
	Password     string `json:"password" validate:"required" example:"U2FsdGVkX1+pqoATrmFheH7j5P6RXSwelcxoiiFOP2U="`
	Mobile       string `json:"mobile" example:"0900000000"`
	Role         string `json:"role" validate:"required" example:"1400000000000000000"`
	ImageProfile string `json:"image_profile" example:"https://media.idev.com/images/1624440506.jpeg"`
	Status       string `json:"status" validate:"required,oneof=active inactive" example:"active"`
	CreatedBy    string `json:"-" swaggerignore:"true"`
	AccessToken  string `json:"-" swaggerignore:"true"`
} // @Name UserCreateInput

func (input *UserCreateInput) ToDomain(datetime util.DateTime) (user *domain.Users) {
	user = &domain.Users{
		ID:           input.ID,
		FirstName:    input.FirstName,
		LastName:     input.LastName,
		Email:        input.Email,
		Mobile:       input.Mobile,
		Username:     input.Email,
		Password:     input.Password,
		Role:         input.Role,
		ImageProfile: input.ImageProfile,
		Status:       input.Status,
		CreatedBy:    input.CreatedBy,
		CreatedAt:    datetime.GetUnixNow(),
		UpdatedAt:    datetime.GetUnixNow(),
	}

	return user
}
