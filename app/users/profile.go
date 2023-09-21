package users

import (
	"strings"

	"idev-cms-service/app/view"
	"idev-cms-service/service/users/inout"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
)

// Update godoc
// @Tags         User Profile
// @Security     BearerAuth
// @Summary      Update contents of me
// @Description  Update me with a given data
// @Accept       json
// @Produce      json
// @Param        Accept-Language  header    string                    false  "Language"  default(en)
// @Param        Request          body      inout.ProfileUpdateInput  true   "Input"
// @Success      200              {object}  view.SuccessUpdateResp{data=object}
// @Failure      400              {object}  view.Error400Resp{errors=[]view.ErrItem}
// @Failure      404              {object}  view.Error404Resp{errors=[]view.ErrItem}
// @Failure      422              {object}  view.Error422Resp{errors=[]view.ErrItem}
// @Failure      500              {object}  view.Error500Resp{errors=[]view.ErrItem}
// @Router       /me [put]
func (ctrl *Controller) UpdateMe(c *gin.Context) {
	sp, ctx := opentracing.StartSpanFromContext(c, "handler.users.Profile")
	defer sp.Finish()

	inp := &inout.ProfileUpdateInput{}
	if err := c.ShouldBindJSON(inp); err != nil {
		view.MakeErrResp(c, err)
		return
	}

	inp.AccessToken = strings.ReplaceAll(c.GetHeader("Authorization"), "Bearer ", "")

	if err := ctrl.service.UpdateMe(ctx, inp); err != nil {
		view.MakeErrResp(c, err)
		return
	}
	view.MakeUpdateSuccessResp(c)
}
