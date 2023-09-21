package users

import (
	"fmt"

	"idev-cms-service/app/view"
	"idev-cms-service/service/users/inout"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
)

// UserChangePassword godoc
// @Tags         User Profile
// @Security     BearerAuth
// @Summary      Update password of a user
// @Description  Update password of a user by token
// @Accept       json
// @Produce      json
// @Param        Accept-Language  header    string                         false  "Language"  default(en)
// @Param        Request          body      inout.UserChangePasswordInput  true   "Input"
// @Success      200              {object}  view.SuccessUpdateResp{data=object}
// @Failure      400              {object}  view.Error400Resp{errors=[]view.ErrItem}
// @Failure      404              {object}  view.Error404Resp{errors=[]view.ErrItem}
// @Failure      422              {object}  view.Error422Resp{errors=[]view.ErrItem}
// @Failure      500              {object}  view.Error500Resp{errors=[]view.ErrItem}
// @Router       /me/password [put]
func (ctrl *Controller) UserChangePassword(c *gin.Context) {
	sp, ctx := opentracing.StartSpanFromContext(c, "handler.users.UserChangePassword")
	defer sp.Finish()

	userID, _ := c.Get("UserID")

	inp := &inout.UserChangePasswordInput{
		UserID: fmt.Sprintf("%v", userID),
	}

	if err := c.ShouldBindJSON(inp); err != nil {
		view.MakeErrResp(c, err)
		return
	}

	if err := ctrl.service.UserChangePassword(ctx, inp); err != nil {
		view.MakeErrResp(c, err)
		return
	}

	view.MakeUpdateSuccessResp(c)
}
