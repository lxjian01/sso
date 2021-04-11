package routers

import (
	"github.com/gin-gonic/gin"
	"sso/server/controllers"
)

func UserRoutes(route *gin.Engine) {
	user := route.Group("/user")
	{
		user.GET("/test", controllers.Test)
		//user.POST("/test", controllers.Test)
	}
}