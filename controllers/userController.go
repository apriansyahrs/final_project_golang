package controllers

import (
	"final_project_golang/database"
	"final_project_golang/helpers"
	"final_project_golang/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var (
	appJson = "application/json"
)

func UserRegister(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)

	User := models.User{}

	if contentType == appJson {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	if err := db.Debug().Create(&User).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
			"msg":   err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":                User.ID,
		"email":             User.Email,
		"username":          User.Username,
		"age":               User.Age,
		"profile_image_url": User.ProfileImageURL,
	})
}

func UserLogin(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)

	User := models.User{}
	if contentType == appJson {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	password := User.Password

	if err := db.Debug().Where("email = ?", User.Email).Take(&User).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Email or Password wrong",
		})
		return
	}

	comparePass := helpers.ComparePass([]byte(User.Password), []byte(password))
	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Email or Password wrong",
		})
		return
	}
	token := helpers.GenerateToken(User.ID, User.Email)
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func UserUpdate(c *gin.Context) {
	db := database.GetDB()

	contentType := helpers.GetContentType(c)

	User := models.User{}

	userId, _ := strconv.Atoi(c.Param("userId"))

	db.First(&User, userId)

	if contentType == appJson {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}
	updates := map[string]interface{}{
		"Email":           User.Email,
		"Username":        User.Username,
		"Age":             User.Age,
		"ProfileImageURL": User.ProfileImageURL,
	}
	err := db.Model(&User).Where("id = ?", userId).Updates(updates).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":                User.ID,
		"email":             User.Email,
		"username":          User.Username,
		"age":               User.Age,
		"profile_image_url": User.ProfileImageURL,
	})
}

func UserDelete(c *gin.Context) {
	db := database.GetDB()

	User := models.User{}
	userId := uint(c.MustGet("userData").(jwt.MapClaims)["id"].(float64))

	err := db.Delete(&User, userId).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your account has been successfully deleted",
	})
}
