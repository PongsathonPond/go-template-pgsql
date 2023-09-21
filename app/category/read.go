package category

import (
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"idev-cms-service/app/view"
	"idev-cms-service/service/category/inout"

	_ "idev-cms-service/service/category/inout"
)

// Read godoc
//
//	@Tags			Category
//
// @Security     BearerAuth
//
//	@Summary		Get a category by category ID
//	@Description	Response a category data with a given category ID
//	@Accept			json
//	@Produce		json
//	@Param			Accept-Language	header		string	false	"Language"	default(en)
//	@Param			id				path		string	true	"Category ID"	default(123456789012345678)
//	@Success		200				{object}	view.SuccessReadResp{data=inout.CategoryView}
//	@Failure		400				{object}	view.Error400Resp{errors=[]view.ErrItem}
//	@Failure		404				{object}	view.Error404Resp{errors=[]view.ErrItem}
//	@Failure		422				{object}	view.Error422Resp{errors=[]view.ErrItem}
//	@Failure		500				{object}	view.Error500Resp{errors=[]view.ErrItem}
//	@Router			/category/{id} [get]
func (ctrl *Controller) Read(c *gin.Context) {
	sp, ctx := opentracing.StartSpanFromContext(c, "handler.category.Read")
	defer sp.Finish()

	inp := &inout.CategoryReadInput{
		ID: c.Param("id"),
	}

	category, err := ctrl.service.Read(ctx, inp)
	if err != nil {
		view.MakeErrResp(c, err)
		return
	}

	view.MakeReadSuccessResp(c, category)
}
