package server

import (
	"github.com/gin-gonic/gin"
	"net"
	"os"
	"sso/config"
	"sso/log"
	"sso/server/controllers"
	"sso/server/middlewares"
	"sso/server/routers"
	"sso/utils/response"
	"strconv"
)

func StartServer(c *config.ServerConfig) {
	router := gin.Default()
	// 添加自定义的logger中间件
	router.Use(middlewares.RequestId(), middlewares.Logger(), gin.Recovery())
	// 404
	router.NoRoute(func(c *gin.Context) {
		response.FailureNoRoute(c,"请求方法不存在")
	})
	// ping healthy
	router.GET("/ping", func(c *gin.Context) {
		response.Success(c,"ok")
	})
	// login
	router.POST("/login", controllers.Login)
	// register
	router.POST("/register", controllers.Register)
	// 添加jwt认证中间件
	router.Use(middlewares.Auth(), gin.Recovery())
	// 添加路由
	routers.UserRoutes(router)      //Added all user routers
	// 拼接host
	Host := c.Host
	Port := strconv.Itoa(c.Port)
	addr := net.JoinHostPort(Host, Port)
	log.Info("Start HTTP server at", addr)
	err := router.Run(addr)
	if err != nil{
		log.Error("Start server error by",err)
		os.Exit(1)
	}
	log.Info("Start server ok")
}
