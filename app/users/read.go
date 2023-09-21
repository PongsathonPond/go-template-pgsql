package users

import (
	"idev-cms-service/app/view"
	"idev-cms-service/service/users/inout"

	_ "idev-cms-service/service/users/inout"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
)

// Read godoc
// @Tags         Users
// @Security     BearerAuth
// @Summary      Get a user by user ID
// @Description  Response a user data with a given user ID
// @Accept       json
// @Produce      json
// @Param        Accept-Language  header    string  false  "Language"  default(en)
// @Param        id               path      string  true   "User ID"   default(123456789012345678)
// @Success      200              {object}  view.SuccessReadResp{data=inout.UserReadView}
// @Failure      400              {object}  view.Error400Resp{errors=[]view.ErrItem}
// @Failure      404              {object}  view.Error404Resp{errors=[]view.ErrItem}
// @Failure      422              {object}  view.Error422Resp{errors=[]view.ErrItem}
// @Failure      500              {object}  view.Error500Resp{errors=[]view.ErrItem}
// @Router       /users/{id} [get]
func (ctrl *Controller) Read(c *gin.Context) {
	sp, ctx := opentracing.StartSpanFromContext(c, "handler.users.Read")
	defer sp.Finish()

	input := &inout.ReadInput{
		ID: c.Param("id"),
	}

	user, err := ctrl.service.Read(ctx, input)
	if err != nil {
		view.MakeErrResp(c, err)
		return
	}

	view.MakeReadSuccessResp(c, user)
}
