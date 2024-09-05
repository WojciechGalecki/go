package middlewares

import (
	"net/http"

	"example.com/app/utils"
	"github.com/gin-gonic/gin"
)

const UserIdKey = "userId"

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{"message": "Not authorized"},
		)
		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{"message": "Not authorized"},
		)
		return
	}

	context.Set(UserIdKey, userId)
	context.Next()
}