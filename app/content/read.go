package content

import (
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"idev-cms-service/app/view"
	"idev-cms-service/service/content/inout"

	_ "idev-cms-service/service/content/inout"
)

// Read godoc
//
//	@Tags			Content
//
// @Security     BearerAuth
//
//	@Summary		Get a content by content ID
//	@Description	Response a content data with a given content ID
//	@Accept			json
//	@Produce		json
//	@Param			Accept-Language	header		string	false	"Language"	default(en)
//	@Param			id				path		string	true	"Content ID"	default(123456789012345678)
//	@Success		200				{object}	view.SuccessReadResp{data=inout.ContentView}
//	@Failure		400				{object}	view.Error400Resp{errors=[]view.ErrItem}
//	@Failure		404				{object}	view.Error404Resp{errors=[]view.ErrItem}
//	@Failure		422				{object}	view.Error422Resp{errors=[]view.ErrItem}
//	@Failure		500				{object}	view.Error500Resp{errors=[]view.ErrItem}
//	@Router			/content/{id} [get]
func (ctrl *Controller) Read(c *gin.Context) {
	sp, ctx := opentracing.StartSpanFromContext(c, "handler.content.Read")
	defer sp.Finish()

	inp := &inout.ContentReadInput{
		ID: c.Param("id"),
	}

	content, err := ctrl.service.Read(ctx, inp)
	if err != nil {
		view.MakeErrResp(c, err)
		return
	}

	view.MakeReadSuccessResp(c, content)
}
