package inout

type CategoryReadInput struct {
	ID string `json:"-" validate:"required"`
}
