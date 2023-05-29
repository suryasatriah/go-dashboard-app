package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/surysatriah/go-dashboard-app/internal/helper"
)

func Authentication() gin.HandlerFunc {
	return func(context *gin.Context) {
		err := helper.ValidateJWT(context)

		if err != nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Authentication required",
				"message": err.Error(),
			})
			return
		}
		context.Next()
	}
}
