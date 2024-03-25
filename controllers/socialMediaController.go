package controllers

import (
	"final_project_golang/database"
	"final_project_golang/helpers"
	"final_project_golang/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func CreateSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userId := uint(c.MustGet("userData").(jwt.MapClaims)["id"].(float64))

	contentType := helpers.GetContentType(c)

	sosmed := models.SocialMedia{}

	if contentType == appJson {
		c.ShouldBindJSON(&sosmed)
	} else {
		c.ShouldBind(&sosmed)
	}

	sosmed.UserID = userId

	if err := db.Debug().Create(&sosmed).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"id":               sosmed.ID,
		"name":             sosmed.Name,
		"social_media_url": sosmed.SocialMediaURL,
		"user_id":          sosmed.UserID,
	})

}

func GetSocialMedia(c *gin.Context) {
	db := database.GetDB()

	userId := uint(c.MustGet("userData").(jwt.MapClaims)["id"].(float64))

	sosmeds := []models.SocialMedia{}

	if err := db.Debug().Preload("User").Where("user_id = ?", userId).Find(&sosmeds).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	var data []gin.H

	for _, sosmed := range sosmeds {
		sosmedData := gin.H{
			"id":               sosmed.ID,
			"name":             sosmed.Name,
			"social_media_url": sosmed.SocialMediaURL,
			"user_id":          sosmed.UserID,
			"user": gin.H{
				"id":       sosmed.User.ID,
				"email":    sosmed.User.Email,
				"username": sosmed.User.Username,
			},
		}
		data = append(data, sosmedData)
	}

	c.JSON(http.StatusOK, data)
}

func GetSocialMediaById(c *gin.Context) {
	db := database.GetDB()

	sosmedId := c.Param("socialMediaId")

	userId := uint(c.MustGet("userData").(jwt.MapClaims)["id"].(float64))

	sosmed := models.SocialMedia{}
	if err := db.Debug().Preload("User").Where("user_id = ? AND id = ?", userId, sosmedId).First(&sosmed).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Social media not found",
		})
		return
	}

	responseData := gin.H{
		"id":               sosmed.ID,
		"name":             sosmed.Name,
		"social_media_url": sosmed.SocialMediaURL,
		"user_id":          sosmed.UserID,
		"user": gin.H{
			"id":       sosmed.User.ID,
			"email":    sosmed.User.Email,
			"username": sosmed.User.Username,
		},
	}

	c.JSON(http.StatusOK, responseData)
}

func UpdateSocialMedia(c *gin.Context) {
	db := database.GetDB()
	sosmedId := c.Param("socialMediaId")
	userId := uint(c.MustGet("userData").(jwt.MapClaims)["id"].(float64))
	contentType := helpers.GetContentType(c)
	sosmed := models.SocialMedia{}
	db.First(&sosmed, sosmedId)

	if contentType == appJson {
		c.ShouldBindJSON(&sosmed)
	} else {
		c.ShouldBind(&sosmed)
	}

	sosmed.UserID = userId

	if err := db.Model(&sosmed).Where("id = ?", sosmedId).Updates(&sosmed).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"id":               sosmed.ID,
		"name":             sosmed.Name,
		"social_media_url": sosmed.SocialMediaURL,
		"user_id":          sosmed.UserID,
	})

}

func DeleteSocialMedia(c *gin.Context) {
	db := database.GetDB()
	sosmedId := c.Param("socialMediaId")
	sosmed := models.SocialMedia{}
	db.First(&sosmed, sosmedId)
	db.Delete(&sosmed)
	c.JSON(http.StatusOK, gin.H{
		"message": "Your social media has been successfully deleted",
	})
}
