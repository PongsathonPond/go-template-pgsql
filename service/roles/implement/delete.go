package implement

import (
	"context"

	"idev-cms-service/domain"
	"idev-cms-service/service/roles/inout"
	"idev-cms-service/service/util"
)

func (impl *implementation) Delete(ctx context.Context, input *inout.RoleDeleteInput) (err error) {
	if err = impl.Validator.Validate(input); err != nil {
		return util.ValidationUpdateErr(err)
	}

	filters := []string{
		impl.FilterString.MakeID(input.ID),
		impl.FilterString.MakeDeletedAtIsNull(),
	}

	if err = impl.Repo.Read(ctx, filters, &domain.Roles{}); err != nil {
		return util.RepoReadErr(err)
	}

	if err = impl.Repo.SoftDelete(ctx, filters); err != nil {
		return util.RepoDeleteErr(err)
	}

	// get user group data (gRPC)
	//userGroup := &domain.UserGroups{}
	//respGetUserGroup, err := impl.GrpcAuthenClientService.GetUserGroup(&pb.UserGroupRequest{RoleId: input.ID})
	//if err != nil {
	//	impl.Log.Error(err)
	//}
	//
	//if respGetUserGroup.Success {
	//	userGroup.KeySlug = respGetUserGroup.KeySlug
	//} else {
	//	impl.Log.Error(respGetUserGroup.Error)
	//}

	// set user group data (gRPC)
	//respSetUserGroup, err := impl.GrpcAuthenClientService.SetUserGroupRoles(&pb.UserGroupRolesRequest{
	//	RoleId:       input.ID,
	//	UserGroupKey: userGroup.KeySlug,
	//	Mode:         "delete",
	//})
	//if err != nil {
	//	impl.Log.Error(err)
	//}
	//
	//if !respSetUserGroup.Success {
	//	impl.Log.Error(respSetUserGroup.Error)
	//}

	return nil
}
