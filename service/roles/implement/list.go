package implement

import (
	"context"

	"idev-cms-service/domain"
	"idev-cms-service/service/roles/inout"
	"idev-cms-service/service/util"
)

func (impl *implementation) List(ctx context.Context, opt *domain.PageOption) (total int, items []*inout.RolesListView, err error) {
	if err = impl.Validator.Validate(opt); err != nil {
		return 0, nil, util.ValidationParamOptionErr(err)
	}

	//opt.Filters = append(opt.Filters, impl.FilterString.MakeIsSuperRoleIsNull())

	total, records, err := impl.Repo.List(ctx, opt, &domain.Roles{})
	if err != nil {
		return 0, nil, util.RepoListErr(err)
	}

	items = make([]*inout.RolesListView, len(records))
	for i, record := range records {
		role := record.(*domain.Roles)

		// get user group data (gRPC)
		//userGroup := &domain.UserGroups{}
		//grpcUserGroup, err := impl.GrpcAuthenClientService.GetUserGroup(&pb.UserGroupRequest{RoleId: role.ID})
		//if err != nil {
		//	impl.Log.Error(err)
		//}
		//
		//if grpcUserGroup.Success {
		//	userGroup.KeySlug = grpcUserGroup.KeySlug
		//	userGroup.Name = grpcUserGroup.Name
		//	userGroup.Status = grpcUserGroup.Status
		//} else {
		//	userGroup = nil
		//	impl.Log.Error(grpcUserGroup.Error)
		//}
		//
		// get user data (gRPC)
		user := &domain.Users{}
		//grpcUser, err := impl.GrpcAuthenClientService.GetUserProfile(&pb.UserProfileRequest{UserId: role.CreatedBy})
		//if err != nil {
		//	impl.Log.Error(err)
		//}
		//
		//if grpcUser.Success {
		//	user.ID = grpcUser.UserId
		//	user.FirstName = grpcUser.FirstName
		//	user.LastName = grpcUser.LastName
		//} else {
		//	user = nil
		//	impl.Log.Error(grpcUser.Error)
		//}

		items[i] = inout.RoleToListView(role, user, impl.DateTime)
	}

	return total, items, nil
}
