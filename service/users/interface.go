package users

import (
	"context"

	"idev-cms-service/domain"
	"idev-cms-service/service/users/inout"
)

//go:generate mockery --name=Service
type Service interface {
	List(ctx context.Context, opt *domain.PageOption) (total int, items []*inout.UserReadView, err error)
	Create(ctx context.Context, input *inout.UserCreateInput) (id string, err error)
	Read(ctx context.Context, input *inout.ReadInput) (user *inout.UserReadView, err error)
	Update(ctx context.Context, input *inout.UserUpdateInput) (err error)
	Delete(ctx context.Context, input *inout.UserDeleteInput) (err error)
	SetVerify(ctx context.Context, input *inout.SetVerifyInput) (resp *inout.SetVerifyResponse, err error)
	GetMe(ctx context.Context, input *inout.GetMeInput) (user *inout.MeView, err error)
	UpdateMe(ctx context.Context, input *inout.ProfileUpdateInput) (err error)
	ReadGRPC(ctx context.Context, input *inout.ReadInput) (user *inout.UserViewGRPC, err error)
	CountByRole(ctx context.Context, input *inout.CountByRoleInput) (total int, err error)
	ReadByToken(ctx context.Context, token *string) (user *inout.UserViewByToken, err error)
	UpdatePassword(ctx context.Context, input *inout.UserUpdatePasswordInput) (err error)
	Activate(ctx context.Context, input *inout.UserActivateInput) (err error)
	ResendActivate(ctx context.Context, input *inout.ResendActivateInput) (err error)
	CheckEmailExist(ctx context.Context, email *string) (err error)
	VerifyActivateToken(ctx context.Context, input *inout.VerifyActivateTokenInput) (err error)
	UserChangePassword(ctx context.Context, input *inout.UserChangePasswordInput) (err error)
}
