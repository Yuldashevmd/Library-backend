package service

import (
	"context"
	"library/internal/repository"
)

type BookService struct {
	repo *repository.BookRepository
}

func NewBookService(repo *repository.BookRepository) *BookService {
	return &BookService{repo: repo}
}

func (s *BookService) GetAllBooks(ctx context.Context) ([]repository.Book, error) {
	return s.repo.GetAll(ctx)
}
