package implement

import (
	"context"
	"github.com/imdario/mergo"
	"idev-cms-service/domain"
	"idev-cms-service/service/content/inout"
	"idev-cms-service/service/util"
)

func (impl *implementation) Update(ctx context.Context, input *inout.ContentUpdateInput) (err error) {
	if err = impl.Validator.Validate(input); err != nil {
		return util.ValidationUpdateErr(err)
	}

	content := &domain.Content{}
	filters := impl.FilterString.MakeIDFilters(input.ID)

	if err = impl.RepoContent.Read(ctx, filters, content); err != nil {
		return util.RepoReadErr(err)
	}

	update := input.ToDomain(impl.DateTime)

	if err = mergo.Merge(content, update, mergo.WithOverride); err != nil {
		return util.UnknownErr(err)
	}

	if input.LinkOnePage == "" {
		content.LinkOnePage = ""
	}
	if input.Content == "" {
		content.Content = ""
	}
	if input.Description == "" {
		content.Description = ""
	}
	if input.Images == "" {
		content.Images = ""
	}

	if err = impl.RepoContent.Update(ctx, filters, content); err != nil {
		return util.RepoUpdateErr(err)
	}

	return nil
}
