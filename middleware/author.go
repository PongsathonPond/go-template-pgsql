package middleware

import (
	"context"
	"net/http"
	"strings"

	"idev-cms-service/app/view"
	"idev-cms-service/service/roles/inout"
	"idev-cms-service/service/util"

	"github.com/gin-gonic/gin"
)

func (mid *Service) Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		header := c.GetHeader("Authorization")

		if len(header) == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// extract token
		token := strings.ReplaceAll(header, "Bearer ", "")

		user, err := mid.UserService.ReadByToken(ctx, &token)
		if err != nil {
			view.MakeErrResp(c, util.RepoReadErr(err))
			c.Abort()
		}

		resp := mid.RoleService.CheckPermission(ctx, &inout.CheckPermissionInput{
			RoleID:   user.Role,
			Method:   c.Request.Method,
			EndPoint: c.FullPath(),
			Service:  "authentication",
		})

		if !resp {
			c.AbortWithStatus(http.StatusForbidden)
		}
	}
}
