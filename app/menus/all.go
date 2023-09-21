package menus

import (
	"idev-cms-service/app/view"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
)

// All godoc
// @Tags         Menus
// @Summary      Get a list of Menus
// @Description  Return a list of Menus filtered by a given filters if any
// @Accept       json
// @Produce      json
// @Param        Accept-Language  header    string  false  "Language"  default(en)
// @Param        user_group       query     string  false  "A User Group Keyslug"
// @Success      200              {object}  view.SuccessGetAllResp{lists=[]inout.MenuView}
// @Failure      400              {object}  view.Error400Resp{errors=[]view.ErrItem}
// @Failure      404              {object}  view.Error404Resp{errors=[]view.ErrItem}
// @Failure      422              {object}  view.Error422Resp{errors=[]view.ErrItem}
// @Failure      500              {object}  view.Error500Resp{errors=[]view.ErrItem}
// @Router       /all/menus [get]
func (ctrl *Controller) All(c *gin.Context) {
	sp, ctx := opentracing.StartSpanFromContext(c, "handler.menus.All")
	defer sp.Finish()

	total, items, err := ctrl.service.All(ctx)
	if err != nil {
		view.MakeErrResp(c, err)
		return
	}

	view.MakeGetAllResp(c, total, items)
}
