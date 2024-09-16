package auth

import "github.com/gin-gonic/gin"

func RegisterRoutes(rg *gin.RouterGroup) {
    auth := rg.Group("/auth")
    {
        auth.POST("/register", Register)
        auth.POST("/login", Login)
    }
}
