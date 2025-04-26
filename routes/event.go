package routes

import (
	"net/http"
	"strconv"

	"example.com/udemy_course/models/event"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := event.ReadAll()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64) 

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	e, err := event.Read(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, *e)
}

func createEvent(context *gin.Context) {
	var e event.Event

	err := context.ShouldBindJSON(&e)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	
	uid := context.GetInt64("uid")
	e.UserId = uid
	err = event.Create(e)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.Status(http.StatusOK)
}

func updateEvent(context *gin.Context) {
	var e event.Event
	
	err := context.ShouldBindJSON(&e)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, err := strconv.ParseInt(context.Param("id"), 10, 64) 
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	oldEvent, err := event.Read(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	e.Id = oldEvent.Id
	e.UserId = oldEvent.UserId
	uid := context.GetInt64("uid")
	if e.UserId != uid {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "you are not the creator"})
		return
	}

	err = event.Update(e)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	context.Status(http.StatusAccepted)
}

func deleteEvent(context *gin.Context) {
	uid := context.GetInt64("uid")
	id, err := strconv.ParseInt(context.Param("id"), 10, 64) 
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	e, err := event.Read(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if e.UserId != uid {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "you are not the creator"})
		return
	}


	err = event.Delete(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	context.Status(http.StatusOK)
}

