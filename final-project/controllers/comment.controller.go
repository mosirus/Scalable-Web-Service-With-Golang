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

func CreateComment(c *gin.Context) {
	db := databases.GetDB()

	contentType := helpers.GetContentType(c)

	Comment := models.Comment{}

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	userData := c.MustGet("userData").(jwt.MapClaims)
	UserID := uint(userData["id"].(float64))

	CommentRequest := models.Comment{
		Message: Comment.Message,
		UserId:  UserID,
		PhotoId: Comment.PhotoId,
	}

	err := db.Debug().Omit("User").Create(&CommentRequest).Error
	//err := db.Debug().Omit(clause.Associations).Create(&CommentRequest).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errror":  "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         CommentRequest.ID,
		"message":    CommentRequest.Message,
		"photo_id":   CommentRequest.PhotoId,
		"user_id":    CommentRequest.UserId,
		"created_at": CommentRequest.CreatedAt,
	})
}

func GetComments(c *gin.Context) {
	db := databases.GetDB()

	Comment := []models.Comment{}

	err := db.Preload("User").Preload("Photo").Find(&Comment).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	APIComment := []models.CommentResponse{}

	for _, v := range Comment {
		temp := models.CommentResponse{}

		temp.ID = v.ID
		temp.Message = v.Message
		temp.PhotoId = v.PhotoId
		temp.UserId = v.UserId
		temp.CreatedAt = v.CreatedAt
		temp.UpdatedAt = v.UpdatedAt
		temp.User.ID = v.User.ID
		temp.User.Email = v.User.Email
		temp.User.Username = v.User.Username
		temp.Photo.ID = v.Photo.ID
		temp.Photo.Title = v.Photo.Title
		temp.Photo.Caption = v.Photo.Caption
		temp.Photo.PhotoUrl = v.Photo.PhotoUrl
		temp.Photo.UserId = v.Photo.UserId

		APIComment = append(APIComment, temp)
	}
	c.JSON(http.StatusOK, APIComment)
}

func UpdateComment(c *gin.Context) {
	db := databases.GetDB()
	contentType := helpers.GetContentType(c)

	commentID, _ := strconv.Atoi(c.Param("commentId"))
	Comment := models.Comment{}

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	err := db.Model(&Comment).Where("id = ?", commentID).Updates(models.Comment{Message: Comment.Message}).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	CommentResponse := models.Comment{}

	db.Preload("Photo").Preload("User").First(&CommentResponse, commentID)

	Response := models.CommentResponseUpdate{}
	Response.Id = CommentResponse.ID
	Response.Title = CommentResponse.Photo.Title
	Response.Caption = CommentResponse.Photo.Caption
	Response.PhotoUrl = CommentResponse.Photo.PhotoUrl
	Response.UserId = CommentResponse.Photo.UserId
	Response.UpdatedAt = CommentResponse.UpdatedAt

	c.JSON(http.StatusOK, Response)
}

func DeleteComment(c *gin.Context) {
	db := databases.GetDB()

	commentID, _ := strconv.Atoi(c.Param("commentId"))
	Comment := models.Comment{}

	err := db.Delete(&Comment, commentID).Error

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
