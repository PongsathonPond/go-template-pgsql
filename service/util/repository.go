package util

import (
	"context"
	"time"

	"idev-cms-service/domain"

	"go.mongodb.org/mongo-driver/bson"
)

//go:generate mockery --name=Repository
type Repository interface {
	List(ctx context.Context, opt *domain.PageOption, itemType interface{}) (total int, items []interface{}, err error)
	Find(ctx context.Context, opt *domain.QueryOption, itemType interface{}) (total int, items []interface{}, err error)
	Create(ctx context.Context, ent interface{}) (id string, err error)
	Read(ctx context.Context, filters []string, out interface{}) (err error)
	Update(ctx context.Context, filters []string, ent interface{}) (err error)
	UpdateBson(ctx context.Context, filters []string, ent bson.D) (err error)
	Delete(ctx context.Context, filters []string) (err error)
	DeleteMany(ctx context.Context, filters []string) (err error)
	SoftDelete(ctx context.Context, filters []string) (err error)
	Count(ctx context.Context, filters []string) (total int, err error)
	Push(ctx context.Context, param *domain.SetOpParam) (err error)
	Pop(ctx context.Context, param *domain.SetOpParam) (err error)
	IsFirst(ctx context.Context, param *domain.SetOpParam) (is bool, err error)
	CountArray(ctx context.Context, param *domain.SetOpParam) (total int, err error)
	ClearArray(ctx context.Context, param *domain.SetOpParam) (err error)
}

//go:generate mockery --name=RepositoryRedis
type RepositoryRedis interface {
	Set(ctx context.Context, key string, value interface{}) (err error)
	SetWithTTL(ctx context.Context, key string, value interface{}, ttl time.Duration) (err error)
	Get(ctx context.Context, key string, dest interface{}) (err error)
	GetTTL(ctx context.Context, key string) (result time.Duration, err error)
	ExpireAt(ctx context.Context, key string, tm time.Time) (err error)
	Clear(ctx context.Context, keys ...string) (err error)
}
