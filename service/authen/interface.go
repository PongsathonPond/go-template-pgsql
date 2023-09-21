package authen

import (
	"context"

	"idev-cms-service/service/authen/inout"
)

//go:generate mockery --name=Service
type Service interface {
	Login(ctx context.Context, input *inout.LoginInput) (view *inout.UserViewWithToken, err error)
	Logout(ctx context.Context, input *inout.LogoutInput) (err error)
	VerifyTokens(ctx context.Context, token *string) (resp *inout.VerifyTokenResponse)
}
