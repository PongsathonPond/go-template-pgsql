package inout

import (
	"idev-cms-service/domain"
	"idev-cms-service/service/util"

	"github.com/modern-go/reflect2"
)

type CategoryUpdateInput struct {
	ID   string `json:"id" swaggerignore:"true"`
	Name string `json:"name"`
}

func (input *CategoryUpdateInput) ToDomain(datetime util.DateTime) (category *domain.Category) {
	if reflect2.IsNil(input) {
		return &domain.Category{}
	}
	return &domain.Category{
		ID:        input.ID,
		Name:      input.Name,
		UpdatedAt: datetime.GetUnixNow(),
	}
}
