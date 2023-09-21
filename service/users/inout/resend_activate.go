package inout

type ResendActivateInput struct {
	UserID string `json:"user_id" validate:"required,not-empty" example:"123456789"`
} // @name ResendActivateInput
