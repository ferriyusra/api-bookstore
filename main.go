package main

import (
	"ferri/api-bookstore/app"
	"ferri/api-bookstore/controller"
	"ferri/api-bookstore/helper"
	"ferri/api-bookstore/middleware"
	"ferri/api-bookstore/repository"
	"ferri/api-bookstore/service"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator"
)

func main() {
	db := app.NewDB()
	validate := validator.New()

	bookRepository := repository.NewBookRepository()
	bookService := service.NewBookService(bookRepository, db, validate)
	bookController := controller.NewBookController(bookService)

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController, bookController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
