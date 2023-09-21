package implement

import (
	"context"

	"idev-cms-service/domain"
	"idev-cms-service/service/category/inout"
	"idev-cms-service/service/util"

	"github.com/imdario/mergo"
)

func (impl *implementation) Update(ctx context.Context, input *inout.CategoryUpdateInput) (err error) {
	if err = impl.Validator.Validate(input); err != nil {
		return util.ValidationUpdateErr(err)
	}

	category := &domain.Category{}
	filters := impl.FilterString.MakeIDFilters(input.ID)

	if err = impl.RepoCategory.Read(ctx, filters, category); err != nil {
		return util.RepoReadErr(err)
	}

	update := input.ToDomain(impl.DateTime)

	if err = mergo.Merge(category, update, mergo.WithOverride); err != nil {
		return util.UnknownErr(err)
	}

	if err = impl.RepoCategory.Update(ctx, filters, category); err != nil {
		return util.RepoUpdateErr(err)
	}

	return nil
}
