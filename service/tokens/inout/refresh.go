package inout

type RefreshTokenInput struct {
	RefreshToken string `json:"refresh_token" validate:"required" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXvCj9..."`
} // @Name RefreshTokenInput
