package authen

import (
	"net/http"
	"strings"

	"idev-cms-service/app/view"
	"idev-cms-service/service/authen/inout"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
)

// Logout godoc
// @Tags         Auth
// @Security     BearerAuth
// @Summary      Invalidate currently used token
// @Description  Invalidate currently used token
// @Accept       json
// @Produce      json
// @Param        Accept-Language  header    string  false  "Language"  default(en)
// @Success      200              {object}  view.SuccessLogoutResp{data=object}
// @Failure      400              {object}  view.Error400Resp{errors=[]view.ErrItem}
// @Failure      404              {object}  view.Error404Resp{errors=[]view.ErrItem}
// @Failure      422              {object}  view.Error422Resp{errors=[]view.ErrItem}
// @Failure      500              {object}  view.Error500Resp{errors=[]view.ErrItem}
// @Router       /logout [post]
func (ctrl *Controller) Logout(c *gin.Context) {
	sp, ctx := opentracing.StartSpanFromContext(c, "handler.authen.Logout")
	defer sp.Finish()

	// extract token
	token := strings.ReplaceAll(c.GetHeader("Authorization"), "Bearer ", "")

	input := &inout.LogoutInput{
		AccessToken: token,
	}

	err := ctrl.service.Logout(ctx, input)
	if err != nil {
		view.MakeErrResp(c, err)
		return
	}

	c.Set("UserID", nil)

	view.MakeSuccessResp(c, http.StatusOK, view.MsgLogoutSuccess(c), nil)
}
