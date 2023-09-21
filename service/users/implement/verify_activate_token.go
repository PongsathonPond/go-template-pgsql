package implement

import (
	"context"

	"idev-cms-service/domain"
	"idev-cms-service/service/users/inout"
	"idev-cms-service/service/util"
)

func (impl *implementation) VerifyActivateToken(ctx context.Context, input *inout.VerifyActivateTokenInput) (err error) {
	if err = impl.Validator.Validate(input); err != nil {
		return util.ValidationUpdateErr(err)
	}

	user := &domain.Users{}
	filters := []string{
		impl.FilterString.MakeActivateTokenString(input.ActivateToken),
		impl.FilterString.MakeVerifiedBool(false),
		impl.FilterString.MakeDeletedAtIsNull(),
	}
	//now := impl.DateTime.GetUnixNow()

	err = impl.RepoUsers.Read(ctx, filters, user)
	if err != nil {
		return util.RepoReadErr(err)
	}

	//if user.ActivateExpiredAt < now {
	//	return util.ConditionActivateExpiredAtErr(errors.New(ActivateTokenHasExpired))
	//}

	return nil
}
