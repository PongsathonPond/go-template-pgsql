package inout

type UserDeleteInput struct {
	ID string `json:"-" validate:"required"`
} // @Name UserDeleteInput
