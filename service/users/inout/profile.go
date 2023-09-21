package inout

import (
	"idev-cms-service/domain"
	"idev-cms-service/service/util"
)

type ProfileUpdateInput struct {
	ID           string `json:"-" swaggerignore:"true"`
	Firstname    string `json:"first_name" validate:"required,not-empty" example:"John"`
	Lastname     string `json:"last_name" validate:"required,not-empty" example:"Smith"`
	Mobile       string `json:"mobile" example:"0900000000"`
	ImageProfile string `json:"image_profile" example:"https://media.idev.com/images/1624440506.jpeg"`
	AccessToken  string `json:"-" swaggerignore:"true"`
} // @Name ProfileUpdateInput

func (input *ProfileUpdateInput) ToDomain(datetime util.DateTime) (user *domain.Users) {
	return &domain.Users{
		ID:           input.ID,
		FirstName:    input.Firstname,
		LastName:     input.Lastname,
		Mobile:       input.Mobile,
		ImageProfile: input.ImageProfile,
		UpdatedAt:    datetime.GetUnixNow(),
	}
}
