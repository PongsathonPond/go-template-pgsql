package inout

type VerifyActivateTokenInput struct {
	ActivateToken string `json:"activate_token" validate:"required,not-empty" example:"JrEjJNZpjuaWQA3Ihp..."`
} // @Name VerifyActivateTokenInput
