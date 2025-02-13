package users

import (
	"errors"
	"fmt"
	"net/http"
	"pt_aka_tech_test/helpers"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type userHandler struct {
	userService Service
}

func NewUserHandler(userService Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) Create(c *gin.Context) {
	var userRequest PostBody

	err := c.ShouldBindJSON(&userRequest)
	if err != nil {
			// Check if the error is a validation error
			var validationErrors validator.ValidationErrors
			if errors.As(err, &validationErrors) {
					errorMessages := []string{}
					for _, e := range validationErrors {
							errorMessage := fmt.Sprintf("Error on field %s, condition %s", e.Field(), e.ActualTag())
							errorMessages = append(errorMessages, errorMessage)
					}
					c.JSON(http.StatusBadRequest, gin.H{
                        "errors": errorMessages,
                    })
			}
	}

	newUser, err := h.userService.Register(userRequest)
	if err != nil {
		var httpErrCode = http.StatusInternalServerError;
		var errMessage = err.Error();
		if (err.Error() == helpers.UNAUTHORIZED) || (err.Error() == helpers.DUPLICATE) {
			httpErrCode = http.StatusConflict;
			errMessage = helpers.USER_ALREADY_EXIST;
		}

		c.JSON(httpErrCode, gin.H{
			"errors": errMessage,
		})
		return
	}

	newUserResponse := convertToUserResponse(newUser)

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": newUserResponse,
	})
}

func (h *userHandler) Login(c *gin.Context) {
	var logInRequest LogInRequest

	err := c.ShouldBindJSON(&logInRequest)
	if err != nil {
			// Check if the error is a validation error
			var validationErrors validator.ValidationErrors
			if errors.As(err, &validationErrors) {
					errorMessages := []string{}
					for _, e := range validationErrors {
							errorMessage := fmt.Sprintf("Error on field %s, condition %s", e.Field(), e.ActualTag())
							errorMessages = append(errorMessages, errorMessage)
					}
					c.JSON(http.StatusBadRequest, gin.H{
                        "errors": errorMessages,
                    })
			}
	}

	token, err := h.userService.LogIn(logInRequest)
	if err != nil {
		var httpErrCode = http.StatusInternalServerError;
		if (err.Error() == helpers.UNAUTHORIZED) {
			httpErrCode = http.StatusConflict;
		}
		var errMessage = err.Error()

		c.JSON(httpErrCode, gin.H{
			"errors": errMessage,
		})
		return
	}

	// c.SetCookie("token", token, 8080, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{
		"data": "success",
		"token": token,
	})
}

func convertToUserResponse(b User) UserResponse {
	return UserResponse{
        ID: b.ID,
		Name: b.Name,
		Email: b.Email,
    }
}