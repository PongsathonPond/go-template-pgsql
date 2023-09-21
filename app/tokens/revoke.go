package tokens

import (
	"net/http"
	"strings"

	"idev-cms-service/app/view"
	"idev-cms-service/service/tokens/inout"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
)

// Revoke godoc
// @Tags         Auth
// @Security     BearerAuth
// @Summary      Invalidate all of your tokens
// @Description  Invalidate all of your tokens
// @Accept       json
// @Produce      json
// @Param        Accept-Language  header    string  false  "Language"  default(en)
// @Success      200              {object}  view.SuccessUpdateResp{data=object}
// @Failure      400              {object}  view.Error400Resp{errors=[]view.ErrItem}
// @Failure      404              {object}  view.Error404Resp{errors=[]view.ErrItem}
// @Failure      422              {object}  view.Error422Resp{errors=[]view.ErrItem}
// @Failure      500              {object}  view.Error500Resp{errors=[]view.ErrItem}
// @Router       /token/revoke [post]
func (ctrl *Controller) Revoke(c *gin.Context) {
	sp, ctx := opentracing.StartSpanFromContext(c, "handler.tokens.Revoke")
	defer sp.Finish()

	// extract token
	token := strings.ReplaceAll(c.GetHeader("Authorization"), "Bearer ", "")

	input := &inout.RevokeTokenInput{
		AccessToken: token,
	}

	err := ctrl.service.RevokeToken(ctx, input)
	if err != nil {
		view.MakeErrResp(c, err)
		return
	}

	c.Set("UserID", nil)

	view.MakeSuccessResp(c, http.StatusOK, view.MsgUpdateDataSuccess(c), nil)
}
