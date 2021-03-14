package middleware

import (
	"cstore/common"
	"cstore/serializer"
	"github.com/gin-gonic/gin"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		code = common.SUCCESS
		token := c.GetHeader("Authorization")
		if token == "" {
			code = 401
			c.JSON(code, &serializer.Response{
				Status: code,
				Msg:    common.GetMsg(code),
			})
			c.Abort()
		} else {
			checkToken, err := common.CheckToken(token)
			if err != nil {
				code = 401
			} else {
				if time.Now().Unix() > checkToken.ExpiresAt {
					code = common.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				}
			}
		}
		if code != common.SUCCESS {
			c.JSON(code, &serializer.Response{
				Status: code,
				Msg:    common.GetMsg(code),
			})
			c.Abort()
		}
		c.Next()
	}
}
