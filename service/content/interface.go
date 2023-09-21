package solutions

import (
	"context"

	"idev-cms-service/domain"
	"idev-cms-service/service/content/inout"
)

//go:generate mockery --name=Service
type Service interface {
	List(ctx context.Context, opt *domain.PageOption) (total int, items []*inout.ContentView, err error)
	Create(ctx context.Context, input *inout.ContentCreateInput) (ID string, err error)
	Read(ctx context.Context, input *inout.ContentReadInput) (solutions *inout.ContentView, err error)
	Update(ctx context.Context, input *inout.ContentUpdateInput) (err error)
	Delete(ctx context.Context, input *inout.ContentDeleteInput) (err error)
}
