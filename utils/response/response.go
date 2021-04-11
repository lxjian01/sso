package response

import (
	"github.com/gin-gonic/gin"
)

// 成功返回
func Success(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{"istrue":true,"data":data})
}

// 失败返回
func Failure(c *gin.Context,code int, data interface{}) {
	c.JSON(code, gin.H{"istrue":false,"data":data})
}

// 请求不存在404
func FailureNoRoute(c *gin.Context,data interface{}) {
	c.JSON(404, gin.H{"istrue":false,"data":data})
}

// 没有认证401
func FailureUnauthorized(c *gin.Context, msg string){
	c.JSON(401, gin.H{"istrue":false,"msg":msg})
}

// 没有权限403
func FailureForbidden(c *gin.Context, msg string){
	c.JSON(403, gin.H{"istrue":false,"msg":msg})
}

// 参数错误
func FailureParams(c *gin.Context, msg string){
	c.JSON(400, gin.H{"istrue":false,"msg":msg})
}

