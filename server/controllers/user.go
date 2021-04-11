package controllers

import (
	"github.com/gin-gonic/gin"
	"sso/log"
	"sso/server/services"
	"sso/server/validates"
	"sso/utils/response"
)

func Login(c *gin.Context) {
	var loginInput validates.LoginInput
	if err := c.ShouldBind(&loginInput); err != nil {
		response.FailureParams(c,err.Error())
		return
	}
	tokenResult,err1 := services.Login(loginInput.UserCode,loginInput.Password)
	if err1 != nil{
		response.FailureParams(c,err1.Error())
		return
	}
	response.Success(c,tokenResult)
	return
}

func Register(c *gin.Context) {
	loginInput := &validates.LoginInput{}//我这里分层了,主要是把参数验证这块单独分离出来了
	if err := c.ShouldBind(&loginInput); err != nil {
		log.Errorf("Login error by params %s \n",err)
	}
	response.Success(c,"111111111111")
}

func Test(c *gin.Context) {
	response.Success(c,"ok")
}
