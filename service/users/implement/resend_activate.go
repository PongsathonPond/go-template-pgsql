package implement

import (
	"context"

	"idev-cms-service/domain"
	"idev-cms-service/service/users/inout"
	"idev-cms-service/service/util"
)

func (impl *implementation) ResendActivate(ctx context.Context, input *inout.ResendActivateInput) (err error) {
	if err = impl.Validator.Validate(input); err != nil {
		return util.ValidationActivationErr(err)
	}

	user := &domain.Users{}
	filters := []string{
		impl.FilterString.MakeID(input.UserID),
		impl.FilterString.MakeVerifiedBool(false),
		impl.FilterString.MakeDeletedAtIsNull(),
	}
	now := impl.DateTime.GetUnixNow()

	if err = impl.RepoUsers.Read(ctx, filters, user); err != nil {
		return util.RepoReadErr(err)
	}

	//if user.ActivateExpiredAt > now {
	//	return util.ConditionActivateExpiredAtErr(errors.New(EmailActivateExpiredAtError))
	//}
	//
	//user.ActivateToken = impl.GenCode.Generate(util.TokenCode, 120)
	//user.ActivateExpiredAt = now + int64(impl.Config.ActivateTokenTTL*60*60)
	user.UpdatedAt = now

	if err = impl.RepoUsers.Update(ctx, filters, user); err != nil {
		return util.RepoUpdateErr(err)
	}

	return nil
}
