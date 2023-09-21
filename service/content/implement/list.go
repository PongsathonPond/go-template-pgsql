package implement

import (
	"context"
	"idev-cms-service/domain"
	"idev-cms-service/service/content/inout"
	"idev-cms-service/service/util"
	"log"
)

func (impl *implementation) List(ctx context.Context, opt *domain.PageOption) (total int, items []*inout.ContentView, err error) {
	if err = impl.Validator.Validate(opt); err != nil {
		return 0, nil, util.ValidationParamOptionErr(err)
	}

	total, records, err := impl.RepoContent.List(ctx, opt, &domain.Content{})
	if err != nil {
		return 0, nil, util.RepoListErr(err)
	}

	items = make([]*inout.ContentView, len(records))
	for i, record := range records {
		content := record.(*domain.Content)
		category := &domain.Category{}
		users := &domain.Users{}

		filters := []string{
			impl.FilterString.MakeID(record.(*domain.Content).CategoryID),
			impl.FilterString.MakeDeletedAtIsNull(),
		}

		if err = impl.RepoCategory.Read(ctx, filters, category); err != nil {
			category = nil
		}

		filtersUser := []string{
			impl.FilterString.MakeID(record.(*domain.Content).CreatedBy),
			impl.FilterString.MakeDeletedAtIsNull(),
		}

		if err = impl.RepoUsers.Read(ctx, filtersUser, users); err != nil {
			users = nil
		}
		log.Println(users)

		items[i] = inout.ContentToView(content, category, users, impl.DateTime)
	}

	return total, items, nil
}
