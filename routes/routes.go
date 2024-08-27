package routes

import (
    "github.com/gin-gonic/gin"
    "go-api-posts/controllers"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    r.Use(corsMiddleware())

    r.POST("/posts", controllers.CreatePost)

    r.GET("/posts", controllers.GetPosts)

    r.POST("/comments", controllers.CreateComment)

    return r
}

func corsMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
        
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}
