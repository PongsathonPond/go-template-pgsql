package inout

import (
	"time"

	"idev-cms-service/config"
	"idev-cms-service/domain"
	"idev-cms-service/service/util"
)

type UserView struct {
	ID            string        `json:"id,omitempty"`
	FirstName     string        `json:"first_name"`
	LastName      string        `json:"last_name"`
	Email         string        `json:"email"`
	Mobile        string        `json:"mobile"`
	Username      string        `json:"username"`
	UserType      string        `json:"user_type"`
	UserGroup     string        `json:"user_group"`
	UserGroupData UserGroupData `json:"user_group_data"`
	Role          string        `json:"role"`
	RoleData      RoleData      `json:"role_data"`
	Verified      bool          `json:"verified"`
	Status        string        `json:"status"`
	CreatedAt     string        `json:"created_at"`
	UpdatedAt     string        `json:"updated_at"`
} // @Name UserView

type UserViewWithToken struct {
	ID              string   `json:"id" example:"123456789012345678"`
	FirstName       string   `json:"first_name" example:"Administrator"`
	LastName        string   `json:"last_name" example:"See it Live"`
	Email           string   `json:"email" example:"admin@seeitlive.com"`
	Mobile          string   `json:"mobile" example:"021002002"`
	Username        string   `json:"username" example:"admin"`
	Role            string   `json:"role" example:"123456789012345678"`
	RoleData        RoleData `json:"role_data"`
	ImageProfile    string   `json:"image_profile" example:"https://media.seeitlivethailand.com/images/3b5e8a59-dc59-4548-862f-89267485f968.png"`
	AccessToken     string   `json:"access_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXvCj9..."`
	AccessTokenTTL  int64    `json:"access_token_expires" example:"3600"`
	RefreshToken    string   `json:"refresh_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXvCj9..."`
	RefreshTokenTTL int64    `json:"refresh_token_expires" example:"24000"`
	Status          string   `json:"status" example:"active"`
	CreatedAt       string   `json:"created_at" example:"2016-01-02 15:04:05"`
	UpdatedAt       string   `json:"updated_at" example:"2016-01-02 15:04:05"`
} // @Name UserViewWithToken

type UserGroupData struct {
	KeySlug string `json:"key_slug" example:"sil"`
	Name    string `json:"name" example:"See it Live"`
} // @Name UserGroupData

type RoleData struct {
	ID   string `json:"id" example:"123456789012345678"`
	Name string `json:"name" example:"Administrator"`
} // @Name RoleData

func UserWithTokenToView(user *domain.Users, role *domain.Roles, datetime util.DateTime, config *config.Config, accessToken, refreshToken *string) (view *UserViewWithToken) {
	view = &UserViewWithToken{
		ID:              user.ID,
		FirstName:       user.FirstName,
		LastName:        user.LastName,
		Email:           user.Email,
		Mobile:          user.Mobile,
		Username:        user.Username,
		Role:            user.Role,
		ImageProfile:    user.ImageProfile,
		AccessToken:     *accessToken,
		RefreshToken:    *refreshToken,
		AccessTokenTTL:  int64((time.Duration(config.AccessTokenTTL) * time.Minute).Seconds()),
		RefreshTokenTTL: int64((time.Duration(config.RefreshTokenTTL) * time.Minute).Seconds()),
		Status:          user.Status,
		CreatedAt:       datetime.ConvertUnixToDateTime(user.CreatedAt),
		UpdatedAt:       datetime.ConvertUnixToDateTime(user.UpdatedAt),
	}
	if role != nil {
		view.RoleData = RoleData{
			ID:   role.ID,
			Name: role.Name,
		}
	}
	return view
}
