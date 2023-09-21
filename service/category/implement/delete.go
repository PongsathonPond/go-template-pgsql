package implement

import (
	"context"

	"idev-cms-service/domain"
	"idev-cms-service/service/category/inout"
	"idev-cms-service/service/util"
)

func (impl *implementation) Delete(ctx context.Context, input *inout.CategoryDeleteInput) (err error) {
	filters := []string{
		impl.FilterString.MakeID(input.ID),
		impl.FilterString.MakeDeletedAtIsNull(),
	}

	if err = impl.RepoCategory.Read(ctx, filters, &domain.Category{}); err != nil {
		return util.RepoReadErr(err)
	}

	if err = impl.RepoCategory.SoftDelete(ctx, filters); err != nil {
		return util.RepoDeleteErr(err)
	}

	return nil
}
