package routes

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"example.com/udemy_course/models/event"
	"example.com/udemy_course/models/registration"
	"github.com/gin-gonic/gin"
)

func register(context *gin.Context) {
	uid := context.GetInt64("uid")
	eid, err := strconv.ParseInt(context.Param("id"), 10, 64) 
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	e, err := event.Read(eid)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	if e.UserId == uid {
		context.JSON(http.StatusConflict, gin.H{"message": "you cannot register on your own event"})
		return
	}

	r := registration.New(uid, eid)
	err = registration.Create(*r)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.Status(http.StatusOK)
}

func cancel(context *gin.Context) {
	uid := context.GetInt64("uid")
	eid, err := strconv.ParseInt(context.Param("id"), 10, 64) 
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	e, err := event.Read(eid)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if e.DateTime.Before(time.Now()) {
		context.JSON(http.StatusBadRequest, gin.H{"message": "you cannot cancel event that is already happened"})
		return
	}

	r, err := registration.FindWithUserIdAndEventId(uid, eid)
	fmt.Println(err)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return 
	}

	err = registration.Delete(r.Id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return 
	}
	context.Status(http.StatusOK)
}