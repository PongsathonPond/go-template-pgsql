package wrapper

import (
	"context"
	"github.com/opentracing/opentracing-go"

	"idev-cms-service/service/content/inout"
)

func (wrp *Wrapper) Create(ctx context.Context, input *inout.ContentCreateInput) (id string, err error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "service.content.Create")
	defer sp.Finish()

	sp.LogKV("Title", input.Title)
	sp.LogKV("LinkOnePage", input.LinkOnePage)

	id, err = wrp.Service.Create(ctx, input)

	sp.LogKV("ID", id)
	sp.LogKV("Error", err)

	return id, err
}
