package implement

import (
	"context"

	"idev-cms-service/domain"
	"idev-cms-service/service/category/inout"
	"idev-cms-service/service/util"
)

func (impl *implementation) Read(ctx context.Context, input *inout.CategoryReadInput) (view *inout.CategoryView, err error) {
	category := &domain.Category{}
	filters := []string{
		impl.FilterString.MakeID(input.ID),
		impl.FilterString.MakeDeletedAtIsNull(),
	}

	if err = impl.RepoCategory.Read(ctx, filters, category); err != nil {
		return nil, util.RepoReadErr(err)
	}

	//createdBy := &domain.Users{}
	//filters = []string{
	//	impl.FilterString.MakeDeletedAtIsNull(),
	//}
	//
	//_ = impl.RepoUsers.Read(ctx, filters, createdBy)

	return inout.CategoryToView(category, impl.DateTime), nil
}
