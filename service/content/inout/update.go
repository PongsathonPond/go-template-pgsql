package inout

import (
	"idev-cms-service/domain"
	"idev-cms-service/service/util"

	"github.com/modern-go/reflect2"
)

type ContentUpdateInput struct {
	ID          string `json:"id" validate:"required" swaggerignore:"true"`
	Title       string `json:"title" validate:"required,not-empty"`
	Images      string `json:"images"`
	Description string `json:"description"`
	CategoryID  string `json:"category_id"`
	Content     string `json:"content"`
	LinkOnePage string `json:"link_one_page"`
}

func (input *ContentUpdateInput) ToDomain(datetime util.DateTime) (content *domain.Content) {
	if reflect2.IsNil(input) {
		return &domain.Content{}
	}
	return &domain.Content{
		Title:       input.Title,
		Images:      input.Images,
		Description: input.Description,
		CategoryID:  input.CategoryID,
		Content:     input.Content,
		LinkOnePage: input.LinkOnePage,
		UpdatedAt:   datetime.GetUnixNow(),
	}
}
