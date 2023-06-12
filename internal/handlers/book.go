package handlers

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/BryanAchmad/go-rest-api/internal/models"
	"github.com/BryanAchmad/go-rest-api/internal/services"
	"github.com/BryanAchmad/go-rest-api/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}
type BookHandler struct {
	bookService services.Service
}

func NewBookHandler(bookService services.Service) *BookHandler {
	return &BookHandler{bookService}
}


func (h *BookHandler) PostBookHandler(c *gin.Context) {
	var bookInput models.BookInput

	err := c.ShouldBindJSON(&bookInput)
	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			errorMessages := make([]string, len(validationErrors))
			for i, e := range validationErrors {
				errorMessages[i] = fmt.Sprintf("error on field %s, condition: %s", e.Field(), e.ActualTag())
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

	book, err := h.bookService.Create(bookInput)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": map[string]interface{}{"message" : "Book successfully saved", "data": book},
	})
}

func (h *BookHandler) GetAllBooksHandler(c *gin.Context) {
	books, err := h.bookService.FindAll();

	if err != nil {
		c.JSON(http.StatusBadRequest, Response{Success: true, Error: "failed To get Data"})
		return
	}

	var booksResponse []utils.BookResponse
	
	for _, b := range books {
		bookResponse := convertBookReponse(b)

		booksResponse = append(booksResponse, bookResponse)
	}

	c.JSON(http.StatusOK, Response{Success: true, Data: booksResponse})

}

func (h *BookHandler) GetBookHandler(c *gin.Context)  {
	bookIDStr := c.Param("id")

	bookID, err := strconv.Atoi(bookIDStr)
	log.Println("id", bookID)

	if err != nil {
		response := Response{
			Success: false,
			Error:   "Invalid book ID",
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	book, err := h.bookService.FindOne(bookID)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Book not found
			response := Response{
				Success: false,
				Error:   "Book not found",
			}
			c.JSON(http.StatusNotFound, response)
			return
		}

		response := Response{
			Success: false,
			Error:   "Failed to retrieve book",
		}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	bookResponse := convertBookReponse(book)

	response := Response{
		Success: true,
		Data:    bookResponse,
	}
	c.JSON(http.StatusOK, response)
}

func (h *BookHandler) UpdateBookHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, _ :=strconv.Atoi(idStr)

	var bookInput models.BookInput

	err := c.ShouldBindJSON(&bookInput)
	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			errorMessages := make([]string, len(validationErrors))
			for i, e := range validationErrors {
				errorMessages[i] = fmt.Sprintf("error on field %s, condition: %s", e.Field(), e.ActualTag())
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

	book, err := h.bookService.Update(id, bookInput)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": map[string]interface{}{"message" : "Book successfully saved", "data": book},
	})
}

func convertBookReponse(b models.Book) utils.BookResponse {
	return utils.BookResponse{
		ID: b.ID,
		Title: b.Title,
		Price: b.Price,
		Description: b.Description,
		Rating: b.Rating,
		Discount: b.Discount,
	}
}