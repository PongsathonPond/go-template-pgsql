package inout

import (
	"github.com/modern-go/reflect2"
	"idev-cms-service/domain"
	"idev-cms-service/service/util"
)

type ContentCreateInput struct {
	ID          string `json:"id" swaggerignore:"true"`
	Title       string `json:"title" validate:"required,not-empty"`
	CategoryID  string `json:"category_id"`
	Images      string `json:"images"`
	Description string `json:"description"`
	Content     string `json:"content"`
	LinkOnePage string `json:"link_one_page"`
	CreatedBy   string `json:"-" swaggerignore:"true"`
	AccessToken string `json:"-" swaggerignore:"true"`
}

func (input *ContentCreateInput) ToDomain(datetime util.DateTime) (content *domain.Content) {
	if reflect2.IsNil(input) {
		return &domain.Content{}
	}
	return &domain.Content{
		ID:          input.ID,
		Title:       input.Title,
		CategoryID:  input.CategoryID,
		Images:      input.Images,
		Description: input.Description,
		Content:     input.Content,
		LinkOnePage: input.LinkOnePage,
		CreatedBy:   input.CreatedBy,
		CreatedAt:   datetime.GetUnixNow(),
		UpdatedAt:   datetime.GetUnixNow(),
	}
}
