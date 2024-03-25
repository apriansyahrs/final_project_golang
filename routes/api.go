package routes

import (
	"final_project_golang/controllers"
	"final_project_golang/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userGroup := r.Group("/users")
	{
		userGroup.POST("/register", controllers.UserRegister)
		userGroup.POST("/login", controllers.UserLogin)
		userGroup.PUT("/:userId", middlewares.Authentication(), controllers.UserUpdate)
		userGroup.DELETE("/delete", middlewares.Authentication(), controllers.UserDelete)
	}

	photoGroup := r.Group("/photos")
	{
		photoGroup.Use(middlewares.Authentication())
		photoGroup.POST("/", controllers.CreatePhoto)
		photoGroup.GET("/", controllers.GetPhoto)
		photoGroup.GET("/:photoId", middlewares.PhotoAuthorization(), controllers.GetPhotoById)
		photoGroup.PUT("/:photoId", middlewares.PhotoAuthorization(), controllers.UpdatePhoto)
		photoGroup.DELETE("/:photoId", middlewares.PhotoAuthorization(), controllers.DeletePhoto)
	}

	commentGroup := r.Group("/comments")
	{
		commentGroup.Use(middlewares.Authentication())
		commentGroup.POST("/", controllers.CreateComment)
		commentGroup.GET("/", controllers.GetComment)
		commentGroup.GET("/:commentId", middlewares.CommentAuthorization(), controllers.GetCommentById)
		commentGroup.PUT("/:commentId", middlewares.CommentAuthorization(), controllers.UpdateComment)
		commentGroup.DELETE("/:commentId", middlewares.CommentAuthorization(), controllers.DeleteComment)
	}

	socialmediasGroup := r.Group("/socialmedias")
	{
		socialmediasGroup.Use(middlewares.Authentication())
		socialmediasGroup.POST("/", controllers.CreateSocialMedia)
		socialmediasGroup.GET("/", controllers.GetSocialMedia)
		socialmediasGroup.GET("/:socialMediaId", middlewares.SocialMediaAuthorization(), controllers.GetSocialMediaById)
		socialmediasGroup.PUT("/:socialMediaId", middlewares.SocialMediaAuthorization(), controllers.UpdateSocialMedia)
		socialmediasGroup.DELETE("/:socialMediaId", middlewares.SocialMediaAuthorization(), controllers.DeleteSocialMedia)
	}

	return r
}
