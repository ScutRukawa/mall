package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST,OPTIONS,GET,DELETE")
		// c.Header("Access-Control-Max-Age", "3600")
		c.Header("Access-Control-Allow-Headers", "accept,x-requested-with,Content-Type,Access-Control-Allow-Origin,Access-Control-Allow-Headers,x-token")
		c.Header("Access-Control-Allow-Credentials", "true")
		// c.Header("Access-Control-Allow-Origin", "http://192.168.10.118:8070")
		c.Header("Access-Control-Expose-Headers", "accept,x-requested-with,Content-Type,Access-Control-Allow-Origin,Access-Control-Allow-Headers,x-token")

		if method == "OPTION" {
			c.AbortWithStatus(http.StatusNoContent)
		}
	}
}
