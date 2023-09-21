package users

import (
	"idev-cms-service/service/util"
)

type customValidateUser struct {
	repo util.Repository
}

func New(repo util.Repository) (customValidate *customValidateUser) {
	return &customValidateUser{
		repo: repo,
	}
}
