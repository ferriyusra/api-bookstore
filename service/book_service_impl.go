package service

import (
	"context"
	"database/sql"
	"ferri/api-bookstore/exception"
	"ferri/api-bookstore/helper"
	"ferri/api-bookstore/model/domain"
	"ferri/api-bookstore/model/web"
	"ferri/api-bookstore/repository"

	"github.com/go-playground/validator"
)

type BookServiceImpl struct {
	BookRepository repository.BookRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewBookService(bookRepository repository.BookRepository, DB *sql.DB, validate *validator.Validate) BookService {
	return &BookServiceImpl{
		BookRepository: bookRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *BookServiceImpl) Create(ctx context.Context, request web.BookCreateRequest) web.BookResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	book := domain.Book{
		CategoryId:    request.CategoryId,
		Title:         request.Title,
		Author:        request.Author,
		Publisher:     request.Publisher,
		PublishedDate: request.PublishedDate,
		Price:         request.Price,
		Stock:         request.Stock,
	}

	book = service.BookRepository.Save(ctx, tx, book)

	return helper.ToBookResponse(book)
}

func (service *BookServiceImpl) Update(ctx context.Context, request web.BookUpdateRequest) web.BookResponse {

	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	book, err := service.BookRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	book.Title = request.Title
	book.Author = request.Author
	book.Publisher = request.Publisher
	book.PublishedDate = request.PublishedDate
	book.Publisher = request.Publisher
	book.Price = request.Price
	book.Stock = request.Stock
	book.CategoryId = request.CategoryId

	book = service.BookRepository.Update(ctx, tx, book)

	return helper.ToBookResponse(book)

}

func (service *BookServiceImpl) Delete(ctx context.Context, bookId int) {

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	category, err := service.BookRepository.FindById(ctx, tx, bookId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))

	}
	service.BookRepository.Delete(ctx, tx, category)

}

func (service *BookServiceImpl) FindById(ctx context.Context, bookId int) web.BookResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	book, _ := service.BookRepository.FindById(ctx, tx, bookId)

	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToBookResponse(book)

}

func (service *BookServiceImpl) FindAll(ctx context.Context) []web.BookResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	books := service.BookRepository.FindAll(ctx, tx)

	return helper.ToBookResponses(books)
}
