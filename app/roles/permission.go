package roles

import (
	"strings"

	"idev-cms-service/app/view"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
)

// GetUserPermission godoc
// @Tags         User Permission
// @Security     BearerAuth
// @Summary      Get user permission
// @Description  Response user permission data
// @Accept       json
// @Produce      json
// @Param        Accept-Language  header    string  false  "Language"  default(en)
// @Success      200              {object}  view.SuccessReadResp{data=[]inout.MenuPermissionView}
// @Failure      400              {object}  view.Error400Resp{errors=[]view.ErrItem}
// @Failure      404              {object}  view.Error404Resp{errors=[]view.ErrItem}
// @Failure      422              {object}  view.Error422Resp{errors=[]view.ErrItem}
// @Failure      500              {object}  view.Error500Resp{errors=[]view.ErrItem}
// @Router       /me/permissions [get]
func (ctrl *Controller) GetUserPermission(c *gin.Context) {
	sp, ctx := opentracing.StartSpanFromContext(c, "handler.menus.GetUserPermission")
	defer sp.Finish()

	items, err := ctrl.service.GetUserPermission(ctx, strings.ReplaceAll(c.GetHeader("Authorization"), "Bearer ", ""))
	if err != nil {
		view.MakeErrResp(c, err)
		return
	}

	view.MakeReadSuccessResp(c, items)
}
