package inout

import (
	"idev-cms-service/config"
	"idev-cms-service/domain"
	"idev-cms-service/service/util"
)

type TokenInput struct {
	ID           string `bson:"_id"`
	UserID       string `bson:"user_id"`
	UserAgent    string `bson:"user_agent"`
	AccessToken  string `bson:"access_token"`
	RefreshToken string `bson:"refresh_token"`
}

type TokenUpdate struct {
	AccessToken  string `bson:"access_token"`
	RefreshToken string `bson:"refresh_token"`
}

func (inp *TokenInput) ToDomain(time util.DateTime, config *config.Config) (token *domain.Tokens) {
	return &domain.Tokens{
		ID:                    inp.ID,
		UserID:                inp.UserID,
		UserAgent:             inp.UserAgent,
		AccessToken:           inp.AccessToken,
		RefreshToken:          inp.RefreshToken,
		AccessTokenExpiredAt:  time.GetNow().AddMinutes(config.AccessTokenTTL).Unix(),
		RefreshTokenExpiredAt: time.GetNow().AddMinutes(config.RefreshTokenTTL).Unix(),
		CreatedAt:             time.GetUnixNow(),
		UpdatedAt:             time.GetUnixNow(),
	}
}

func (upd *TokenUpdate) ToDomain(time util.DateTime, config *config.Config) (token *domain.Tokens) {
	return &domain.Tokens{
		AccessToken:           upd.AccessToken,
		RefreshToken:          upd.RefreshToken,
		AccessTokenExpiredAt:  time.GetNow().AddMinutes(config.AccessTokenTTL).Unix(),
		RefreshTokenExpiredAt: time.GetNow().AddMinutes(config.RefreshTokenTTL).Unix(),
		UpdatedAt:             time.GetUnixNow(),
	}
}
