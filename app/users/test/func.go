package test

import "idev-cms-service/service/users/inout"

func (suite *PackageTestSuite) makeTestReadInput() (input *inout.ReadInput) {
	return &inout.ReadInput{
		ID: "test",
	}
}

func (suite *PackageTestSuite) makeTestCreateInput() (input *inout.UserCreateInput) {
	return &inout.UserCreateInput{
		ID:           "",
		FirstName:    "test",
		LastName:     "test",
		Email:        "test",
		Mobile:       "test",
		Role:         "test",
		Status:       "test",
		ImageProfile: "test",
	}
}

func (suite *PackageTestSuite) makeTestUpdateInput() (input *inout.UserUpdateInput) {
	return &inout.UserUpdateInput{
		ID:           "test",
		FirstName:    "test",
		LastName:     "test",
		Mobile:       "test",
		Role:         "test",
		Status:       "test",
		ImageProfile: "test",
	}
}

func (suite *PackageTestSuite) makeTestDeleteInput() (input *inout.UserDeleteInput) {
	return &inout.UserDeleteInput{
		ID: "test",
	}
}

func (suite *PackageTestSuite) makeTestGetMeInput() (input *inout.GetMeInput) {
	return &inout.GetMeInput{
		AccessToken: "",
	}
}

func (suite *PackageTestSuite) makeTestUpdateMeInput() (input *inout.ProfileUpdateInput) {
	return &inout.ProfileUpdateInput{
		ID:           "",
		Firstname:    "test",
		Lastname:     "test",
		Mobile:       "test",
		ImageProfile: "test",
		AccessToken:  "",
	}
}

func (suite *PackageTestSuite) makeTestResendActivateInput() (input *inout.ResendActivateInput) {
	return &inout.ResendActivateInput{UserID: givenTest}
}

func (suite *PackageTestSuite) makeTestActivateInput() (input *inout.UserActivateInput) {
	return &inout.UserActivateInput{
		ActivateToken: givenTest,
		Password:      givenTest,
	}
}
