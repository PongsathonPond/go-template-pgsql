package inout

import (
	"idev-cms-service/config"
	"idev-cms-service/domain"
	"idev-cms-service/service/util"
)

type AccessTokenUpdate struct {
	AccessToken string `bson:"access_token"`
}

func (upd *AccessTokenUpdate) ToDomain(time util.DateTime, config *config.Config) (token *domain.Tokens) {
	return &domain.Tokens{
		AccessToken:          upd.AccessToken,
		AccessTokenExpiredAt: time.GetNow().AddMinutes(config.AccessTokenTTL).Unix(),
		UpdatedAt:            time.GetUnixNow(),
	}
}
