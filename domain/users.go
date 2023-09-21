package domain

type Users struct {
	ID           string `bson:"id,omitempty" gorm:"primaryKey"`
	FirstName    string `bson:"first_name"`
	LastName     string `bson:"last_name"`
	Email        string `bson:"email"`
	Mobile       string `bson:"mobileac"`
	Username     string `bson:"username"`
	Password     string `bson:"password"`
	Role         string `bson:"role"`
	ImageProfile string `bson:"image_profile"`
	Status       string `bson:"status"`
	CreatedBy    string `bson:"created_by,omitempty"`
	CreatedAt    int64  `bson:"created_at"`
	UpdatedAt    int64  `bson:"updated_at"`
	DeletedAt    int64  `bson:"deleted_at,omitempty"`
}

type UsersVerify struct {
	Verified  bool  `bson:"verified"`
	UpdatedAt int64 `bson:"updated_at"`
}

type UsersPassword struct {
	Password  string `bson:"password"`
	UpdatedAt int64  `bson:"updated_at"`
}

type UsersActivate struct {
	Verified  bool   `bson:"verified"`
	Password  string `bson:"password"`
	UpdatedAt int64  `bson:"updated_at"`
}
