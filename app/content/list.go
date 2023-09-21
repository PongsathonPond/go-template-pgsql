package content

import (
	"idev-cms-service/app/view"
	"idev-cms-service/domain"

	_ "idev-cms-service/service/content/inout"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
)

// List godoc
//
//	@Tags			Content
//
// @Security     BearerAuth
//
//	@Summary		Get a list of content
//	@Description	Return a list of content filtered by a given filters if any
//	@Accept			json
//	@Produce		json
//	@Param			Accept-Language	header		string		false	"Language"							default(en)
//	@Param			page			query		int			false	"A page number"						default(1)
//	@Param			per_page		query		int			false	"A total number of items per page"	default(15)
//	@Param			filters			query		[]string	false	"Condition for content retrieval, ex. `status:eq:active` | `age:gte:25`"
//	@Param			search			query		[]string	false	"Search with like condition for content retrieval, ex. `name:john`"
//	@Param			sorts			query		[]string	false	"Sort for content data, ex. `created_at:asc`"
//	@Success		200				{object}	view.SuccessPaginatorResp{data=view.dataListWithOption{lists=[]inout.ContentView}}
//	@Failure		400				{object}	view.Error400Resp{errors=[]view.ErrItem}
//	@Failure		404				{object}	view.Error404Resp{errors=[]view.ErrItem}
//	@Failure		422				{object}	view.Error422Resp{errors=[]view.ErrItem}
//	@Failure		500				{object}	view.Error500Resp{errors=[]view.ErrItem}
//	@Router			/content [get]
func (ctrl *Controller) List(c *gin.Context) {
	sp, ctx := opentracing.StartSpanFromContext(c, "handler.content.List")
	defer sp.Finish()

	opt := &domain.PageOption{}
	if err := c.ShouldBind(opt); err != nil {
		view.MakeErrResp(c, err)
		return
	}

	opt.Filters = append(opt.Filters, "deleted_at:isNull")
	if len(opt.Sorts) < 1 {
		opt.Sorts = append(opt.Sorts, "created_at:desc")
	}

	total, items, err := ctrl.service.List(ctx, opt)
	if err != nil {
		view.MakeErrResp(c, err)
		return
	}

	view.MakePaginatorResp(c, total, opt, items)
}
