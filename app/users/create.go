package users

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"

	"idev-cms-service/app/view"
	"idev-cms-service/service/users/inout"
)

// Create godoc
// @Tags         Users
// @Security     BearerAuth
// @Summary      Create a new user
// @Description  A newly created user
// @Accept       json
// @Produce      json
// @Param        Accept-Language  header    string                 false  "Language"  default(en)
// @Param        Request          body      inout.UserCreateInput  true   "Input"
// @Success      201              {object}  view.SuccessCreateResp{}
// @Failure      400              {object}  view.Error400Resp{errors=[]view.ErrItem}
// @Failure      404              {object}  view.Error404Resp{errors=[]view.ErrItem}
// @Failure      422              {object}  view.Error422Resp{errors=[]view.ErrItem}
// @Failure      500              {object}  view.Error500Resp{errors=[]view.ErrItem}
// @Router       /users [post]
func (ctrl *Controller) Create(c *gin.Context) {
	sp, ctx := opentracing.StartSpanFromContext(c, "handler.users.Create")
	defer sp.Finish()

	inp := &inout.UserCreateInput{}
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
