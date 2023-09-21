package inout

import (
	"idev-cms-service/domain"
	"idev-cms-service/service/util"
)

type UserReadView struct {
	ID            string         `json:"id,omitempty" example:"1400000000000000000"`
	FirstName     string         `json:"first_name" example:"John"`
	LastName      string         `json:"last_name" example:"Smith"`
	Email         string         `json:"email" example:"email@example.com"`
	Mobile        string         `json:"mobile" example:"0900000000"`
	UserName      string         `json:"username" example:"email@example.com"`
	Role          string         `json:"role" example:"1400000000000000000"`
	RoleData      *RoleData      `json:"role_data"`
	ImageProfile  string         `json:"image_profile" example:"https://media.idev.com/images/1624440506.jpeg"`
	Status        string         `json:"status" example:"active"`
	CreatedBy     string         `json:"created_by" example:"1400000000000000000"`
	CreatedByData *CreatedByData `json:"created_by_data"`
	CreatedAt     string         `json:"created_at" example:"2016-01-02 15:04:05"`
	UpdatedAt     string         `json:"updated_at" example:"2016-01-02 15:04:05"`
} // @Name UserReadView

type UserViewGRPC struct {
	ID        string `json:"id,omitempty" example:"1400000000000000000"`
	FirstName string `json:"first_name" example:"John"`
	LastName  string `json:"last_name" example:"Smith"`
	Email     string `json:"email" example:"example@example.com"`
	UserType  string `json:"user_type" example:"admin"`
	Status    string `json:"status" example:"active"`
} // @Name UserViewGRPC

type UserViewByToken struct {
	ID        string `json:"id,omitempty" example:"1400000000000000000"`
	FirstName string `json:"first_name" example:"John"`
	LastName  string `json:"last_name" example:"Smith"`
	UserGroup string `json:"user_group" example:"sil"`
	Role      string `json:"role" example:"1400000000000000000"`
} // @Name UserViewByToken

type MeView struct {
	ID           string    `json:"id,omitempty" example:"1400000000000000000"`
	FirstName    string    `json:"first_name" example:"John"`
	LastName     string    `json:"last_name" example:"Smith"`
	Email        string    `json:"email" example:"email@example.com"`
	UserName     string    `json:"username" example:"email@example.com"`
	Role         string    `json:"role" example:"1400000000000000000"`
	RoleData     *RoleData `json:"role_data"`
	Status       string    `json:"status" example:"active"`
	Mobile       string    `json:"mobile" example:"0900000000"`
	ImageProfile string    `json:"image_profile" example:"https://media.idev.com/images/1624440506.jpeg"`
	Verified     bool      `json:"verified" example:"false"`
	CreatedAt    string    `json:"created_at" example:"2016-01-02 15:04:05"`
	UpdatedAt    string    `json:"updated_at" example:"2016-01-02 15:04:05"`
} // @Name MeView

type UserGroupData struct {
	KeySlug string `json:"key_slug" example:"sil"`
	Name    string `json:"name" example:"Admin"`
} // @Name UserGroupData

type RoleData struct {
	ID   string `json:"id" example:"1400000000000000000"`
	Name string `json:"name" example:"Admin"`
} // @Name RoleData

type CreatedByData struct {
	ID        string `json:"id" example:"1400000000000000000"`
	FirstName string `json:"first_name" example:"Admin"`
	LastName  string `json:"last_name" example:"Admin"`
} // @Name CreatedByData

func UserToView(user *domain.Users, role *domain.Roles, createdBy *domain.Users, datetime util.DateTime) (view *UserReadView) {
	view = &UserReadView{
		ID:           user.ID,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		UserName:     user.Username,
		Email:        user.Email,
		Mobile:       user.Mobile,
		Role:         user.Role,
		Status:       user.Status,
		ImageProfile: user.ImageProfile,
		CreatedBy:    user.CreatedBy,
		CreatedAt:    datetime.ConvertUnixToDateTime(user.CreatedAt),
		UpdatedAt:    datetime.ConvertUnixToDateTime(user.UpdatedAt),
	}
	if role != nil {
		view.RoleData = &RoleData{
			ID:   role.ID,
			Name: role.Name,
		}
	}
	if createdBy != nil {
		view.CreatedByData = &CreatedByData{
			ID:        createdBy.ID,
			FirstName: createdBy.FirstName,
			LastName:  createdBy.LastName,
		}
	}
	return view
}

func UserToViewGRPC(user *domain.Users) (view *UserViewGRPC) {
	return &UserViewGRPC{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Status:    user.Status,
	}
}

func UserByTokenToView(user *domain.Users) (view *UserViewByToken) {
	return &UserViewByToken{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Role:      user.Role,
	}
}

func MeToView(user *domain.Users, role *domain.Roles, datetime util.DateTime) (view *MeView) {
	view = &MeView{
		ID:           user.ID,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Email:        user.Email,
		UserName:     user.Username,
		Role:         user.Role,
		Status:       user.Status,
		Mobile:       user.Mobile,
		ImageProfile: user.ImageProfile,
		CreatedAt:    datetime.ConvertUnixToDateTime(user.CreatedAt),
		UpdatedAt:    datetime.ConvertUnixToDateTime(user.UpdatedAt),
	}
	if role != nil {
		view.RoleData = &RoleData{
			ID:   role.ID,
			Name: role.Name,
		}
	}
	return view
}
