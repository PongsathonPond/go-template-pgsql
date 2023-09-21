package implement

import (
	"context"

	"idev-cms-service/domain"
	"idev-cms-service/service/users/inout"
	"idev-cms-service/service/util"
)

func (impl *implementation) Read(ctx context.Context, input *inout.ReadInput) (view *inout.UserReadView, err error) {
	user := &domain.Users{}
	filters := []string{
		impl.FilterString.MakeID(input.ID),
	}

	err = impl.RepoUsers.Read(ctx, filters, user)
	if err != nil {
		return nil, util.RepoReadErr(err)
	}

	// get user group data
	//userGroup := &domain.UserGroups{}
	filters = []string{
		//impl.FilterString.MakeKeySlug(user.UserGroup),
	}
	//if err = impl.RepoUserGroups.Read(ctx, filters, userGroup); err != nil {
	//	userGroup = nil
	//}

	// get user role data (gRPC)
	role := &domain.Roles{}
	//grpcRole, err := impl.GRPCClientService.GetRole(&pb.RoleRequest{RoleId: user.Role})
	//if err == nil {
	//	role.ID = grpcRole.RoleId
	//	role.Name = grpcRole.Name
	//	role.Status = grpcRole.Status
	//} else {
	//	role = nil
	//}

	// get user createdBy
	createdBy := &domain.Users{}
	filters = []string{
		impl.FilterString.MakeID(user.CreatedBy),
		impl.FilterString.MakeDeletedAtIsNull(),
	}

	_ = impl.RepoUsers.Read(ctx, filters, createdBy)

	return inout.UserToView(user, role, createdBy, impl.DateTime), nil
}
