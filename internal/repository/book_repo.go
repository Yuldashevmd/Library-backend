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

func (r *BookRepository) GetByID(ctx context.Context, id int) (*Book, error) {
	row := r.db.QueryRow(ctx, "select id, title, author from books where id = $1", id)
	var b Book
	if err := row.Scan(&b.ID, &b.Title, &b.Author); err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &b, nil
}

func (r *BookRepository) Create(ctx context.Context, book Book) error {
	_, err := r.db.Exec(ctx, "insert into books (title, author) values ($1, $2)", book.Title, book.Author)
	return err
}

func (r *BookRepository) Delete(ctx context.Context, id int) error {
	_, err := r.db.Exec(ctx, "delete from books where id = $1", id)
	return err
}

func (r *BookRepository) Update(ctx context.Context, book Book) error {
	_, err := r.db.Exec(ctx, "update books set title = $1, author = $2 where id = $3", book.Title, book.Author, book.ID)
	return err
}
