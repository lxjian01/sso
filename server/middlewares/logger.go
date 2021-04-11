package middlewares

import "C"
import (
	"github.com/gin-gonic/gin"
	"sso/log"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		host := c.Request.Host
		url := c.Request.URL
		method := c.Request.Method
		requestId := c.GetString("X-Request-Id")
		startTime := time.Now().Format("2006-01-02 15:04:05")
		log.Infof("%s %s %s%s, request id is %s \n",startTime, method, host, url,requestId)
		c.Next()
		endTime := time.Now().Format("2006-01-02 15:04:05")
		log.Infof("%s %s %s%s, response status is %d, request id is %s \n",endTime, method, host, url,c.Writer.Status(),requestId)
	}
}
