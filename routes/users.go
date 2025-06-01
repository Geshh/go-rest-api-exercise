package routes

import (
	"net/http"
	"strconv"

	"example.com/event-booking-rest-api/models"
	"example.com/event-booking-rest-api/utils"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created!", "user": user})
}

func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate."})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successful.", "token": token})
}

func getUser(context *gin.Context) {
	userId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse user id."})
		return
	}

	user, err := models.GetUserById(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch user."})
		return
	}

	context.JSON(http.StatusOK, user)
}
