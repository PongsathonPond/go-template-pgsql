package roles

import (
	"idev-cms-service/app/view"
	"idev-cms-service/service/roles/inout"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
)

// Read godoc
// @Tags         Roles
// @Summary      Get a role by role ID
// @Description  Response a role data with a given staff ID
// @Accept       json
// @Produce      json
// @Param        Accept-Language  header    string  false  "Language"  default(en)
// @Param        id               path      string  true   "Role ID"   default(123456789012345678)
// @Success      200              {object}  view.SuccessReadResp{data=inout.RoleReadView}
// @Failure      400              {object}  view.Error400Resp{errors=[]view.ErrItem}
// @Failure      404              {object}  view.Error404Resp{errors=[]view.ErrItem}
// @Failure      422              {object}  view.Error422Resp{errors=[]view.ErrItem}
// @Failure      500              {object}  view.Error500Resp{errors=[]view.ErrItem}
// @Router       /roles/{id} [get]
func (ctrl *Controller) Read(c *gin.Context) {
	sp, ctx := opentracing.StartSpanFromContext(c, "handler.roles.Read")
	defer sp.Finish()

	inp := &inout.RoleReadInput{
		ID: c.Param("id"),
	}

	role, err := ctrl.service.Read(ctx, inp)
	if err != nil {
		view.MakeErrResp(c, err)
		return
	}

	view.MakeReadSuccessResp(c, role)
}
