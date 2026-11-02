package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func NewDB() (*pgx.Conn, error) {
	ctx := context.Background()
	connectString := "postgres://postgres:1234@localhost:5432/library_db"
	conn, err := pgx.Connect(ctx, connectString)
	if err != nil {
		return nil, fmt.Errorf("bazaga ulanib bo'lmadi: %w", err)
	}
	if err := conn.Ping(ctx); err != nil {
		return nil, fmt.Errorf("baza javob bermayapti: %v", err)
	}
	fmt.Println("PostgreSQL bazasiga muvaffaqiyatli ulandi!")
	return conn, nil
}
