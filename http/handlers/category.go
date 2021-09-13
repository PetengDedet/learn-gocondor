package handlers

import (
	"errors"
	"fmt"
	"learn-gocondor/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

type CategoryCreateRequestValidator struct {
	Name   string `form:"name" json:"name" binding:"required,printascii,min=1,max=255"`
	Parent string `form:"parent" json:"parent" binding:"omitempty,required,printascii,min=1,max=255"`
}

func GetCategories(c *gin.Context) {
	var categories []models.Category

	DB.Order("created_at DESC").Limit(10).Find(&categories)

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   categories,
	})
}

func GetCategoryById(c *gin.Context) {
	var category models.Category
	result := DB.Preload("Parent").First(&category, c.Param("id"))

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "fail",
			"message": "Category not found!",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   category,
	})
}

func CreateCategory(c *gin.Context) {
	var request CategoryCreateRequestValidator
	if err := c.ShouldBind(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"error":  err.Error(),
		})

		return
	}

	// Verify unique slug
	slug := slug.Make(request.Name)
	isSlugExists, _ := IsSlugExists(slug)
	if isSlugExists {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": fmt.Sprintf("Category with slug `%s` already exists", slug),
		})

		return
	}

	newCategory := models.Category{
		Name: request.Name,
		Slug: slug,
	}

	// Verify parent exist
	if len(request.Parent) > 0 {
		isParentSlugExists, parentID := IsSlugExists(request.Parent)
		if !isParentSlugExists {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": fmt.Sprintf("Parent category `%s` does not exist in the database", request.Parent),
			})

			return
		}

		newCategory.ParentID = parentID
	}

	DB.Create(&newCategory)

	c.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"data":   &newCategory,
	})

}

func IsSlugExists(slug string, skipID ...uint) (isExist bool, ID *uint) {
	var category models.Category
	exists := DB.Where("slug = ?", slug).Limit(1).Find(&category)

	isExist = exists.RowsAffected > 0
	ID = &category.ID

	if len(skipID) > 0 {
		isExist = isExist && category.ID != skipID[0]
		return
	}

	return
}
