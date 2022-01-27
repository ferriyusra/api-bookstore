package repository

import (
	"context"
	"database/sql"
	"ferri/api-bookstore/model/domain"
)

type BookRepository interface {
	Save(ctx context.Context, tx *sql.Tx, book domain.BookCreateOrUpdate) domain.BookCreateOrUpdate
	FindById(ctx context.Context, tx *sql.Tx, bookId int) (domain.Book, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Book
}
