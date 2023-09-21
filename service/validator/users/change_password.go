package users

import (
	"regexp"

	"idev-cms-service/service/users/inout"

	"github.com/go-playground/validator/v10"
)

func (c *customValidateUser) ChangePasswordValidation(structLV validator.StructLevel) {
	input := structLV.Current().Interface().(inout.UserChangePasswordInput)

	re := regexp.MustCompile(`[^\s-]`)

	if input.NewPassword != "" && re.MatchString(input.NewPassword) {
		c.checkPasswordLength(structLV, input.NewPassword)
	}

	if input.OldPassword != "" && re.MatchString(input.OldPassword) {
		c.checkPasswordLength(structLV, input.OldPassword)
	}
}
