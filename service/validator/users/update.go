package users

import (
	"context"

	"idev-cms-service/service/users/inout"

	"github.com/go-playground/validator/v10"
)

func (c *customValidateUser) UpdateStructLevelValidation(structLV validator.StructLevel) {
	ctx := context.Background()
	input := structLV.Current().Interface().(inout.UserUpdateInput)

	c.checkUpdateEmailUnique(ctx, structLV, input.ID, input.Email)
	//c.checkUserGroupsExist(ctx, structLV, input.UserGroup)
	//c.checkRolesExist(structLV, input.Role)
}
