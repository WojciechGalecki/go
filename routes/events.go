package routes

import (
	"net/http"
	"strconv"

	"example.com/app/middlewares"
	"example.com/app/models"
	"github.com/gin-gonic/gin"
)

func GetEvents(context *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"massage": "Couldn't fetch events"},
		)
		return
	}
	context.JSON(http.StatusOK, events)
}

func GetEvent(context *gin.Context) {
	id, err := extractId(context)

	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"message": "Invalid event id"},
		)
		return
	}

	event, err := models.GetEvent(id)

	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"message": "Couldn't fetch event"},
		)
		return
	}

	context.JSON(http.StatusOK, event)
}

func CreateEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"message": "Couldn't parse request data"},
		)
		return
	}

	userId := context.GetInt64(middlewares.UserIdKey)
	event.UserID = userId

	err = event.Save()

	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"message": "Couldn't create event"},
		)
		return
	}

	context.JSON(
		http.StatusCreated,
		gin.H{
			"message": "Event created",
			"event":   event,
		})
}

func UpdateEvent(context *gin.Context) {
	id, err := extractId(context)

	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"message": "Invalid event id"},
		)
		return
	}

	userId := extractUserId(context)
	event, err := models.GetEvent(id)

	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"message": "Couldn't fetch event"},
		)
		return
	}

	if event.UserID != userId {
		context.JSON(
			http.StatusUnauthorized,
			gin.H{"message": "Not authorized to update event"},
		)
		return
	}

	var updatedEvent models.Event

	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"message": "Couldn't parse request data"},
		)
		return
	}

	updatedEvent.ID = id

	err = updatedEvent.Update()

	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"message": "Couldn't update event"},
		)
		return
	}

	context.JSON(
		http.StatusOK,
		gin.H{"message": "Event updated"},
	)
}

func DeleteEvent(context *gin.Context) {
	id, err := extractId(context)

	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"message": "Invalid event id"},
		)
		return
	}

	userId := extractUserId(context)
	event, err := models.GetEvent(id)

	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"message": "Couldn't fetch event"},
		)
		return
	}

	if event.UserID != userId {
		context.JSON(
			http.StatusUnauthorized,
			gin.H{"message": "Not authorized to delete event"},
		)
		return
	}

	err = event.Delete()

	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"message": "Couldn't delete event"},
		)
		return
	}

	context.JSON(
		http.StatusOK,
		gin.H{"message": "Event deleted"},
	)
}

func extractId(context *gin.Context) (int64, error) {
	return strconv.ParseInt(context.Param("id"), 10, 64)
}

func extractUserId(context *gin.Context) int64 {
	return context.GetInt64(middlewares.UserIdKey)
}
