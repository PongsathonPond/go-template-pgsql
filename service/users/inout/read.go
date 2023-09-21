package inout

type ReadInput struct {
	ID          string `json:"-" validate:"required"`
	AccessToken string `json:"-"`
}
