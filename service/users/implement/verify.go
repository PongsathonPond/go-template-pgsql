package implement

import (
	"context"

	"idev-cms-service/domain"
	"idev-cms-service/service/users/inout"

	"go.mongodb.org/mongo-driver/bson"
)

func (impl *implementation) SetVerify(ctx context.Context, input *inout.SetVerifyInput) (resp *inout.SetVerifyResponse, err error) {
	token := &domain.Tokens{}
	filters := []string{
		impl.FilterString.MakeAccessToken(input.Token),
	}
	if err = impl.RepoTokens.Read(ctx, filters, token); err != nil {
		return nil, err
	}

	// set user verify
	filters = []string{
		impl.FilterString.MakeID(token.UserID),
	}
	update := &domain.UsersVerify{
		Verified:  true,
		UpdatedAt: impl.DateTime.GetUnixNow(),
	}
	if err = impl.RepoUsers.UpdateBson(ctx, filters, bson.D{{"$set", update}}); err != nil {
		return nil, err
	}

	return &inout.SetVerifyResponse{
		ID:       token.UserID,
		Email:    input.Email,
		Verified: true,
	}, nil
}
