package users

import (
	"idev-cms-service/app/view"
	"idev-cms-service/service/users/inout"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
)

// Update godoc
// @Tags         Users
// @Security     BearerAuth
// @Summary      Update contents of a user
// @Description  Update user with a given user ID according to a given data
// @Accept       json
// @Produce      json
// @Param        Accept-Language  header    string                 false  "Language"  default(en)
// @Param        id               path      string                 true   "User ID"   default(123456789012345678)
// @Param        Request          body      inout.UserUpdateInput  true   "Input"
// @Success      200              {object}  view.SuccessUpdateResp{data=object}
// @Failure      400              {object}  view.Error400Resp{errors=[]view.ErrItem}
// @Failure      404              {object}  view.Error404Resp{errors=[]view.ErrItem}
// @Failure      422              {object}  view.Error422Resp{errors=[]view.ErrItem}
// @Failure      500              {object}  view.Error500Resp{errors=[]view.ErrItem}
// @Router       /users/{id} [put]
func (ctrl *Controller) Update(c *gin.Context) {
	sp, ctx := opentracing.StartSpanFromContext(c, "handler.users.Update")
	defer sp.Finish()

	inp := &inout.UserUpdateInput{
		ID: c.Param("id"),
	}
	if err := c.ShouldBindJSON(inp); err != nil {
		view.MakeErrResp(c, err)
		return
	}

	if err := ctrl.service.Update(ctx, inp); err != nil {
		view.MakeErrResp(c, err)
		return
	}

	view.MakeUpdateSuccessResp(c)
}
