package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"learn-gocondor/http/input"
	"learn-gocondor/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func GetUsers(c *gin.Context) {

	var users []models.User
	DB.Order("created_at DESC").Limit(10).Find(&users)

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   users,
	})
}

func CreateUser(c *gin.Context) {

	var request input.UserCreateRequestValidator

	if err := c.ShouldBind(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"error":  err.Error(),
		})

		return
	}

	// Check Username
	isUsernameExists := IsUsernameExists(request.Username)
	if isUsernameExists {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"error":  fmt.Sprintf("Username `%s` has been taken", request.Username),
		})

		return
	}

	// Check Email
	isEmailExists := IsEmailExists(request.Email)
	if isEmailExists {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"error":  fmt.Sprintf("Email `%s` has been taken", request.Email),
		})

		return
	}

	hashedPassword, _ := HashPassword(request.Password)
	newUser := models.User{
		Username: request.Username,
		Name:     request.Name,
		Email:    request.Email,
		Password: hashedPassword,
	}

	DB.Create(&newUser)

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "User created!",
		"data":    &newUser,
	})
}

func GetUserById(c *gin.Context) {
	var user models.User

	result := DB.First(&user, c.Param("id"))

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "fail",
			"message": "User not found!",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   user,
	})
}

func UpdateUserById(c *gin.Context) {
	var user models.User

	result := DB.First(&user, c.Param("id"))

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "fail",
			"message": "User not found!",
		})

		return
	}

	// DB.Model(&user)
	var request input.UserUpdateRequestValidator
	if err := c.ShouldBind(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"error":  err.Error(),
		})

		return
	}

	// Update name only if set
	if len(request.Name) > 0 {
		user.Name = request.Name
	}

	// Update username if set
	if len(request.Username) > 0 {
		// Check exist username
		isUsernameExists := IsUsernameExists(request.Username, user.ID)
		if isUsernameExists {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": fmt.Sprintf("Username `%s` has been taken", request.Username),
			})

			return
		}

		user.Username = request.Username
	}

	// Update email only if set
	if len(request.Email) > 0 {
		// Check exist email
		isEmailExists := IsEmailExists(request.Email, user.ID)
		if isEmailExists {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": fmt.Sprintf("Email `%s` has been taken", request.Email),
			})

			return
		}

		user.Email = request.Email
	}

	DB.Save(&user)

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "User updated!",
		"data":    &user,
	})
}

func DeleteUserById(c *gin.Context) {
	var user models.User

	result := DB.First(&user, c.Param("id"))

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "fail",
			"message": "User not found!",
		})

		return
	}

	DB.Delete(&user)

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "User deleted!",
		"data":    &user.ID,
	})
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func IsUsernameExists(username string, skipID ...uint) bool {
	var user models.User
	exists := DB.Where("username = ?", username).Limit(1).Find(&user)

	if len(skipID) > 0 {
		return exists.RowsAffected > 0 && user.ID != skipID[0]
	}

	return exists.RowsAffected > 0
}

func IsEmailExists(email string, skipID ...uint) bool {
	var user models.User
	exists := DB.Where("email = ?", email).Limit(1).Find(&user)

	if len(skipID) > 0 {
		return exists.RowsAffected > 0 && user.ID != skipID[0]
	}

	return exists.RowsAffected > 0
}
