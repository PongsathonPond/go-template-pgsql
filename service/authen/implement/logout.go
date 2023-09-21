package implement

import (
	"context"
	"fmt"

	"idev-cms-service/domain"
	"idev-cms-service/service/authen/inout"
	"idev-cms-service/service/util"
)

func (impl *implementation) Logout(ctx context.Context, input *inout.LogoutInput) (err error) {
	if err = impl.Validator.Validate(input); err != nil {
		return util.ValidationErr(err)
	}

	filters := []string{
		impl.FilterString.MakeAccessToken(input.AccessToken),
	}
	err = impl.RepoTokens.Read(ctx, filters, &domain.Tokens{})
	if err != nil {
		return util.RepoReadErr(err)
	}

	// clear token from database
	if err = impl.RepoTokens.Delete(ctx, filters); err != nil {
		return util.RepoDeleteErr(err)
	}

	// clear token from redis
	_ = impl.RepoRedis.Clear(ctx, fmt.Sprintf("%s-%s", PREFIX_AUTH_TOKEN, impl.Encrypt.MD5(input.AccessToken)))

	return nil
}
