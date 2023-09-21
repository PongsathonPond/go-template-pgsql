package wrapper

import (
	"context"
	"github.com/opentracing/opentracing-go"

	"idev-cms-service/service/content/inout"
)

func (wrp *Wrapper) Read(ctx context.Context, input *inout.ContentReadInput) (view *inout.ContentView, err error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "service.content.Read")
	defer sp.Finish()

	sp.LogKV("ID", input.ID)

	view, err = wrp.Service.Read(ctx, input)

	sp.LogKV("Data", view)
	sp.LogKV("Err", err)

	return view, err
}
