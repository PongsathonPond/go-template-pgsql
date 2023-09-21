package wrapper

import (
	"context"

	"github.com/opentracing/opentracing-go"
)

func (wrp *Wrapper) GetUserIDByToken(ctx context.Context, token *string) (userID string) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "service.authen.UserIDByToken")
	defer sp.Finish()

	sp.LogKV("Token", *token)

	userID = wrp.Service.GetUserIDByToken(ctx, token)

	sp.LogKV("UserID", userID)

	return userID
}
