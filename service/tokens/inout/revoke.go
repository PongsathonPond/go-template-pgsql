package inout

type RevokeTokenInput struct {
	AccessToken string `json:"access_token" validate:"required"`
}

type RevokeTokenByUserIdInput struct {
	UserID string `json:"-"`
}
