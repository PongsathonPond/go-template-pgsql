package implement

import (
	"context"

	"idev-cms-service/service/users/inout"
)

func (impl *implementation) Activate(ctx context.Context, input *inout.UserActivateInput) (err error) {
	//if err = impl.Validator.Validate(input); err != nil {
	//	return util.ValidationUpdateErr(err)
	//}
	//
	//user := &domain.Users{}
	//filters := []string{
	//	impl.FilterString.MakeActivateTokenString(input.ActivateToken),
	//	impl.FilterString.MakeVerifiedBool(false),
	//	impl.FilterString.MakeDeletedAtIsNull(),
	//}
	//now := impl.DateTime.GetUnixNow()
	//
	//err = impl.RepoUsers.Read(ctx, filters, user)
	//if err != nil {
	//	return util.RepoReadErr(err)
	//}
	//
	//if user.ActivateExpiredAt < now {
	//	return util.ConditionActivateExpiredAtErr(errors.New(ActivateTokenHasExpired))
	//}
	//
	//update := &domain.UsersActivate{
	//	Verified:  true,
	//	Password:  impl.Encrypt.HashPassword(input.Password),
	//	UpdatedAt: impl.DateTime.GetUnixNow(),
	//}
	//
	//bsonSet := bson.D{
	//	{"$set", update},
	//	{"$unset", bson.D{{"activate_token", nil}, {"activate_expired_at", nil}}},
	//}
	//
	//filters = []string{
	//	impl.FilterString.MakeID(user.ID),
	//}
	//
	//if err = impl.RepoUsers.UpdateBson(ctx, filters, bsonSet); err != nil {
	//	return util.RepoUpdateErr(err)
	//}

	return nil
}
