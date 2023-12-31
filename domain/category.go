package domain

type Category struct {
	ID        string `bson:"_id,omitempty"`
	Name      string `bson:"name"`
	CreatedAt int64  `bson:"created_at"`
	UpdatedAt int64  `bson:"updated_at"`
	DeletedAt int64  `bson:"deleted_at,omitempty"`
}
