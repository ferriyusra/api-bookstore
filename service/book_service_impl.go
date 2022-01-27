package service

import (
	"context"
	"database/sql"
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

	book := domain.BookCreateOrUpdate{
		CategoryId:    request.CategoryId,
		Title:         request.Title,
		Author:        request.Author,
		Publisher:     request.Publisher,
		PublishedDate: request.PublishedDate,
		Price:         request.Price,
		Stock:         request.Stock,
	}

	book = service.BookRepository.Save(ctx, tx, book)

	return helper.ToBookResponseCreateOrUpdate(book)
}

func (service *BookServiceImpl) FindAll(ctx context.Context) []web.BookResponses {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	books := service.BookRepository.FindAll(ctx, tx)

	return helper.ToBookResponses(books)
}
