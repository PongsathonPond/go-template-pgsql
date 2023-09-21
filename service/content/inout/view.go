package inout

import (
	"idev-cms-service/domain"
	"idev-cms-service/service/util"
)

type ContentView struct {
	ID            string         `json:"id" example:"123456789012345678"`
	Title         string         `json:"title"`
	Images        string         `json:"images"`
	Description   string         `json:"description"`
	CategoryData  *CategoryData  `json:"category_data"`
	Content       string         `json:"content"`
	LinkOnePage   string         `json:"link_one_page"`
	CreatedByData *CreatedByData `json:"created_by_data"`
	CreatedAt     string         `json:"created_at" example:"2016-01-02 15:04:05"`
	UpdatedAt     string         `json:"updated_at" example:"2016-01-02 15:04:05"`
}

type CategoryData struct {
	ID   string `json:"id" example:"1"`
	Name string `json:"name"`
}

type CreatedByData struct {
	ID        string `json:"id" example:"1400000000000000000"`
	FirstName string `json:"first_name" example:"Admin"`
	LastName  string `json:"last_name" example:"Admin"`
}

func CreatedUserByData(users *domain.Users) (create *CreatedByData) {
	return &CreatedByData{

		ID:        users.ID,
		FirstName: users.FirstName,
		LastName:  users.LastName,
	}

}
func ContentToView(content *domain.Content, category *domain.Category, users *domain.Users, datetime util.DateTime) (view *ContentView) {
	return &ContentView{
		ID:          content.ID,
		Title:       content.Title,
		Images:      content.Images,
		Description: content.Description,
		CategoryData: &CategoryData{
			ID:   category.ID,
			Name: category.Name,
		},
		Content:     content.Content,
		LinkOnePage: content.LinkOnePage,
		CreatedByData: &CreatedByData{
			ID:        users.ID,
			FirstName: users.FirstName,
			LastName:  users.LastName,
		},
		CreatedAt: datetime.ConvertUnixToDateTime(content.CreatedAt),
		UpdatedAt: datetime.ConvertUnixToDateTime(content.UpdatedAt),
	}

	return view

}
