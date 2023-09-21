package implement

import (
	"context"
	"log"

	"idev-cms-service/domain"
	"idev-cms-service/service/menus/inout"
	"idev-cms-service/service/util"
)

func (impl *implementation) All(ctx context.Context) (total int, items []*inout.MenuView, err error) {
	//var grpcUserGroup = &pb.UserGroupByKeyReply{}
	//if userGroups != "" { // get user group data by keySlug (gRPC)
	//	grpcUserGroup, err = impl.GrpcAuthenClientService.GetUserGroupByKey(&pb.UserGroupByKeyRequest{KeySlug: userGroups})
	//	if err != nil {
	//		impl.Log.Error(err)
	//	}
	//	// if user group not exits
	//	if !grpcUserGroup.Success {
	//		return 0, nil, util.RepoListErr(err)
	//	}
	//}

	filters := []string{
		impl.FilterString.MakeStatus("active"),
		impl.FilterString.MakeDeletedAtIsNull(),
	}
	sorts := []string{
		"sort:asc",
	}

	opt := &domain.QueryOption{
		Filters: filters,
		Search:  nil,
		Sorts:   sorts,
	}
	total, records, err := impl.Repo.Find(ctx, opt, &domain.Menu{})
	if err != nil {
		log.Println(err)
		return 0, nil, util.RepoListErr(err)
	}

	items = make([]*inout.MenuView, len(records))
	for i, record := range records {
		rec := record.(*domain.Menu)

		//if rec.KeySlug == "cms" && strings.Contains(strings.ToLower(userGroups), Depa) {
		//	for _, sub := range rec.SubMenus {
		//		if sub.KeySlug == "cctv" {
		//			sub.CanSetPermission = 6
		//			break
		//		}
		//	}
		//}

		items[i] = inout.MenuToView(rec, impl.DateTime, impl.Perm)
	}

	return total, items, nil
}
