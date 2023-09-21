package implement

import (
	"context"

	"idev-cms-service/domain"
	"idev-cms-service/service/category/inout"
	"idev-cms-service/service/util"
)

func (impl *implementation) List(ctx context.Context, opt *domain.PageOption) (total int, items []*inout.CategoryView, err error) {
	if err = impl.Validator.Validate(opt); err != nil {
		return 0, nil, util.ValidationParamOptionErr(err)
	}

	total, records, err := impl.RepoCategory.List(ctx, opt, &domain.Category{})
	if err != nil {
		return 0, nil, util.RepoListErr(err)
	}

	items = make([]*inout.CategoryView, len(records))
	for i, record := range records {

		//createdBy := &domain.Users{}
		//filters := []string{
		//	impl.FilterString.MakeID(record.(*domain.Users).CreatedBy),
		//	impl.FilterString.MakeDeletedAtIsNull(),
		//}
		//if err = impl.RepoUsers.Read(ctx, filters, createdBy); err != nil {
		//	createdBy = nil
		//}

		items[i] = inout.CategoryToView(record.(*domain.Category), impl.DateTime)
	}

	return total, items, nil
}
