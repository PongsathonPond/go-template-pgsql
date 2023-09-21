package content

import (
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"strings"

	"idev-cms-service/app/view"
	"idev-cms-service/service/content/inout"
)

// Create godoc
//
//	@Tags			Content
//
// @Security     BearerAuth
//
//	@Summary		Create a new content
//	@Description	A newly created content, ID will be given in a Content-Location response header
//	@Accept			json
//	@Produce		json
//	@Param			Accept-Language	header		string					false	"Language"	default(en)
//	@Param			Request			body		inout.ContentCreateInput	true	"Input"
//	@Success		201				{object}	view.SuccessCreateResp{data=string}
//	@Failure		400				{object}	view.Error400Resp{errors=[]view.ErrItem}
//	@Failure		404				{object}	view.Error404Resp{errors=[]view.ErrItem}
//	@Failure		422				{object}	view.Error422Resp{errors=[]view.ErrItem}
//	@Failure		500				{object}	view.Error500Resp{errors=[]view.ErrItem}
//	@Router			/content [post]
func (ctrl *Controller) Create(c *gin.Context) {
	sp, ctx := opentracing.StartSpanFromContext(c, "handler.content.Create")
	defer sp.Finish()

	inp := &inout.ContentCreateInput{}
	if err := c.ShouldBindJSON(inp); err != nil {
		view.MakeErrResp(c, err)
		return
	}

	inp.AccessToken = strings.ReplaceAll(c.GetHeader("Authorization"), "Bearer ", "")
	id, err := ctrl.service.Create(ctx, inp)
	if err != nil {
		view.MakeErrResp(c, err)
		return
	}

	view.MakeCreateSuccessResp(c, id)
}
