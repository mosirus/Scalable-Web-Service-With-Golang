package controllers

import (
	"final-project/databases"
	"final-project/helpers"
	"final-project/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var (
	appJSON = "application/json"
)

func UserRegister(c *gin.Context) {
	db := databases.GetDB()
	contentType := helpers.GetContentType(c)

	User := models.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	err := db.Debug().Create(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"age":      User.Age,
		"email":    User.Email,
		"id":       User.ID,
		"username": User.Username,
	})
}

func UserLogin(c *gin.Context) {
	db := databases.GetDB()

	contentType := helpers.GetContentType(c)

	User := models.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	password := User.Password

	err := db.Debug().Where("email = ?", User.Email).Take(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email/passowrd",
		})
		return
	}

	comparePassword := helpers.ComparePassword([]byte(User.Password), []byte(password))

	if !comparePassword {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Password Incorrect",
		})
		return
	}

	token := helpers.GenerateToken(User.ID, User.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func UpdateUser(c *gin.Context) {
	db := databases.GetDB()

	contentType := helpers.GetContentType(c)
	user := models.User{}
	userIDFromParam, err := strconv.Atoi(c.Param("userId"))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "invalid parameter",
		})
		return
	}

	if contentType == appJSON {
		c.ShouldBindJSON(&user)
	} else {
		c.ShouldBind(&user)
	}

	err = db.Model(&user).Where("id = ?", userIDFromParam).Updates(models.User{Email: user.Email, Username: user.Username}).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	userResponse := models.User{}

	db.First(&userResponse, userIDFromParam)

	c.JSON(http.StatusOK, gin.H{
		"id":         userResponse.ID,
		"email":      userResponse.Email,
		"username":   userResponse.Username,
		"age":        userResponse.Age,
		"updated_at": userResponse.UpdatedAt,
	})
}

func DeleteUser(c *gin.Context) {

	db := databases.GetDB()

	userData := c.MustGet("userData").(jwt.MapClaims)
	UserID := int(userData["id"].(float64))

	user := models.User{}

	err := db.Delete(&user, UserID).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "your account has been successfully deleted",
	})
}
