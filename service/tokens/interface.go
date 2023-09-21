package tokens

import (
	"context"

	"idev-cms-service/service/tokens/inout"
)

//go:generate mockery --name=Service
type Service interface {
	RefreshToken(ctx context.Context, input *inout.RefreshTokenInput) (view *inout.RefreshTokenView, err error)
	RevokeToken(ctx context.Context, input *inout.RevokeTokenInput) (err error)
	RevokeTokenByUserId(ctx context.Context, userID *string) (err error)
	GetUserIDByToken(ctx context.Context, token *string) (userID string)
}
