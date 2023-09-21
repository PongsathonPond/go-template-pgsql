package authen

import (
	"net/http"

	"idev-cms-service/app/view"
	"idev-cms-service/service/authen/inout"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
)

// Login godoc
// @Tags         Auth
// @Summary      Create auth token from username and password
// @Description  Use this operation to obtain an authorization token which could be used in all other operations
// @Accept       json
// @Produce      json
// @Param        Accept-Language  header    string            false  "Language"  default(en)
// @Param        Request          body      inout.LoginInput  true   "Input"
// @Success      200              {object}  view.SuccessLoginResp{data=inout.UserViewWithToken}
// @Failure      400              {object}  view.Error400Resp{errors=[]view.ErrItem}
// @Failure      404              {object}  view.Error404Resp{errors=[]view.ErrItem}
// @Failure      422              {object}  view.Error422Resp{errors=[]view.ErrItem}
// @Failure      500              {object}  view.Error500Resp{errors=[]view.ErrItem}
// @Router       /login [post]
func (ctrl *Controller) Login(c *gin.Context) {
	sp, ctx := opentracing.StartSpanFromContext(c, "handler.authen.Login")
	defer sp.Finish()

	input := &inout.LoginInput{}
	if err := c.ShouldBindJSON(input); err != nil {
		view.MakeErrResp(c, err)
		return
	}
	input.UserAgent = c.Request.Header.Get("User-Agent")

	user, err := ctrl.service.Login(ctx, input)
	if err != nil {
		view.MakeErrResp(c, err)
		return
	}

	c.Set("UserID", user.ID)

	view.MakeSuccessResp(c, http.StatusOK, view.MsgLoginSuccess(c), user)
}
