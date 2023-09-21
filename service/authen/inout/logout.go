package inout

type LogoutInput struct {
	AccessToken string `json:"access_token" validate:"required"`
}
