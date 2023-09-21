package wrapper

import (
	"context"
	"github.com/opentracing/opentracing-go"

	"idev-cms-service/service/content/inout"
)

func (wrp *Wrapper) Delete(ctx context.Context, input *inout.ContentDeleteInput) (err error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "service.content.Delete")
	defer sp.Finish()

	sp.LogKV("ID", input.ID)

	err = wrp.Service.Delete(ctx, input)

	sp.LogKV("Error", err)

	return err
}
