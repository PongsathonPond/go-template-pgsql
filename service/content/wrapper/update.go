package wrapper

import (
	"context"
	"github.com/opentracing/opentracing-go"

	"idev-cms-service/service/content/inout"
)

func (wrp *Wrapper) Update(ctx context.Context, input *inout.ContentUpdateInput) (err error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "service.content.Update")
	defer sp.Finish()

	sp.LogKV("ID", input.ID)
	sp.LogKV("Title", input.Title)
	sp.LogKV("LinkOnePage", input.LinkOnePage)

	err = wrp.Service.Update(ctx, input)

	sp.LogKV("Error", err)

	return err
}
