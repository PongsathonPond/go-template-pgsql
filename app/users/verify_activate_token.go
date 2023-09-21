package users

import (
	"net/http"
	"reflect"

	"idev-cms-service/app/view"
	"idev-cms-service/service/users/inout"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
)

// Verify Activate Token godoc
// @Tags         Verify
// @Summary      Verify activate token form Invitation link
// @Description  Verify activate token before set password for first time login form Invitation link
// @Accept       json
// @Produce      json
// @Param        Accept-Language  header    string                          false  "Language"  default(en)
// @Param        Request          body      inout.VerifyActivateTokenInput  true   "Input"
// @Success      200              {object}  view.SuccessDefaultResp{data=object}
// @Failure      400              {object}  view.Error400Resp{errors=[]view.ErrItem}
// @Failure      404              {object}  view.Error404Resp{errors=[]view.ErrItem}
// @Failure      422              {object}  view.Error422Resp{errors=[]view.ErrItem}
// @Failure      500              {object}  view.Error500Resp{errors=[]view.ErrItem}
// @Router       /activate/verify [post]
func (ctrl *Controller) VerifyActivateToken(c *gin.Context) {
	sp, ctx := opentracing.StartSpanFromContext(c, "handler.users.VerifyActivateToken")
	defer sp.Finish()

	inp := &inout.VerifyActivateTokenInput{}
	if err := c.ShouldBindJSON(inp); err != nil {
		view.MakeErrResp(c, err)
		return
	}

	err := ctrl.service.VerifyActivateToken(ctx, inp)
	if err != nil {
		view.MakeErrResp(c, err)
		return
	}

	view.MakeSuccessResp(c, http.StatusOK, "Success", reflect.ValueOf(nil))
}
