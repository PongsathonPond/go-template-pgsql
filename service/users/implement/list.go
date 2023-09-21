package implement

import (
	"context"
	"idev-cms-service/domain"
	"idev-cms-service/service/users/inout"
	"idev-cms-service/service/util"
)

func (impl *implementation) List(ctx context.Context, opt *domain.PageOption) (total int, items []*inout.UserReadView, err error) {
	if err = impl.Validator.Validate(opt); err != nil {
		return 0, nil, util.ValidationParamOptionErr(err)
	}

	total, records, err := impl.RepoUsers.List(ctx, opt, &domain.Users{})
	if err != nil {
		return 0, nil, util.RepoListErr(err)
	}

	items = make([]*inout.UserReadView, len(records))
	for i, record := range records {
		user := record.(*domain.Users)

		role := &domain.Roles{}
		filters := []string{
			impl.FilterString.MakeID(record.(*domain.Users).Role),
			//impl.FilterString.MakeDeletedAtIsNull(),

		}
		if err = impl.RepoRoles.Read(ctx, filters, role); err != nil {
			role = nil
		}
		createdBy := &domain.Users{}
		items[i] = inout.UserToView(user, role, createdBy, impl.DateTime)
	}

	return total, items, nil
}
