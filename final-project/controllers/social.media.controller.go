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

func CreateSocialmedia(c *gin.Context) {
	db := databases.GetDB()

	contentType := helpers.GetContentType(c)

	SocialMedia := models.SocialMedia{}

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	userData := c.MustGet("userData").(jwt.MapClaims)
	UserID := uint(userData["id"].(float64))

	SocialMediaRequest := models.SocialMedia{
		Name:           SocialMedia.Name,
		SocialMediaUrl: SocialMedia.SocialMediaUrl,
		UserId:         UserID,
	}

	err := db.Create(&SocialMediaRequest).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errror":  "Bad Request from create social media",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":               SocialMediaRequest.ID,
		"name":             SocialMediaRequest.Name,
		"social_media_url": SocialMediaRequest.SocialMediaUrl,
		"user_id":          SocialMediaRequest.UserId,
		"created_at":       SocialMediaRequest.CreatedAt,
	})
}

func GetSocialMedias(c *gin.Context) {
	db := databases.GetDB()
	SocialMedia := []models.SocialMedia{}

	err := db.Preload("User").Find(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ResponseSocialMedia := []models.ResponseSocialMedia{}

	for _, v := range SocialMedia {
		temp := models.ResponseSocialMedia{}

		temp.ID = v.ID
		temp.Name = v.Name
		temp.SocialMediaUrl = v.SocialMediaUrl
		temp.UserId = v.UserId
		temp.CreatedAt = v.CreatedAt
		temp.UpdatedAt = v.UpdatedAt
		temp.User.ID = v.User.ID
		temp.User.Username = v.User.Username
		temp.User.Email = v.User.Email

		ResponseSocialMedia = append(ResponseSocialMedia, temp)
	}

	c.JSON(http.StatusOK, gin.H{
		"social_medias": ResponseSocialMedia,
	})

}

func UpdateSocialmedia(c *gin.Context) {
	db := databases.GetDB()
	contentType := helpers.GetContentType(c)

	SocialMediaID, err := strconv.Atoi(c.Param("socialMediaId"))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "invalid parameter for SocialMediaID",
		})
		return
	}

	SocialMedia := models.SocialMedia{}

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	err = db.Model(&SocialMedia).Where("id = ?", SocialMediaID).Updates(models.SocialMedia{Name: SocialMedia.Name, SocialMediaUrl: SocialMedia.SocialMediaUrl}).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	SocialmediaResponse := models.SocialMedia{}

	db.First(&SocialmediaResponse, SocialMediaID)

	c.JSON(http.StatusOK, gin.H{
		"id":               SocialmediaResponse.ID,
		"name":             SocialmediaResponse.Name,
		"social_media_url": SocialmediaResponse.SocialMediaUrl,
		"user_id":          SocialmediaResponse.UserId,
		"updated_at":       SocialmediaResponse.UpdatedAt,
	})
}

func DeleteSocialmedia(c *gin.Context) {
	db := databases.GetDB()

	SocialMediaID, _ := strconv.Atoi(c.Param("socialMediaId"))

	SocialMedia := models.SocialMedia{}

	err := db.Delete(&SocialMedia, SocialMediaID).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "your social media has been successfully deleted",
	})
}
