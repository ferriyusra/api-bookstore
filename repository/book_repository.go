package repository

import (
	"context"
	"database/sql"
	"ferri/api-bookstore/model/domain"
)

type BookRepository interface {
	Save(ctx context.Context, tx *sql.Tx, book domain.Book) domain.Book
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Book
}
