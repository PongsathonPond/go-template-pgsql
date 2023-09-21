package implement

import (
	"context"
	"errors"

	"idev-cms-service/domain"
	"idev-cms-service/service/users/inout"
	"idev-cms-service/service/util"

	"go.mongodb.org/mongo-driver/bson"
)

func (impl *implementation) UserChangePassword(ctx context.Context, input *inout.UserChangePasswordInput) (err error) {
	if err = impl.Validator.Validate(input); err != nil {
		return util.ValidationUpdateErr(err)
	}

	user := &domain.Users{}

	filters := []string{
		impl.FilterString.MakeID(input.UserID),
		impl.FilterString.MakeDeletedAtIsNull(),
	}

	if err = impl.RepoUsers.Read(ctx, filters, user); err != nil {
		return util.RepoReadErr(err)
	}

	if !impl.Encrypt.CheckPasswordHash(input.OldPassword, user.Password) {
		return util.RepoReadErr(errors.New("password not match"))
	}

	update := &domain.UsersPassword{
		Password:  impl.Encrypt.HashPassword(input.NewPassword),
		UpdatedAt: impl.DateTime.GetUnixNow(),
	}

	if err = impl.RepoUsers.UpdateBson(ctx, filters, bson.D{{"$set", update}}); err != nil {
		return err
	}

	return nil
}
