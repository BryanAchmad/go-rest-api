package main

import (
	"fmt"
	"log"

	"github.com/BryanAchmad/go-rest-api/internal/container"
	"github.com/BryanAchmad/go-rest-api/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/go-rest-api?charset=utf8mb4&parseTime=True&loc=Local"
  	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Connection to DB failed")
	}
	fmt.Println("Connection to DB Success")

	db.AutoMigrate(&models.Book{})

	bookContainer := container.BuildContainerBook(db)
	
	router := gin.Default()

	v1 := router.Group("/v1")
	v1.POST("/books", bookContainer.PostBookHandler)
	v1.GET("/books", bookContainer.GetAllBooksHandler)
	v1.GET("/book/:id", bookContainer.GetBookHandler)
	v1.POST("/book/:id", bookContainer.UpdateBookHandler)

	err = router.Run(":8000")
	if err != nil {
		log.Fatal("failed to start server: ", err)
	}
}
