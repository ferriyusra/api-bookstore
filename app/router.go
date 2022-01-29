package app

import (
	"ferri/api-bookstore/controller"
	"ferri/api-bookstore/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(categoryController controller.CategoryController, bookController controller.BookController) *httprouter.Router {

	router := httprouter.New()

	router.GET("/api/books", bookController.FindAll)
	router.GET("/api/books/:bookId", bookController.FindById)
	router.POST("/api/books", bookController.Create)
	router.PUT("/api/books/:bookId", bookController.Update)
	router.DELETE("/api/books/:bookId", bookController.Delete)

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
