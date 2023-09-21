package tokens

import (
	"net/http"

	"idev-cms-service/app/view"
	"idev-cms-service/service/tokens/inout"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
)

// Refresh godoc
// @Tags         Auth
// @Summary      Renew your access token by refresh token
// @Description  Use this operation to obtain new authorization token which could be used in all other operations
// @Accept       json
// @Produce      json
// @Param        Accept-Language  header    string                   false  "Language"  default(en)
// @Param        Request          body      inout.RefreshTokenInput  true   "Input"
// @Success      200              {object}  view.SuccessReadResp{data=inout.RefreshTokenView}
// @Failure      400              {object}  view.Error400Resp{errors=[]view.ErrItem}
// @Failure      404              {object}  view.Error404Resp{errors=[]view.ErrItem}
// @Failure      422              {object}  view.Error422Resp{errors=[]view.ErrItem}
// @Failure      500              {object}  view.Error500Resp{errors=[]view.ErrItem}
// @Router       /token/refresh [post]
func (ctrl *Controller) Refresh(c *gin.Context) {
	sp, ctx := opentracing.StartSpanFromContext(c, "handler.tokens.Refresh")
	defer sp.Finish()

	input := &inout.RefreshTokenInput{}
	if err := c.ShouldBindJSON(input); err != nil {
		view.MakeErrResp(c, err)
		return
	}

	data, err := ctrl.service.RefreshToken(ctx, input)
	if err != nil {
		view.MakeErrResp(c, err)
		return
	}

	view.MakeSuccessResp(c, http.StatusOK, view.MsgGetDataSuccess(c), data)
}
