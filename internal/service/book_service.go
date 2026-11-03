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

func (s *BookService) GetBookByID(ctx context.Context, id int) (*repository.Book, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *BookService) CreateBook(ctx context.Context, title, author string) error {
	book := repository.Book{
		Title:  title,
		Author: author,
	}
	return s.repo.Create(ctx, book)
}

func (s *BookService) UpdateBook(ctx context.Context, id int, title, author string) error {
	book := repository.Book{
		ID:     id,
		Title:  title,
		Author: author,
	}
	return s.repo.Update(ctx, book)
}

func (s *BookService) DeleteBook(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
