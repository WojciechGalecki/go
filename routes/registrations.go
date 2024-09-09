package routes

import (
	"net/http"
	"strconv"

	"example.com/app/middlewares"
	"example.com/app/models"
	"github.com/gin-gonic/gin"
)

func RegisterForEvent(context *gin.Context) {
	userId := context.GetInt64(middlewares.UserIdKey)
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"message": "Invalid event id"},
		)
		return
	}

	event, err := models.GetEvent(eventId)

	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"message": "Couldn't fetch event"},
		)
		return
	}

	err = event.Register(userId)

	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"message": "Couldn't register event"},
		)
		return
	}

	context.JSON(
		http.StatusCreated,
		gin.H{"message": "Event registered"},
	)
}

func CancelRegistration(context *gin.Context) {
	userId := context.GetInt64(middlewares.UserIdKey)
	eventId, _ := strconv.ParseInt(context.Param("id"), 10, 64)

	var event models.Event
	event.ID = eventId

	err := event.CancelRegistration(userId)

	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"message": "Couldn't cancel registration"},
		)
		return
	}

	context.JSON(
		http.StatusOK,
		gin.H{"message": "Registration cancelled"},
	)
}
