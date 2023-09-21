package roles

import (
	"strings"

	"idev-cms-service/app/view"
	"idev-cms-service/service/roles/inout"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
)

// Create godoc
// @Tags         Roles
// @Security     BearerAuth
// @Summary      Create a new role
// @Description  A newly created role, ID will be given in a Content-Location response header
// @Accept       json
// @Produce      json
// @Param        Accept-Language  header    string                 false  "Language"  default(en)
// @Param        Request          body      inout.RoleCreateInput  true   "Input"
// @Success      201              {object}  view.SuccessCreateResp{data=object}
// @Failure      400              {object}  view.Error400Resp{errors=[]view.ErrItem}
// @Failure      404              {object}  view.Error404Resp{errors=[]view.ErrItem}
// @Failure      422              {object}  view.Error422Resp{errors=[]view.ErrItem}
// @Failure      500              {object}  view.Error500Resp{errors=[]view.ErrItem}
// @Router       /roles [post]
func (ctrl *Controller) Create(c *gin.Context) {
	sp, ctx := opentracing.StartSpanFromContext(c, "handler.roles.Create")
	defer sp.Finish()

	inp := &inout.RoleCreateInput{}
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
