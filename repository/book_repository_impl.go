package repository

import (
	"context"
	"database/sql"
	"ferri/api-bookstore/helper"
	"ferri/api-bookstore/model/domain"
)

type BookRepositoryImpl struct {
}

func NewBookRepository() BookRepository {
	return &BookRepositoryImpl{}
}

func (repository *BookRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, book domain.BookCreateOrUpdate) domain.BookCreateOrUpdate {
	SQL := "INSERT INTO book(category_id, title, author, publisher, published_date, price, stock) VALUES (?, ?, ?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, book.CategoryId, book.Title, book.Author, book.Publisher, book.PublishedDate, book.Price, book.Stock)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	book.Id = int(id)

	return book

}

// func (repository *BookRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, book domain.Category) domain.Category {
// 	SQL := "UPDATE category SET name = ? WHERE id = ?"
// 	_, err := tx.ExecContext(ctx, SQL, category.Name, category.Id)
// 	helper.PanicIfError(err)

// 	return category

// }
// func (repository *BookRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, book domain.Category) {
// 	SQL := "DELETE from category WHERE id = ?"
// 	_, err := tx.ExecContext(ctx, SQL, category.Id)
// 	helper.PanicIfError(err)

// }

// func (repository *BookRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, book int) (domain.Category, error) {
// 	SQL := "SELECT id,name FROM category WHERE id = ?"
// 	rows, err := tx.QueryContext(ctx, SQL, categoryId)
// 	helper.PanicIfError(err)

// 	defer rows.Close()

// 	category := domain.Category{}
// 	if rows.Next() {
// 		category := domain.Category{}

// 		err := rows.Scan(&category.Id, &category.Name)
// 		helper.PanicIfError(err)

// 		return category, nil
// 	} else {
// 		return category, errors.New("category is not found")
// 	}

// }

func (repository *BookRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Book {
	SQL := "SELECT book.id, book.category_id, book.title, book.author, book.publisher, book.published_date, book.price, book.stock, category.name FROM book LEFT JOIN category ON book.category_id = category.id"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	defer rows.Close()

	var books []domain.Book
	for rows.Next() {
		book := domain.Book{}
		// getDate := &book.PublishedDate

		// date := *getDate

		// newDateFormat, _ := time.Parse("01-02-2006", date)

		err := rows.Scan(&book.Id, &book.CategoryId, &book.Title, &book.Author, &book.Publisher, &book.PublishedDate, &book.Price, &book.Stock, &book.Category)
		helper.PanicIfError(err)

		books = append(books, book)
	}
	return books
}
