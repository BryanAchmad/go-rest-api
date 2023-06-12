package container

import (
	"github.com/BryanAchmad/go-rest-api/internal/handlers"
	"github.com/BryanAchmad/go-rest-api/internal/repositories"
	"github.com/BryanAchmad/go-rest-api/internal/services"
	"gorm.io/gorm"
)

func BuildContainerBook(db *gorm.DB) *handlers.BookHandler {
	bookRepository := repositories.NewBookRepository(db)
	bookService := services.NewBookService(bookRepository)
	bookHandler := handlers.NewBookHandler(bookService)
	return bookHandler
}