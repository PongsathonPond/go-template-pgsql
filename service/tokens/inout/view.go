package inout

import (
	"time"

	"idev-cms-service/config"
)

type RefreshTokenView struct {
	AccessToken    string `json:"access_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXvCj9..."`
	AccessTokenTTL int64  `json:"access_token_expires" example:"3600"`
} // @Name RefreshTokenView

func RefreshTokenToView(config *config.Config, accessToken *string) (view *RefreshTokenView) {
	return &RefreshTokenView{
		AccessToken:    *accessToken,
		AccessTokenTTL: int64((time.Duration(config.AccessTokenTTL) * time.Minute).Seconds()),
	}
}
