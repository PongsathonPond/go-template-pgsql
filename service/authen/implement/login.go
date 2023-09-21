package implement

import (
	"context"
	"errors"
	"fmt"

	"idev-cms-service/domain"
	"idev-cms-service/service/authen/inout"
	"idev-cms-service/service/util"
)

func (impl *implementation) Login(ctx context.Context, input *inout.LoginInput) (view *inout.UserViewWithToken, err error) {
	if err = impl.Validator.Validate(input); err != nil {
		return nil, util.ValidationErr(err)
	}

	user := &domain.Users{}
	filters := []string{
		impl.FilterString.MakeUserName(input.Username),
		impl.FilterString.MakeStatusIn("active", "inactive"),
	}

	if err = impl.RepoUsers.Read(ctx, filters, user); err != nil {
		return nil, util.RepoReadErr(err)
	}

	if !impl.Encrypt.CheckPasswordHash(input.Password, user.Password) {
		return nil, util.RepoReadErr(errors.New("username or password not match"))
	}

	if user.Status != "active" {
		return nil, util.RepoReadErr(errors.New(fmt.Sprintf("user status is %s", user.Status)))
	}

	accessToken, refreshToken, err := impl.generateTokens(user)
	if err != nil {
		return nil, util.RepoReadErr(err)
	}

	// save token to database
	filters = []string{
		impl.FilterString.MakeUserID(user.ID),
		impl.FilterString.MakeUserAgent(input.UserAgent),
	}
	cntToken, err := impl.RepoTokens.Count(ctx, filters)
	if err != nil {
		return nil, util.RepoReadErr(err)
	}
	if cntToken > 0 {
		tokenData := inout.TokenUpdate{
			AccessToken:  *accessToken,
			RefreshToken: *refreshToken,
		}
		if err = impl.RepoTokens.Update(ctx, filters, tokenData.ToDomain(impl.DateTime, impl.Config)); err != nil {
			return nil, util.RepoUpdateErr(err)
		}
	} else {
		tokenData := inout.TokenInput{
			ID:           impl.UUID.Generate(),
			UserID:       user.ID,
			UserAgent:    input.UserAgent,
			AccessToken:  *accessToken,
			RefreshToken: *refreshToken,
		}
		_, err = impl.RepoTokens.Create(ctx, tokenData.ToDomain(impl.DateTime, impl.Config))
		if err != nil {
			return nil, util.RepoCreateErr(err)
		}
	}

	// save token to redis
	//if err = impl.RepoRedis.SetWithTTL(ctx, fmt.Sprintf("%s-%s", PREFIX_AUTH_TOKEN, impl.Encrypt.MD5(*accessToken)), user.ID, 10*time.Minute); err != nil {
	//	impl.Log.Error(err)
	//}

	// get user group data
	//userGroup := &domain.UserGroups{}
	filters = []string{
		//impl.FilterString.MakeKeySlug(user.UserGroup),
	}
	//if err := impl.RepoUserGroups.Read(ctx, filters, userGroup); err != nil {
	//	return nil, util.RepoReadErr(err)
	//}

	// get user role data (gRPC)
	role := &domain.Roles{}

	return inout.UserWithTokenToView(user, role, impl.DateTime, impl.Config, accessToken, refreshToken), nil
}
