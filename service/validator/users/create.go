package users

import (
	"context"

	"idev-cms-service/service/users/inout"

	"github.com/go-playground/validator/v10"
)

func (c *customValidateUser) CreateStructLevelValidation(structLV validator.StructLevel) {
	ctx := context.Background()
	input := structLV.Current().Interface().(inout.UserCreateInput)

	c.checkEmailUnique(ctx, structLV, input.Email)
	//c.checkUserGroupsExist(ctx, structLV, input.UserGroup)
	//c.checkRolesExist(structLV, input.Role)
}
