package controller

import (
	"ferri/api-bookstore/helper"
	"ferri/api-bookstore/model/web"
	"ferri/api-bookstore/service"
	"net/http"
	"strconv"

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
	BookWebResponse := web.BookWebResponse{
		Code:   200,
		Status: "OK",
		Data:   bookResponse,
	}

	helper.WriteToResponseBody(writer, BookWebResponse)
}

func (controller *BookControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	bookUpdateRequest := web.BookUpdateRequest{}
	helper.ReadFromRequestBody(request, &bookUpdateRequest)

	bookId := params.ByName("bookId")
	id, err := strconv.Atoi(bookId)
	helper.PanicIfError(err)

	bookUpdateRequest.Id = id

	bookResponse := controller.BookService.Update(request.Context(), bookUpdateRequest)
	BookWebResponse := web.BookWebResponse{
		Code:   200,
		Status: "OK",
		Data:   bookResponse,
	}

	helper.WriteToResponseBody(writer, BookWebResponse)

}
func (controller *BookControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	bookId := params.ByName("bookId")
	id, err := strconv.Atoi(bookId)
	helper.PanicIfError(err)

	controller.BookService.Delete(request.Context(), id)
	BookWebResponse := web.BookWebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, BookWebResponse)

}

func (controller *BookControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	bookId := params.ByName("bookId")
	id, err := strconv.Atoi(bookId)
	helper.PanicIfError(err)

	bookResponse := controller.BookService.FindById(request.Context(), id)

	BookWebResponse := web.BookWebResponse{
		Code:   200,
		Status: "OK",
		Data:   bookResponse,
	}

	helper.WriteToResponseBody(writer, BookWebResponse)

}

func (controller *BookControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	bookResponses := controller.BookService.FindAll(request.Context())
	BookWebResponse := web.BookWebResponse{
		Code:   200,
		Status: "OK",
		Data:   bookResponses,
	}

	helper.WriteToResponseBody(writer, BookWebResponse)

}
