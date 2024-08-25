package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "go-api-posts/models"
    "go-api-posts/database"
)

// Função para criar um post
func CreatePost(c *gin.Context) {
    var post models.Post
    if err := c.ShouldBindJSON(&post); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    database.DB.Create(&post)
    c.JSON(http.StatusOK, post)
}

// Função para listar todos os posts
func GetPosts(c *gin.Context) {
    var posts []models.Post
    database.DB.Preload("Comments").Find(&posts)
    c.JSON(http.StatusOK, posts)
}

// Função para criar um comentário
func CreateComment(c *gin.Context) {
    var comment models.Comment
    if err := c.ShouldBindJSON(&comment); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    database.DB.Create(&comment)
    c.JSON(http.StatusOK, comment)
}
