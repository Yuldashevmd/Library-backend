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

	bookHandler := NewBookHandler(bookService)
	// ENDPOINTS v1 group
	api := r.Group("/api/v1")
	{
		books := api.Group("/books")
		{
			books.GET("", bookHandler.GetBooks)
			books.GET("/:id", bookHandler.GetBookByID)
			books.POST("", bookHandler.CreateBook)
			books.PUT("/:id", bookHandler.UpdateBook)
			books.DELETE("/:id", bookHandler.DeleteBook)
		}
	}
	return r
}
