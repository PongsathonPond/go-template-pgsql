package roles

import (
	"idev-cms-service/app/view"
	"idev-cms-service/service/roles/inout"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
)

// All godoc
// @Tags         Roles
// @Summary      Get all data of roles
// @Description  Return a list of all roles
// @Accept       json
// @Produce      json
// @Param        Accept-Language  header    string  false  "Language"  default(en)
// @Success      200              {object}  view.SuccessReadResp{data=[]inout.RolesAllView}
// @Failure      400              {object}  view.Error400Resp{errors=[]view.ErrItem}
// @Failure      404              {object}  view.Error404Resp{errors=[]view.ErrItem}
// @Failure      422              {object}  view.Error422Resp{errors=[]view.ErrItem}
// @Failure      500              {object}  view.Error500Resp{errors=[]view.ErrItem}
// @Router       /all/roles [get]
func (ctrl *Controller) All(c *gin.Context) {
	sp, ctx := opentracing.StartSpanFromContext(c, "handler.roles.All")
	defer sp.Finish()

	input := &inout.RoleAllInput{}
	if err := c.ShouldBind(input); err != nil {
		view.MakeErrResp(c, err)
		return
	}

	data, err := ctrl.service.All(ctx, input)
	if err != nil {
		view.MakeErrResp(c, err)
		return
	}

	view.MakeReadSuccessResp(c, data)
}
