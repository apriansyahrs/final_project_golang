package middlewares

import (
	"final_project_golang/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifiyToken, err := helpers.VerifiyToken(c)
		_ = verifiyToken

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthenticated",
				"message": err.Error(),
			})
			return
		}
		c.Set("userData", verifiyToken)
		c.Next()
	}
}
