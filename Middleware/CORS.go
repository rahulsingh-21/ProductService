package middleware

import "github.com/gin-gonic/gin"

func CORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PATCH, PUT, DELETE")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, X-Requested-With, "+
			"Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, refresh_token")
		ctx.Writer.Header().Set("Access-Control-Expose-Headers", "Auth_token, Refresh_token")

		ctx.Next()
	}
}
