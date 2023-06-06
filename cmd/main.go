package main

import (
	"net/http"
	"go-rest-api/internal/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", rootHandler)
	v1.GET("/hello", helloHandler)
	v1.GET("/books/:id", booksHandler)
	v1.GET("/query", queryHandler)
	v1.POST("/books", handlers.postBookHandler)

	router.Run(":8000")
}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name" : "Bryan",
		"bio" : "I'm a unawarded professional programmer",
	})
}

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"title" : "Hello Page",
		"subtitle" : "This page contains something not usefull",
	})
}

func booksHandler(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"id" : id})
}

func queryHandler(c *gin.Context) {
	id := c.Query("title")
	
	c.JSON(http.StatusOK, gin.H{"id" : id})
}

