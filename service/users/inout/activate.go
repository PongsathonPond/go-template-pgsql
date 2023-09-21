package inout

type UserActivateInput struct {
	ActivateToken string `json:"activate_token" validate:"required,not-empty" example:"JrEjJNZpjuaWQA3Ihp..."`
	Password      string `json:"password" validate:"required,not-empty" example:"John"`
} // @Name UserActivateInput
