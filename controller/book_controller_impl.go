package controller

import (
	"ferri/api-bookstore/helper"
	"ferri/api-bookstore/model/web"
	"ferri/api-bookstore/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type BookControllerImpl struct {
	BookService service.BookService
}

func NewBookController(bookService service.BookService) BookController {
	return &BookControllerImpl{
		BookService: bookService,
	}
}

func (controller *BookControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	bookCreateRequest := web.BookCreateRequest{}
	helper.ReadFromRequestBody(request, &bookCreateRequest)

	bookResponse := controller.BookService.Create(request.Context(), bookCreateRequest)
	BookWebResponse := web.BookWebResponseCreateOrUpdate{
		Code:   200,
		Status: "OK",
		Data:   bookResponse,
	}

	helper.WriteToResponseBody(writer, BookWebResponse)
}

func (controller *BookControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	bookResponses := controller.BookService.FindAll(request.Context())
	BookWebResponse := web.BookWebResponse{
		Code:      200,
		Status:    "OK",
		CountData: len(bookResponses),
		Data:      bookResponses,
	}

	helper.WriteToResponseBody(writer, BookWebResponse)

}
