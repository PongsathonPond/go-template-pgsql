package implement

import (
	"context"

	"idev-cms-service/domain"
)

func (impl *implementation) GetUserIDByToken(ctx context.Context, token *string) (userID string) {
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
		return ""
	}
	return tokenData.UserID
	//}
	//return userID
}
