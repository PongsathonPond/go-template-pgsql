package inout

import (
	"github.com/modern-go/reflect2"
	"idev-cms-service/domain"
	"idev-cms-service/service/util"
)

type CategoryCreateInput struct {
	ID   string `json:"id" swaggerignore:"true"`
	Name string `json:"name"`
}

func (input *CategoryCreateInput) ToDomain(datetime util.DateTime) (category *domain.Category) {
	if reflect2.IsNil(input) {
		return &domain.Category{}
	}
	return &domain.Category{
		ID:        input.ID,
		Name:      input.Name,
		CreatedAt: datetime.GetUnixNow(),
		UpdatedAt: datetime.GetUnixNow(),
	}
}
