package users

import (
	"net/http"
	"reflect"

	"idev-cms-service/app/view"
	"idev-cms-service/service/users/inout"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
)

// ResendActivate godoc
// @Tags         Activation
// @Security     BearerAuth
// @Summary      Resend activation link
// @Description  Resend activation link to Email
// @Accept       json
// @Produce      json
// @Param        Accept-Language  header    string                     false  "Language"  default(en)
// @Param        Request          body      inout.ResendActivateInput  true   "Input"
// @Success      200              {object}  view.SuccessDefaultResp{data=object}
// @Failure      400              {object}  view.Error400Resp{errors=[]view.ErrItem}
// @Failure      404              {object}  view.Error404Resp{errors=[]view.ErrItem}
// @Failure      422              {object}  view.Error422Resp{errors=[]view.ErrItem}
// @Failure      500              {object}  view.Error500Resp{errors=[]view.ErrItem}
// @Router       /activate/resend [post]
func (ctrl *Controller) ResendActivate(c *gin.Context) {
	sp, ctx := opentracing.StartSpanFromContext(c, "handler.users.ResendActivate")
	defer sp.Finish()

	inp := &inout.ResendActivateInput{}
	if err := c.ShouldBindJSON(inp); err != nil {
		view.MakeErrResp(c, err)
		return
	}

	err := ctrl.service.ResendActivate(ctx, inp)
	if err != nil {
		view.MakeErrResp(c, err)
		return
	}

	view.MakeSuccessResp(c, http.StatusOK, "Success", reflect.ValueOf(nil))
}
