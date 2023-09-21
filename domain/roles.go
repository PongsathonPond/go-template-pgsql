package domain

type Roles struct {
	ID          string            `bson:"_id,omitempty"`
	Name        string            `bson:"name"`
	Description string            `bson:"description"`
	Permissions []MenuPermissions `bson:"permissions"`
	Status      string            `bson:"status"`
	CreatedBy   string            `bson:"created_by"`
	CreatedAt   int64             `bson:"created_at"`
	UpdatedAt   int64             `bson:"updated_at"`
	DeletedAt   int64             `bson:"deleted_at,omitempty"`
}

type MenuPermissions struct {
	KeySlug    string `bson:"key_slug"`
	Permission uint8  `bson:"permission"`
}
