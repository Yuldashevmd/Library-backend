package handler

import (
	"library/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

func InitRouter(bookService *service.BookService) *gin.Engine {
	r := gin.Default()

	// CORS middleware
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")

		if c.Request.Method == "HTTP_METHOD_OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
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
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch books"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"books": books})
		})

		api.POST("/books", func(c *gin.Context) {
			var input CreateBookInput
			if err := c.ShouldBindJSON(&input); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
				return
			}

			err := bookService.CreateBook(c.Request.Context(), input.Title, input.Author)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book"})
				return
			}
			c.JSON(http.StatusCreated, gin.H{"message": "Book created successfully"})
		})
	}
	return r
}
