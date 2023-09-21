package implement

import (
	"context"

	"idev-cms-service/service/category/inout"
	"idev-cms-service/service/util"
)

func (impl *implementation) Create(ctx context.Context, input *inout.CategoryCreateInput) (id string, err error) {
	if err = impl.Validator.Validate(input); err != nil {
		return "", util.ValidationCreateErr(err)
	}

	input.ID = impl.UUID.Generate()
	category := input.ToDomain(impl.DateTime)

	_, err = impl.RepoCategory.Create(ctx, category)
	if err != nil {
		return "", util.RepoCreateErr(err)
	}

	return input.ID, nil
}
