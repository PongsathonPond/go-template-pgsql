package test

import (
	"errors"
	"net/http"

	"idev-cms-service/app/view"
	"idev-cms-service/service/util"
)

func (suite *PackageTestSuite) TestMakeErrResp() {
	err := util.ConvertInputToDomainErr(errors.New("test"))
	view.MakeErrResp(suite.ctx, err)
	suite.Equal(http.StatusBadRequest, suite.ctx.Writer.Status())
}
