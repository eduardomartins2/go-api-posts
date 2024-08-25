package routes

import (
    "github.com/gin-gonic/gin"
    "go-api-posts/controllers"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    // Rota para criação de um post
    r.POST("/posts", controllers.CreatePost)

    // Rota para listar todos os posts
    r.GET("/posts", controllers.GetPosts)

    // Rota para criar um comentário
    r.POST("/comments", controllers.CreateComment)

    return r
}
