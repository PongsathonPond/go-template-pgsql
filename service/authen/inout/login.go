package inout

type LoginInput struct {
	Username  string `json:"username" validate:"required" example:"admin@idev.com"`
	Password  string `json:"password" validate:"required" example:"adminidev1234"`
	UserAgent string `json:"user_agent" swaggerignore:"true"`
} // @Name LoginInput
