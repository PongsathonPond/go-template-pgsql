package test

import (
	"errors"
	"net/http"

	"idev-cms-service/domain"
	"idev-cms-service/service/users/inout"

	"github.com/stretchr/testify/mock"
)

func (suite *PackageTestSuite) TestList() {
	req, resp, err := suite.makeListRequest()
	suite.NoError(err)

	opt := &domain.PageOption{
		Page:    1,
		PerPage: 15,
		Filters: []string{"deleted_at:isNull"},
		Sorts:   []string{"created_at:desc"},
	}

	suite.service.On("List", mock.Anything, opt).Once().Return(0, []*inout.UserReadView{}, nil)
	suite.router.ServeHTTP(resp, req)

	suite.Equal(http.StatusNoContent, resp.Code)
}

func (suite *PackageTestSuite) TestListInvalidQueryParam() {
	req, resp, err := suite.makeListRequestInvalidQueryParam()
	suite.NoError(err)

	suite.service.On("List", mock.Anything, nil).Once()
	suite.router.ServeHTTP(resp, req)

	suite.Equal(http.StatusInternalServerError, resp.Code)
}

func (suite *PackageTestSuite) TestListServiceError() {
	req, resp, err := suite.makeListRequest()
	suite.NoError(err)

	opt := &domain.PageOption{
		Page:    1,
		PerPage: 15,
		Filters: []string{"deleted_at:isNull"},
		Sorts:   []string{"created_at:desc"},
	}

	givenErr := errors.New("some error message")

	suite.service.On("List", mock.Anything, opt).Once().Return(0, nil, givenErr)
	suite.router.ServeHTTP(resp, req)

	suite.Equal(http.StatusInternalServerError, resp.Code)
}
