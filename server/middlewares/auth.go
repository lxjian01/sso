package middlewares

import (
	"github.com/gin-gonic/gin"
	"sso/log"
	"sso/utils/jwt"
	"sso/utils/response"
)

// 判断用户书否登陆中间件
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			response.FailureUnauthorized(c, "请求未携带token，无权限访问")
			c.Abort()
			return
		}
		log.Info("get token: ", token)
		j := jwt.NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		log.Info("claims is ", claims)
		if err != nil {
			response.FailureUnauthorized(c, err.Error())
			c.Abort()
			return
		}
		// 继续交由下一个路由处理,并将解析出的信息传递下去
		c.Set("claims", claims)
	}
}