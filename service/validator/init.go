package validator

import (
	ioUser "idev-cms-service/service/users/inout"
	"idev-cms-service/service/util"
	vUser "idev-cms-service/service/validator/users"

	"github.com/go-playground/validator/v10"
)

type GoPlayGroundValidator struct {
	validate *validator.Validate
	Repo     *ValidatorRepository
}

type ValidatorRepository struct {
	User util.Repository
}

func New(vRepo *ValidatorRepository) (v *GoPlayGroundValidator) {
	v = &GoPlayGroundValidator{
		validate: validator.New(),
		Repo:     vRepo,
	}

	_ = v.validate.RegisterValidation("not-empty", notEmptyStringFunc)
	_ = v.validate.RegisterValidation("name-th", nameThFunc)
	_ = v.validate.RegisterValidation("name-en", nameEnFunc)

	// custom validate user
	cUser := vUser.New(vRepo.User)
	v.validate.RegisterStructValidation(cUser.CreateStructLevelValidation, &ioUser.UserCreateInput{})
	v.validate.RegisterStructValidation(cUser.UpdateStructLevelValidation, &ioUser.UserUpdateInput{})
	v.validate.RegisterStructValidation(cUser.UpdateMeStructLevelValidation, &ioUser.ProfileUpdateInput{})
	v.validate.RegisterStructValidation(cUser.ChangePasswordValidation, &ioUser.UserChangePasswordInput{})
	return v
}

func (v *GoPlayGroundValidator) Validate(item interface{}) (err error) {
	return v.validate.Struct(item)
}
