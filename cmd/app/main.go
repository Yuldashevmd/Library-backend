package main

import (
	"context"
	"fmt"
	"library/internal/handler"
	"library/internal/repository"
	"library/internal/service"
)

func main() {
	// DB connection
	db, err := repository.NewDB()
	if err != nil {
		panic(err)
	}
	defer db.Close(context.Background())

	bookRepo := repository.NewBookRepository(db)
	bookService := service.NewBookService(bookRepo)
	r := handler.InitRouter(bookService)

	fmt.Println("Server 8080-portda ishga tushmoqda...")
	// Start the server
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
