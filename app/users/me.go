package users

import (
	"net/http"
	"strings"

	"idev-cms-service/app/view"
	"idev-cms-service/service/users/inout"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
)

// GetMe godoc
// @Tags         User Profile
// @Security     BearerAuth
// @Summary      Get me data
// @Description  Response me data
// @Accept       json
// @Produce      json
// @Param        Accept-Language  header    string  false  "Language"  default(en)
// @Success      200              {object}  view.SuccessReadResp{data=inout.MeView}
// @Failure      400              {object}  view.Error400Resp{errors=[]view.ErrItem}
// @Failure      404              {object}  view.Error404Resp{errors=[]view.ErrItem}
// @Failure      422              {object}  view.Error422Resp{errors=[]view.ErrItem}
// @Failure      500              {object}  view.Error500Resp{errors=[]view.ErrItem}
// @Router       /me [get]
func (ctrl *Controller) GetMe(c *gin.Context) {
	sp, ctx := opentracing.StartSpanFromContext(c, "handler.users.GetMe")
	defer sp.Finish()

	input := &inout.GetMeInput{
		AccessToken: strings.ReplaceAll(c.GetHeader("Authorization"), "Bearer ", ""),
	}

	user, err := ctrl.service.GetMe(ctx, input)
	if err != nil {
		view.MakeErrResp(c, err)
		return
	}

	view.MakeSuccessResp(c, http.StatusOK, view.MsgGetDataSuccess(c), user)
}
