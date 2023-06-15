package main

import (
	"log"

	"github.com/BryanAchmad/go-rest-api/internal/config"
	"github.com/BryanAchmad/go-rest-api/internal/container"
	"github.com/BryanAchmad/go-rest-api/internal/middleware"
	"github.com/BryanAchmad/go-rest-api/internal/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectDB()

	db.AutoMigrate(&models.Book{})

	bookContainer := container.BuildContainerBook(db)
	
	router := gin.Default()

	router.Use(middleware.LoggerMiddleware)
	v1 := router.Group("/v1")

	v1.POST("/books", bookContainer.PostBookHandler)
	v1.GET("/books", bookContainer.GetAllBooksHandler)
	v1.GET("/book/:id", bookContainer.GetBookHandler)
	v1.POST("/book/:id", bookContainer.UpdateBookHandler)

	err := router.Run(":8000")
	if err != nil {
		log.Fatal("failed to start server: ", err)
	}
}
