package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"~/go-rest-api/internal/input"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)


func postBookHandler(c *gin.Context) {
	var bookInput input.BookInput

	err := c.ShouldBindJSON(&bookInput)
	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			errorMessages := []string{}
			for _, e := range validationErrors {
				errorMessage := fmt.Sprintf("error on field %s, condition: %s", e.Field(), e.ActualTag())
				errorMessages = append(errorMessages, errorMessage)
			}
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errorMessages,
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"title": bookInput.Title,
		"price": bookInput.Price,
	})
}