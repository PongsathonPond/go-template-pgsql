package inout

import (
	"idev-cms-service/domain"
	"idev-cms-service/service/util"
)

type CategoryView struct {
	ID        string `json:"id" example:"123456789012345678"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at" example:"2016-01-02 15:04:05"`
	UpdatedAt string `json:"updated_at" example:"2016-01-02 15:04:05"`
}

func CategoryToView(category *domain.Category, datetime util.DateTime) (view *CategoryView) {
	return &CategoryView{
		ID:        category.ID,
		Name:      category.Name,
		CreatedAt: datetime.ConvertUnixToDateTime(category.CreatedAt),
		UpdatedAt: datetime.ConvertUnixToDateTime(category.UpdatedAt),
	}

}
