package implement

import (
	"context"
	"errors"
	"fmt"

	"idev-cms-service/domain"
	"idev-cms-service/service/tokens/inout"
	"idev-cms-service/service/util"
)

func (impl *implementation) RefreshToken(ctx context.Context, input *inout.RefreshTokenInput) (view *inout.RefreshTokenView, err error) {
	if err = impl.Validator.Validate(input); err != nil {
		return nil, util.ValidationErr(err)
	}

	token := &domain.Tokens{}
	filters := []string{
		impl.FilterString.MakeRefreshToken(input.RefreshToken),
	}
	if err = impl.RepoTokens.Read(ctx, filters, token); err != nil {
		return nil, util.RepoReadErr(err)
	}

	// check refresh token has expired
	if token.RefreshTokenExpiredAt < impl.DateTime.GetUnixNow() {
		return nil, util.ValidationTokenExpired(errors.New("token has expired"))
	}

	user := &domain.Users{}
	filters = []string{
		impl.FilterString.MakeID(token.UserID),
	}
	if err = impl.RepoUsers.Read(ctx, filters, user); err != nil {
		return nil, util.RepoReadErr(err)
	}

	if user.Status != "active" {
		return nil, util.RepoReadErr(errors.New(fmt.Sprintf("user status is %s", user.Status)))
	}

	// generate new access token
	accessToken, err := impl.generateAccessToken(user)
	if err != nil {
		return nil, util.RepoReadErr(err)
	}

	tokenData := inout.AccessTokenUpdate{
		AccessToken: *accessToken,
	}
	filters = []string{
		impl.FilterString.MakeID(token.ID),
	}
	if err = impl.RepoTokens.Update(ctx, filters, tokenData.ToDomain(impl.DateTime, impl.Config)); err != nil {
		return nil, util.RepoUpdateErr(err)
	}

	return inout.RefreshTokenToView(impl.Config, accessToken), nil
}
