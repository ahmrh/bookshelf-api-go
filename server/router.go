package server

import (
	"github.com/ahmrh/bookshelf-api-go/controllers"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)
	router.GET("/health", health.Status)
	
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	bookGroup := router.Group("books")
	{
		book := new(controllers.BookController)
		
		bookGroup.POST("", book.Create)

		bookGroup.GET("", book.Read)
		bookGroup.GET("/:id", book.Read)

		bookGroup.PUT("/:id", book.Update)

		bookGroup.DELETE("/:id", book.Delete)
	}

	return router
}