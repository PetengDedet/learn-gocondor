package handlers

import (
	"net/http"

	"learn-gocondor/http/input"
	"learn-gocondor/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserCollection struct {
	Name     string
	Username string
}

func GetUser(c *gin.Context) {

	var users []models.User
	// result := DB.Model(&models.User{}).Find(&UserCollection{})

	DB.Find(&users)
	// json.Unmarshal(result.)
	// fmt.Println(&result)
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
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

	exists := UsernameOrEmailExists(request.Username, request.Email)
	if exists {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"error":  "Email or Password already exists",
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

func UsernameOrEmailExists(username, email string) bool {
	var user models.User
	exists := DB.Where("username = ?", username).Or("email = ?", email).First(&user)

	return exists.RowsAffected > 0
}
