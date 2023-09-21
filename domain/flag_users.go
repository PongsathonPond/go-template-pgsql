package domain

type FlagUser struct {
	ID         string   `bson:"_id,omitempty"`
	UserIDs    []string `bson:"user_ids"`
	UserGroups string   `bson:"user_groups"`
	Status     string   `bson:"status"`
	CreatedAt  int64    `bson:"created_at"`
}
