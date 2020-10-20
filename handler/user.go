package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"startup/helper"
	"startup/user"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(service user.Service) *userHandler {
	return &userHandler{service}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errorMessage := helper.FormatValidationError(err)
		response := helper.APIRespose("Register account failed.", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		errorMessage := helper.FormatValidationError(err)
		response := helper.APIRespose("Register account failed.", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(newUser, "token")
	response := helper.APIRespose("Account has been created.", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context) {
	var input user.LoginUser

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorMessage := helper.FormatValidationError(err)
		response := helper.APIRespose("Login failed.", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loginUser, err := h.userService.Login(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIRespose("Login failed.", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := user.FormatUser(loginUser, "token")
	response := helper.APIRespose("Successfully login.", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
	return
}

func (h *userHandler) CheckEmailAvailability(c *gin.Context) {
	var input user.EmailCheckInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorMessage := helper.FormatValidationError(err)
		response := helper.APIRespose("Email checking failed.", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isEmailAvailable, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		errorMessage := gin.H{"error": "Server error"}
		response := helper.APIRespose("Email checking failed.", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	metaMessage := "Email has been registered"
	if isEmailAvailable {
		metaMessage = "Email is available"
	}
	data := gin.H{"is_available": isEmailAvailable}
	response := helper.APIRespose(metaMessage, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
	return
}
