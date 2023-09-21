package domain

type Content struct {
	ID          string `bson:"_id,omitempty"`
	Title       string `bson:"title"`
	CategoryID  string `bson:"category_id"`
	Images      string `bson:"images"`
	Description string `bson:"description"`
	Content     string `bson:"content"`
	LinkOnePage string `bson:"link_one_page"`
	CreatedBy   string `bson:"created_by,omitempty"`
	CreatedAt   int64  `bson:"created_at"`
	UpdatedAt   int64  `bson:"updated_at"`
	DeletedAt   int64  `bson:"deleted_at,omitempty"`
}
