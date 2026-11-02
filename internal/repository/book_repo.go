package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
)

// MODEL
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

type BookRepository struct {
	db *pgx.Conn
}

func NewBookRepository(db *pgx.Conn) *BookRepository {
	return &BookRepository{db: db}
}

func (r *BookRepository) GetAll(ctx context.Context) ([]Book, error) {
	rows, err := r.db.Query(ctx, "select id, title, author from books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var b Book
		if err := rows.Scan(&b.ID, &b.Title, &b.Author); err != nil {
			return nil, err
		}
		books = append(books, b)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return books, nil
}
