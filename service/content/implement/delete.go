package implement

import (
	"context"

	"idev-cms-service/domain"
	"idev-cms-service/service/content/inout"
	"idev-cms-service/service/util"
)

func (impl *implementation) Delete(ctx context.Context, input *inout.ContentDeleteInput) (err error) {
	filters := []string{
		impl.FilterString.MakeID(input.ID),
		impl.FilterString.MakeDeletedAtIsNull(),
	}

	if err = impl.RepoContent.Read(ctx, filters, &domain.Content{}); err != nil {
		return util.RepoReadErr(err)
	}

	if err = impl.RepoContent.SoftDelete(ctx, filters); err != nil {
		return util.RepoDeleteErr(err)
	}

	return nil
}
