package handler

import (
	"library/internal/service"

	"github.com/gin-gonic/gin"
)

func InitRouter(bookService *service.BookService) *gin.Engine {
	r := gin.Default()

	// CORS middleware
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")

		if c.Request.Method == "HTTP_METHOD_OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// ENDPOINTS v1 group
	api := r.Group("/api/v1")
	{
		api.GET("/books", func(c *gin.Context) {
			books, err := bookService.GetAllBooks(c.Request.Context())
			if err != nil {
				c.JSON(500, gin.H{"error": "Failed to fetch books"})
				return
			}
			c.JSON(200, gin.H{"books": books})
		})
	}
	return r
}
