package test

import (
	"context"

	"idev-cms-service/service/util/mocks"
	"idev-cms-service/service/validator"

	"github.com/stretchr/testify/suite"
)

type PackageTestSuite struct {
	suite.Suite
	ctx           context.Context
	validator     validator.Validator
	repoUserGroup *mocks.Repository
	repoUser      *mocks.Repository
}

func (suite *PackageTestSuite) SetupTest() {
	suite.ctx = context.Background()
	suite.repoUserGroup = &mocks.Repository{}
	suite.repoUser = &mocks.Repository{}
	//vRepo := validator.ValidatorRepository{
	//	User: suite.repoUser,
	//}
	//vService := validator.ValidatorService{}
	//suite.validator = validator.New(&vRepo, &vService)
}
