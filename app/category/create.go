package category

import (
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"

	"idev-cms-service/app/view"
	"idev-cms-service/service/category/inout"
)

// Create godoc
//
//	@Tags			Category
//
// @Security     BearerAuth
//
//	@Summary		Create a new category
//	@Description	A newly created category, ID will be given in a Category-Location response header
//	@Accept			json
//	@Produce		json
//	@Param			Accept-Language	header		string					false	"Language"	default(en)
//	@Param			Request			body		inout.CategoryCreateInput	true	"Input"
//	@Success		201				{object}	view.SuccessCreateResp{data=string}
//	@Failure		400				{object}	view.Error400Resp{errors=[]view.ErrItem}
//	@Failure		404				{object}	view.Error404Resp{errors=[]view.ErrItem}
//	@Failure		422				{object}	view.Error422Resp{errors=[]view.ErrItem}
//	@Failure		500				{object}	view.Error500Resp{errors=[]view.ErrItem}
//	@Router			/category [post]
func (ctrl *Controller) Create(c *gin.Context) {
	sp, ctx := opentracing.StartSpanFromContext(c, "handler.category.Create")
	defer sp.Finish()

	inp := &inout.CategoryCreateInput{}
	if err := c.ShouldBindJSON(inp); err != nil {
		view.MakeErrResp(c, err)
		return
	}

	id, err := ctrl.service.Create(ctx, inp)
	if err != nil {
		view.MakeErrResp(c, err)
		return
	}

	view.MakeCreateSuccessResp(c, id)
}
