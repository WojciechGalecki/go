package routes

import (
	"net/http"

	"example.com/app/models"
	"github.com/gin-gonic/gin"
)

func Signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"message": "Couldn't parse request data"},
		)
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"message": "Couldn't create user"},
		)
		return
	}

	context.JSON(
		http.StatusCreated,
		gin.H{"message": "User created"})
}

func Login(context *gin.Context) {
  var user models.User
	err := context.ShouldBindJSON(&user)

  if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"message": "Couldn't parse request data"},
		)
		return
	}

  err = user.ValidateCredentials()

  if err != nil {
		context.JSON(
			http.StatusUnauthorized,
			gin.H{"message": "Couldn't authenticate user"},
		)
		return
	}

  context.JSON(
		http.StatusOK,
		gin.H{"message": "Login successful"})
}
