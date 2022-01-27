package helper

import (
	"ferri/api-bookstore/model/domain"
	"ferri/api-bookstore/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToCategoryResponses(categries []domain.Category) []web.CategoryResponse {
	var categoryResponses []web.CategoryResponse
	for _, category := range categries {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}

	return categoryResponses
}

func ToBookResponseCreateOrUpdate(book domain.BookCreateOrUpdate) web.BookResponse {
	return web.BookResponse{
		Id:            book.Id,
		CategoryId:    book.CategoryId,
		Title:         book.Title,
		Author:        book.Author,
		Publisher:     book.Publisher,
		PublishedDate: book.PublishedDate,
		Stock:         book.Stock,
		Price:         book.Price,
	}
}

func ToBookResponse(book domain.Book) web.BookResponses {
	return web.BookResponses{
		Id:            book.Id,
		Category:      book.Category,
		Title:         book.Title,
		Author:        book.Author,
		Publisher:     book.Publisher,
		PublishedDate: book.PublishedDate,
		Stock:         book.Stock,
		Price:         book.Price,
	}
}

func ToBookResponses(books []domain.Book) []web.BookResponses {
	var bookResponses []web.BookResponses
	for _, book := range books {
		bookResponses = append(bookResponses, ToBookResponse(book))
	}

	return bookResponses
}
