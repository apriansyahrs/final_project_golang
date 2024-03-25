package controllers

import (
	"final_project_golang/database"
	"final_project_golang/helpers"
	"final_project_golang/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func CreatePhoto(c *gin.Context) {

	db := database.GetDB()
	userId := uint(c.MustGet("userData").(jwt.MapClaims)["id"].(float64))

	contentType := helpers.GetContentType(c)

	Photo := models.Photo{}

	if contentType == appJson {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userId

	if err := db.Debug().Create(&Photo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"id":        Photo.ID,
		"title":     Photo.Title,
		"caption":   Photo.Caption,
		"photo_url": Photo.PhotoURL,
		"user_id":   Photo.UserID,
	})

}

func GetPhoto(c *gin.Context) {
	db := database.GetDB()
	userId := uint(c.MustGet("userData").(jwt.MapClaims)["id"].(float64))
	photos := []models.Photo{}

	if err := db.Debug().Preload("User").Where("user_id = ?", userId).Find(&photos).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	var response []gin.H
	for _, photo := range photos {
		photoData := gin.H{
			"id":        photo.ID,
			"caption":   photo.Caption,
			"title":     photo.Title,
			"photo_url": photo.PhotoURL,
			"user_id":   photo.UserID,
		}
		userData := gin.H{
			"id":       photo.User.ID,
			"email":    photo.User.Email,
			"username": photo.User.Username,
		}
		photoData["user"] = userData
		response = append(response, photoData)
	}
	c.JSON(http.StatusOK, response)
}

func GetPhotoById(c *gin.Context) {
	db := database.GetDB()
	userId := uint(c.MustGet("userData").(jwt.MapClaims)["id"].(float64))
	photoId := c.Param("photoId")
	photo := models.Photo{}

	if err := db.Debug().Preload("User").Where("user_id = ? AND id = ?", userId, photoId).First(&photo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Photo not found",
		})
		return
	}

	responseData := gin.H{
		"id":        photo.ID,
		"caption":   photo.Caption,
		"title":     photo.Title,
		"photo_url": photo.PhotoURL,
		"user_id":   photo.UserID,
		"user": gin.H{
			"id":       photo.User.ID,
			"email":    photo.User.Email,
			"username": photo.User.Username,
		},
	}

	c.JSON(http.StatusOK, responseData)
}

func UpdatePhoto(c *gin.Context) {
	db := database.GetDB()

	photoId := c.Param("photoId")

	userId := uint(c.MustGet("userData").(jwt.MapClaims)["id"].(float64))

	contentType := helpers.GetContentType(c)

	Photo := models.Photo{}
	db.First(&Photo, photoId)

	if contentType == appJson {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userId

	if err := db.Model(&Photo).Where("id = ?", photoId).Updates(&Photo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"id":        Photo.ID,
		"caption":   Photo.Caption,
		"title":     Photo.Title,
		"photo_url": Photo.PhotoURL,
		"user_id":   Photo.UserID,
	})

}

func DeletePhoto(c *gin.Context) {
	db := database.GetDB()
	photoId := c.Param("photoId")
	Photo := models.Photo{}
	db.First(&Photo, photoId)
	db.Delete(&Photo)
	c.JSON(http.StatusOK, gin.H{
		"message": "Your photo has been successfully deleted",
	})
}
