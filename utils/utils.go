package utils

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var Logger = log.Default()

func ReturnResponse(ctx *gin.Context, status int, message string, payload any) {
	ctx.Writer.Header().Set("Content-Type", "application/json")
	ctx.Status(status)
	ctx.JSON(status, gin.H{
		"status":  status,
		"message": message,
		"payload": payload,
	})
}

func HandleValidationErr(err error, ctx *gin.Context) {
	messages := []string{}

	switch errs := err.(type) {
	case validator.ValidationErrors:
		for _, err := range errs {
			messages = append(messages, MakeUserFriendlyMessage(&err))
		}
		ReturnResponse(ctx, http.StatusBadRequest, ValidationError, gin.H{"error": messages})
		return
	default:
		// log the error to check and fix for future
		Logger.Println("Unexpected Error: ", err)
		ReturnResponse(ctx, http.StatusBadRequest, ValidationError, nil)
	}
}

func MakeUserFriendlyMessage(fieldErr *validator.FieldError) string {
	switch (*fieldErr).Tag() {
	// can add multiple cases based on tag name and can generate message
	case "required":
		return fmt.Sprintf("'%s' is a required field", (*fieldErr).Field())
	default:
		return fmt.Sprintf("'%s' is invalid or missing", (*fieldErr).Field())
	}
}
