package implement

import (
	"context"

	"idev-cms-service/domain"
	"idev-cms-service/service/content/inout"
	"idev-cms-service/service/util"
)

func (impl *implementation) Read(ctx context.Context, input *inout.ContentReadInput) (view *inout.ContentView, err error) {
	content := &domain.Content{}
	filters := []string{
		impl.FilterString.MakeID(input.ID),
		impl.FilterString.MakeDeletedAtIsNull(),
	}

	if err = impl.RepoContent.Read(ctx, filters, content); err != nil {
		return nil, util.RepoReadErr(err)
	}

	category := &domain.Category{}

	filtersCategory := []string{
		impl.FilterString.MakeID(content.CategoryID),
		impl.FilterString.MakeDeletedAtIsNull(),
	}
	if err = impl.RepoCategory.Read(ctx, filtersCategory, category); err != nil {
		return nil, util.RepoReadErr(err)
	}

	createdBy := &domain.Users{}
	filters = []string{
		impl.FilterString.MakeDeletedAtIsNull(),
	}

	_ = impl.RepoUsers.Read(ctx, filters, createdBy)

	return inout.ContentToView(content, category, createdBy, impl.DateTime), nil
}
