// pkg/middleware/cors.go
package middleware

import "github.com/gin-gonic/gin"

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
        c.Next()
    }
}
