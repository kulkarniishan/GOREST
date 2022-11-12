package middleware

import "github.com/gin-gonic/gin"

func Authorization() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		"ishan": "123Testing",
	})
}
