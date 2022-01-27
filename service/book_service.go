package service

import (
	"context"
	"ferri/api-bookstore/model/web"
)

type BookService interface {
	Create(ctx context.Context, request web.BookCreateRequest) web.BookResponse
	FindById(ctx context.Context, bookId int) web.BookResponses
	FindAll(ctx context.Context) []web.BookResponses
}
