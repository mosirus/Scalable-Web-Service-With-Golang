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

func CreatePhoto(c *gin.Context) {
	db := databases.GetDB()
	contentType := helpers.GetContentType(c)

	Photo := models.Photo{}

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	userData := c.MustGet("userData").(jwt.MapClaims)
	UserID := uint(userData["id"].(float64))

	PhotoRequest := models.Photo{
		Title:    Photo.Title,
		Caption:  Photo.Caption,
		PhotoUrl: Photo.PhotoUrl,
		UserId:   UserID,
	}

	err := db.Debug().Create(&PhotoRequest).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errror":  "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         PhotoRequest.ID,
		"title":      PhotoRequest.Title,
		"caption":    PhotoRequest.Caption,
		"photo_url":  PhotoRequest.PhotoUrl,
		"user_id":    PhotoRequest.UserId,
		"created_at": PhotoRequest.CreatedAt,
	})
}

func GetAllPhotos(c *gin.Context) {
	db := databases.GetDB()

	Photo := []models.Photo{}

	result := db.Preload("User").Find(&Photo)

	APIPhoto := []models.APIPhoto{}

	for _, v := range Photo {
		temp := models.APIPhoto{}

		temp.ID = v.ID
		temp.Title = v.Title
		temp.Caption = v.Caption
		temp.PhotoUrl = v.PhotoUrl
		temp.UserId = v.UserId
		temp.CreatedAt = v.CreatedAt
		temp.UpdatedAt = v.UpdatedAt
		temp.User.Username = v.User.Username
		temp.User.Email = v.User.Email

		APIPhoto = append(APIPhoto, temp)
	}

	if result.Error != nil {
		panic(result.Error)
	}

	c.JSON(http.StatusOK, APIPhoto)

}

func UpdatePhoto(c *gin.Context) {
	db := databases.GetDB()
	contentType := helpers.GetContentType(c)

	Photo := models.Photo{}

	photoID, err := strconv.Atoi(c.Param("photoId"))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "invalid parameter for photoID",
		})
		return
	}

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	err = db.Model(&Photo).Where("id = ?", photoID).Updates(models.Photo{Title: Photo.Title, Caption: Photo.Caption, PhotoUrl: Photo.PhotoUrl}).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	photoResponse := models.Photo{}

	db.First(&photoResponse, photoID)

	c.JSON(http.StatusOK, gin.H{
		"id":         photoResponse.ID,
		"title":      photoResponse.Title,
		"caption":    photoResponse.Caption,
		"photo_url":  photoResponse.PhotoUrl,
		"user_id":    photoResponse.UserId,
		"updated_at": photoResponse.UpdatedAt,
	})

}

func DeletePhoto(c *gin.Context) {
	db := databases.GetDB()

	Photo := models.Photo{}

	photoID, _ := strconv.Atoi(c.Param("photoId"))

	err := db.Delete(&Photo, photoID).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "your photo has been successfully deleted",
	})

}
