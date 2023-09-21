package pgsql

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"idev-cms-service/domain"
	"reflect"
)

func (repo *Repository) List(ctx context.Context, opt *domain.PageOption, itemType interface{}) (total int, items []interface{}, err error) {
	var query string
	var args []interface{}
	var optFilter []string

	if opt.Filters != nil && len(opt.Filters) > 0 {
		optFilter = opt.Filters
		query, args = repo.makeFilters(opt.Filters)
	}
	skip := (opt.Page - 1) * opt.PerPage

	total, err = repo.Count(ctx, optFilter)
	if err != nil {
		return 0, nil, err
	}

	data := reflect.New(reflect.SliceOf(reflect.TypeOf(itemType))).Interface()

	db := repo.DB.Table(repo.TBName).Where(query, args...).Offset(skip).Limit(opt.PerPage)
	sort := repo.makeSorts(opt.Sorts)
	if sort != "" {
		db = db.Order(sort)
	}
	err = db.Find(data).Error
	if err != nil {
		return 0, nil, err
	}

	items = repo.interfaceToSlice(data)
	return total, items, nil
}

func (r Repository) Find(ctx context.Context, opt *domain.QueryOption, itemType interface{}) (total int, items []interface{}, err error) {
	//TODO implement me
	panic("implement me")
}

func (repo *Repository) Create(ctx context.Context, data interface{}) (id string, err error) {
	res := repo.DB.Table(repo.TBName).Create(data)
	if res.Error != nil {
		return "", err
	}
	return "", err
}

func (repo *Repository) Read(ctx context.Context, filters []string, out interface{}) (err error) {
	query, args := repo.makeFilters(filters)
	sort := repo.makeSorts([]string{"created_at:desc"})
	return repo.DB.Table(repo.TBName).Where(query, args...).Order(sort).First(out).Error
}

func (repo *Repository) Update(ctx context.Context, filters []string, data interface{}) (err error) {
	query, args := repo.makeFilters(filters)
	return repo.DB.Table(repo.TBName).Where(query, args...).Updates(data).Error
}

func (repo *Repository) Delete(ctx context.Context, filters []string) (err error) {
	query, args := repo.makeFilters(filters)
	return repo.DB.Table(repo.TBName).Where(query, args...).Delete(nil).Error
}

func (repo *Repository) Count(ctx context.Context, filters []string) (total int, err error) {
	query, args := repo.makeFilters(filters)
	var cnt int64
	err = repo.DB.Table(repo.TBName).Where(query, args...).Count(&cnt).Error
	if err != nil {
		return 0, err
	}
	return int(cnt), nil
}

func (r Repository) UpdateBson(ctx context.Context, filters []string, ent bson.D) (err error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) DeleteMany(ctx context.Context, filters []string) (err error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) SoftDelete(ctx context.Context, filters []string) (err error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) Push(ctx context.Context, param *domain.SetOpParam) (err error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) Pop(ctx context.Context, param *domain.SetOpParam) (err error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) IsFirst(ctx context.Context, param *domain.SetOpParam) (is bool, err error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) CountArray(ctx context.Context, param *domain.SetOpParam) (total int, err error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) ClearArray(ctx context.Context, param *domain.SetOpParam) (err error) {
	//TODO implement me
	panic("implement me")
}
