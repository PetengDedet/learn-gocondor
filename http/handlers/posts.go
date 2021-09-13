package handlers

import (
	"learn-gocondor/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllPosts(c *gin.Context) {
	var posts []models.Post

	DB.Order("created_at DESC").Limit(10).Find(&posts)

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   posts,
	})
}

func CreatePost(c *gin.Context) {

}
