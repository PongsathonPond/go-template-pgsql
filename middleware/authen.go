package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (mid *Service) Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		header := c.GetHeader("Authorization")

		if len(header) == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// extract token
		token := strings.ReplaceAll(header, "Bearer ", "")

		resp := mid.AuthenService.VerifyTokens(ctx, &token)
		if resp.Verify {
			c.Set("UserID", resp.UserID)
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
