package implement

import (
	"context"

	"idev-cms-service/domain"
	"idev-cms-service/service/roles/inout"
	"idev-cms-service/service/util"
)

func (impl *implementation) All(ctx context.Context, input *inout.RoleAllInput) (items []*inout.RolesAllView, err error) {
	input.Filters = append(
		input.Filters,
		impl.FilterString.MakeStatus("active"),
		impl.FilterString.MakeDeletedAtIsNull(),
	)
	sorts := []string{
		"name:asc",
	}

	opt := &domain.QueryOption{
		Filters: input.Filters,
		Search:  nil,
		Sorts:   sorts,
	}
	_, records, err := impl.Repo.Find(ctx, opt, &domain.Roles{})
	if err != nil {
		return nil, util.RepoFindErr(err)
	}

	items = make([]*inout.RolesAllView, len(records))
	for i, record := range records {
		role := record.(*domain.Roles)
		items[i] = inout.RoleAllToView(role, impl.DateTime)
	}

	return items, nil
}
