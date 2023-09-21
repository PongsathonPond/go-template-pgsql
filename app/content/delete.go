package content

import (
	"idev-cms-service/app/view"
	"idev-cms-service/service/content/inout"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
)

// Delete godoc
//
//	@Tags			Content
//
// @Security     BearerAuth
//
//	@Summary		Delete contents of a content
//	@Description	Delete content with a given content ID
//	@Accept			json
//	@Produce		json
//	@Param			Accept-Language	header		string	false	"Language"	default(en)
//	@Param			id				path		string	true	"Content ID"	default(123456789012345678)
//	@Success		200				{object}	view.SuccessDeleteResp{data=object}
//	@Failure		400				{object}	view.Error400Resp{errors=[]view.ErrItem}
//	@Failure		404				{object}	view.Error404Resp{errors=[]view.ErrItem}
//	@Failure		422				{object}	view.Error422Resp{errors=[]view.ErrItem}
//	@Failure		500				{object}	view.Error500Resp{errors=[]view.ErrItem}
//	@Router			/content/{id} [delete]
func (ctrl *Controller) Delete(c *gin.Context) {
	sp, ctx := opentracing.StartSpanFromContext(c, "handler.content.Delete")
	defer sp.Finish()

	inp := &inout.ContentDeleteInput{
		ID: c.Param("id"),
	}

	if err := ctrl.service.Delete(ctx, inp); err != nil {
		view.MakeErrResp(c, err)
		return
	}

	view.MakeDeleteSuccessResp(c)
}
