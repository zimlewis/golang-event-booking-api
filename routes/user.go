package routes

import (
	"net/http"

	"example.com/udemy_course/models/user"
	"example.com/udemy_course/utils"

	"github.com/gin-gonic/gin"
)

func signUp(context *gin.Context) {
	var u user.User

	err := context.ShouldBindJSON(&u)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	u.Password, err = utils.HashPassword(u.Password)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	err = user.Create(u)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.Status(http.StatusCreated)
}

func signIn(context *gin.Context) {
	var u user.User
	
	err := context.ShouldBindJSON(&u)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	token, err := user.ValidateUser(u)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}