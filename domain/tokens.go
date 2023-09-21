package domain

type Tokens struct {
	ID                    string `bson:"_id,omitempty"`
	UserID                string `bson:"user_id,omitempty"`
	UserAgent             string `bson:"user_agent,omitempty"`
	AccessToken           string `bson:"access_token"`
	RefreshToken          string `bson:"refresh_token,omitempty"`
	AccessTokenExpiredAt  int64  `bson:"access_token_exp_at"`
	RefreshTokenExpiredAt int64  `bson:"refresh_token_exp_at,omitempty"`
	CreatedAt             int64  `bson:"created_at,omitempty"`
	UpdatedAt             int64  `bson:"updated_at"`
}
