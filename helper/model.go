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

func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	var categoryResponses []web.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}

	return categoryResponses
}

func ToBookResponse(book domain.Book) web.BookResponse {
	return web.BookResponse{
		Id:            book.Id,
		CategoryId:    book.CategoryId,
		CategoryName:  book.CategoryName,
		Title:         book.Title,
		Author:        book.Author,
		Publisher:     book.Publisher,
		PublishedDate: book.PublishedDate,
		Stock:         book.Stock,
		Price:         book.Price,
	}
}

func ToBookResponses(books []domain.Book) []web.BookResponse {
	var bookResponses []web.BookResponse
	for _, book := range books {
		bookResponses = append(bookResponses, ToBookResponse(book))
	}

	return bookResponses
}
