package implement

import (
	"context"

	"idev-cms-service/domain"
	"idev-cms-service/service/tokens/inout"
	"idev-cms-service/service/util"
)

func (impl *implementation) RevokeToken(ctx context.Context, input *inout.RevokeTokenInput) (err error) {
	if err = impl.Validator.Validate(input); err != nil {
		return util.ValidationErr(err)
	}

	token := &domain.Tokens{}
	filters := []string{
		impl.FilterString.MakeAccessToken(input.AccessToken),
	}
	err = impl.RepoTokens.Read(ctx, filters, token)
	if err != nil {
		return util.RepoReadErr(err)
	}

	//clear token from database
	opt := &domain.QueryOption{
		Filters: []string{
			impl.FilterString.MakeUserID(token.UserID),
		},
		Search: nil,
		Sorts:  nil,
	}
	_, items, err := impl.RepoTokens.Find(ctx, opt, &domain.Tokens{})
	if err != nil {
		return util.RepoListErr(err)
	}

	for _, item := range items {
		tokenItem := item.(*domain.Tokens)
		filters = []string{
			impl.FilterString.MakeID(tokenItem.ID),
		}
		if err = impl.RepoTokens.Delete(ctx, filters); err != nil {
			return util.RepoDeleteErr(err)
		}
	}

	return nil
}
