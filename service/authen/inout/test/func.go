package test

import (
	"idev-cms-service/domain"
	"idev-cms-service/service/authen/inout"
)

func (suite *PackageTestSuite) makeTestTokenInput() (input *inout.TokenInput) {
	return &inout.TokenInput{
		ID:           "test",
		UserID:       "test",
		UserAgent:    "test",
		AccessToken:  "test",
		RefreshToken: "test",
	}
}

func (suite *PackageTestSuite) makeTestTokenUpdate() (input *inout.TokenUpdate) {
	return &inout.TokenUpdate{
		AccessToken:  "test",
		RefreshToken: "test",
	}
}

func (suite *PackageTestSuite) makeTestUsers() (users *domain.Users) {
	return &domain.Users{
		ID:           "test",
		FirstName:    "test",
		LastName:     "Test",
		Email:        "test",
		Mobile:       "test",
		Username:     "Test",
		Password:     "test",
		Role:         "test",
		ImageProfile: "test",
		Status:       "test",
		CreatedBy:    "test",
		CreatedAt:    0,
		UpdatedAt:    0,
		DeletedAt:    0,
	}
}

func (suite *PackageTestSuite) makeTestRoles() (roles *domain.Roles) {
	return &domain.Roles{
		ID:          "test",
		Name:        "test",
		Description: "test",
		Permissions: nil,
		Status:      "test",
		CreatedBy:   "test",
		CreatedAt:   0,
		UpdatedAt:   0,
		DeletedAt:   0,
	}
}
