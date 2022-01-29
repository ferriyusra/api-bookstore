package repository

import (
	"context"
	"database/sql"
	"errors"
	"ferri/api-bookstore/helper"
	"ferri/api-bookstore/model/domain"
	"time"
)

type BookRepositoryImpl struct {
}

func NewBookRepository() BookRepository {
	return &BookRepositoryImpl{}
}

func (repository *BookRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, book domain.Book) domain.Book {
	SQL := "INSERT INTO book(category_id, title, author, publisher, published_date, price, stock) VALUES (?, ?, ?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, book.CategoryId, book.Title, book.Author, book.Publisher, book.PublishedDate, book.Price, book.Stock)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	book.Id = int(id)

	return book

}

func (repository *BookRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, book domain.Book) domain.Book {

	SQL := "UPDATE book SET title = ?, author = ?, publisher = ?, published_date = ?, price = ?, stock = ?, category_id = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, book.Title, book.Author, book.Publisher, book.PublishedDate, book.Price, book.Stock, book.CategoryId, book.Id)
	helper.PanicIfError(err)

	return book

}

func (repository *BookRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Book) {
	SQL := "DELETE from book WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Id)
	helper.PanicIfError(err)

}

func (repository *BookRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, bookId int) (domain.Book, error) {

	SQL := "SELECT book.id, book.category_id, book.title, book.author, book.publisher, book.published_date, book.price, book.stock, category.name FROM book LEFT JOIN category ON book.category_id = category.id WHERE book.id = ?"
	rows, err := tx.QueryContext(ctx, SQL, bookId)
	helper.PanicIfError(err)

	defer rows.Close()

	book := domain.Book{}
	if rows.Next() {
		var pubDate time.Time

		err := rows.Scan(&book.Id, &book.CategoryId, &book.Title, &book.Author, &book.Publisher, &pubDate, &book.Price, &book.Stock, &book.CategoryName)
		helper.PanicIfError(err)

		book.PublishedDate = pubDate.Format("02-01-2006")

		return book, nil
	} else {
		return book, errors.New("book is not found")
	}

}

func (repository *BookRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Book {
	SQL := "SELECT book.id, book.category_id, book.title, book.author, book.publisher, book.published_date, book.price, book.stock, category.name FROM book LEFT JOIN category ON book.category_id = category.id"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	defer rows.Close()

	var books []domain.Book
	for rows.Next() {
		book := domain.Book{}

		var pubDate time.Time

		err := rows.Scan(&book.Id, &book.CategoryId, &book.Title, &book.Author, &book.Publisher, &pubDate, &book.Price, &book.Stock, &book.CategoryName)
		helper.PanicIfError(err)

		book.PublishedDate = pubDate.Format("02-01-2006")
		books = append(books, book)
	}
	return books
}
