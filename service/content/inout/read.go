package inout

type ContentReadInput struct {
	ID string `json:"-" validate:"required"`
}
