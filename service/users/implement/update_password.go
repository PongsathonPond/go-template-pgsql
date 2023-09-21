package implement

import (
	"context"

	"idev-cms-service/domain"
	"idev-cms-service/service/users/inout"

	"go.mongodb.org/mongo-driver/bson"
)

func (impl *implementation) UpdatePassword(ctx context.Context, input *inout.UserUpdatePasswordInput) (err error) {
	filters := []string{
		impl.FilterString.MakeEmail(input.Email),
		impl.FilterString.MakeStatus("active"),
		impl.FilterString.MakeDeletedAtIsNull(),
	}

	if err = impl.RepoUsers.Read(ctx, filters, &domain.Users{}); err != nil {
		return err
	}

	update := &domain.UsersPassword{
		Password:  input.Password,
		UpdatedAt: impl.DateTime.GetUnixNow(),
	}

	if err = impl.RepoUsers.UpdateBson(ctx, filters, bson.D{{"$set", update}}); err != nil {
		return err
	}

	return nil
}
