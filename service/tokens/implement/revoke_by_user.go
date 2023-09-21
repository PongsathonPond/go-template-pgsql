package implement

import (
	"context"
)

func (impl *implementation) RevokeTokenByUserId(ctx context.Context, userID *string) (err error) {
	//filters := []string{
	//	impl.FilterString.MakeUserID(*userID),
	//}
	//_, items, err := impl.RepoTokens.Find(ctx, filters, nil, &domain.Tokens{})
	//if err != nil {
	//	return util.RepoListErr(err)
	//}
	//
	//for _, item := range items {
	//	tokenItem := item.(*domain.Tokens)
	//	_ = impl.RepoTokens.Delete(ctx, impl.FilterString.MakeIDFilters(tokenItem.ID))
	//	_ = impl.RepoRedis.Clear(ctx, fmt.Sprintf("%s-%s", PREFIX_AUTH_TOKEN, impl.md5(tokenItem.AccessToken)))
	//}

	return nil
}
