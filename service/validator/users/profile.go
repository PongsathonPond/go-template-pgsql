package users

import (
	"idev-cms-service/service/users/inout"

	"github.com/go-playground/validator/v10"
)

func (c *customValidateUser) UpdateMeStructLevelValidation(structLV validator.StructLevel) {
	input := structLV.Current().Interface().(inout.ProfileUpdateInput)
	c.checkFormatMobileNumber(structLV, input.Mobile)
}
