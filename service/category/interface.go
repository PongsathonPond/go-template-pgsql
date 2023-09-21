package solutions

import (
	"context"

	"idev-cms-service/domain"
	"idev-cms-service/service/category/inout"
)

//go:generate mockery --name=Service
type Service interface {
	List(ctx context.Context, opt *domain.PageOption) (total int, items []*inout.CategoryView, err error)
	Create(ctx context.Context, input *inout.CategoryCreateInput) (ID string, err error)
	Read(ctx context.Context, input *inout.CategoryReadInput) (solutions *inout.CategoryView, err error)
	Update(ctx context.Context, input *inout.CategoryUpdateInput) (err error)
	Delete(ctx context.Context, input *inout.CategoryDeleteInput) (err error)
}
