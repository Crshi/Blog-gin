package jwt

import (
	"net/http"
	"time"

	"github.com/Crshi/Blog/pkg/e"
	"github.com/Crshi/Blog/pkg/util"
	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = e.SUCCESS
		token := c.Query("token")
		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			//检查Token是否可以解析
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else {
				//检查Token是否过期
				if time.Now().Unix() > claims.ExpiresAt {
					code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				}
			}
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
