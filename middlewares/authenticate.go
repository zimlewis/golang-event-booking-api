package middlewares

import (
	"net/http"

	"example.com/udemy_course/models/user"
	"example.com/udemy_course/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	uid, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
		return
	}

	_, err = user.Read(uid)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
		return
	}

	context.Set("uid", uid)
	context.Next()
} 