package implement

import (
	"context"

	"idev-cms-service/domain"
	"idev-cms-service/service/authen/inout"
)

func (impl *implementation) VerifyTokens(ctx context.Context, token *string) (resp *inout.VerifyTokenResponse) {
	resp = &inout.VerifyTokenResponse{
		Verify: false,
	}
	var userID string

	// find in redis first
	//err := impl.RepoRedis.Get(ctx, fmt.Sprintf("%s-%s", PREFIX_AUTH_TOKEN, impl.Encrypt.MD5(*token)), &userID)
	//if err == redis.Nil {
	// find in database if not found in redis
	tokenData := &domain.Tokens{}
	filters := []string{
		impl.FilterString.MakeAccessToken(*token),
		impl.FilterString.MakeAccessTokenExpiredAtGreaterThan(impl.DateTime.GetUnixNow()),
	}
	if err := impl.RepoTokens.Read(ctx, filters, tokenData); err != nil {
		return resp
	}
	userID = tokenData.UserID

	// save data to redis
	//var ttl time.Duration
	//diffTTL := time.Duration(tokenData.AccessTokenExpiredAt - impl.DateTime.GetUnixNow())
	//if diffTTL*time.Second >= 10*time.Minute {
	//	ttl = 10 * time.Minute
	//} else if diffTTL*time.Second > 0 {
	//	ttl = diffTTL * time.Second
	//} else {
	//	ttl = 0 * time.Second
	//}
	//_ = impl.RepoRedis.SetWithTTL(ctx, fmt.Sprintf("%s-%s", implement.PREFIX_AUTH_TOKEN, impl.Encrypt.MD5(*token)), tokenData.UserID, ttl)
	//}

	if userID == "" {
		return resp
	}

	// response data
	resp.Verify = true
	resp.UserID = userID

	// get user data'
	userData := &domain.Users{}
	filters = []string{
		impl.FilterString.MakeID(userID),
	}
	if err := impl.RepoUsers.Read(ctx, filters, userData); err != nil {
		return resp
	}

	//resp.UserGroup = userData.UserGroup
	resp.Role = userData.Role
	//resp.UserVerified = userData.Verified

	return resp
}
